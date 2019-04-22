package main

import (
	"fmt"
	"github.com/hjiangsse/conf"
	"github.com/hjiangsse/decpool"
	"github.com/hjiangsse/msgqueue"
	"github.com/hjiangsse/rcpool"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

//config file path
const TEC_BUS_MAP_CNF string = "./conf/tec_bus_map.txt"
const TEC_RC_MAP_CNF string = "./conf/tec_rc_map.txt"
const RC_ADDR_MAP_CNF string = "./conf/rc_addr_map.txt"

//load global config; only read by go routines
var tec_bus_map map[string]string = conf.LoadPbuMap(TEC_BUS_MAP_CNF)
var tec_rc_map map[string]string = conf.LoadPbuMap(TEC_RC_MAP_CNF)
var rc_addr_map map[string]string = conf.LoadRcToAddr(RC_ADDR_MAP_CNF)

func main() {
	//initialize rc connection pool
	initRcConn()

	go msgqueue.TellAccount()

	listener, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		log.Fatal(err)
	}

	//the channel of handleConn and SendToRc
	rcsend_chan := make(chan int64, 10)

	//goroutine send decsim's message to RC
	go sendMsgToRc(&rcsend_chan)

	//process messages from RC
	procMsgFromRc()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn, &rcsend_chan) //go routines handle a single client
	}
}

/*
* Init connection map between cs and rc
 */
func initRcConn() {
	for rc_num, addr := range rc_addr_map {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			log.Fatal(err)
			return
		}

		if !rcpool.AddConnMap(rc_num, conn) {
			fmt.Printf("init rc %s failed!\n", rc_num)
		}
	}
}

/*
* handle messages from decsim
 */
func handleConn(c net.Conn, cl *chan int64) {
	var buff [1024]byte
	var tech_pbu string

	for {
		n, err := c.Read(buff[0:])
		if err != nil { //net connection broken
			if decpool.DelConnMap(tech_pbu) {
				fmt.Printf("%s delete conn map ok!", tech_pbu)
			}
			return
		}

		//get the message(client sent it)
		buff_msg := string(buff[:n])
		buff_head, buff_body := deal_msg(buff_msg)

		switch buff_head {
		case "login":
			tech_pbu = strings.TrimSpace(buff_body) /* login|techpbu */

			//get semaphone and update the connection map
			if decpool.AddConnMap(tech_pbu, c) {
				fmt.Printf("%s add conn map ok!\n", buff_body)
			}

			//send login message to RC, then wait login Responce
			login_slice := []byte(buff_msg)
			rcNo := tec_rc_map[tech_pbu]
			rcConn := rcpool.GetConnMap(rcNo)

			_, err := rcConn.Write(login_slice[:])
			if err != nil {
				log.Fatal(err)
				return
			}
		case "msg":
			//process the message send by client
			body_segs := strings.Split(buff_body, ",")
			bus_pbu := body_segs[1] /* msg|A0001,B0001,This is a message! */

			if tec_bus_map[bus_pbu] != tech_pbu {
				fmt.Fprintf(os.Stderr, "Invalid tech and bus pbu map!\n")
				return
			}

			msgNum := msgqueue.AddMsg(buff_body)
			*cl <- msgNum
		default:
			fmt.Println("Unknow message type!")
		}
	}
}

/*
* Send Ord Msg to RC
 */
func sendMsgToRc(cl *chan int64) {
	for {
		msgNo := <-*cl

		msg := msgqueue.GetMsg(msgNo)
		msg_fields := strings.Split(msg, ",")
		tec_pbu := msg_fields[0]
		bus_pbu := msg_fields[1]
		real_msg := msg_fields[2]

		//routine msg to rc
		rcNo := tec_rc_map[tec_pbu]
		rcConn := rcpool.GetConnMap(rcNo)

		msgStr := fmt.Sprintf("%s|%s,%s,%d,%s", "msg", tec_pbu, bus_pbu, msgNo, real_msg)

		msgSli := []byte(msgStr)
		_, err := rcConn.Write(msgSli[:])
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}

/*
* Process messages from RC
 */
func procMsgFromRc() {
	rcConns := rcpool.GetAllConn()

	for rcNo, rcConn := range rcConns {
		go hndlRcConn(rcConn, rcNo)
	}
}

/*
* Process message receive from RC
 */
func hndlRcConn(conn net.Conn, rcNo string) {
	var buff [1024]byte

	for {
		n, err := conn.Read(buff[0:])
		if err != nil { //net connection broken
			if rcpool.DelConnMap(rcNo) {
				fmt.Printf("%s delete conn map ok!", rcNo)
			}
			return
		}

		recvMsg := string(buff[0:n])
		//fmt.Println(recvMsg)
		msgSegs := strings.Split(recvMsg, "|")
		bodySegs := strings.Split(msgSegs[1], ",")
		tech_pbu := bodySegs[0]

		switch msgSegs[0] {
		case "logrsp":
			//Send the rc login respond to delsim
			logStatus := bodySegs[1]

			//send login respond back to decsim
			decConn := decpool.GetConnMap(tech_pbu)
			rspSli := []byte(recvMsg)
			_, err := decConn.Write(rspSli[:])
			if err != nil {
				log.Fatal(err)
				return
			}

			//if login fail, delete connection from decsim
			if logStatus != "SUCCESS" {
				decpool.DelConnMap(tech_pbu)
			}
		case "msgrsp":
			//Send msg respond message to delsim
			decConn := decpool.GetConnMap(tech_pbu)
			rspMsg := fmt.Sprintf("%s|%s,%s,%s,%s", msgSegs[0], tech_pbu, bodySegs[1], bodySegs[2], bodySegs[3])
			rspSli := []byte(rspMsg)
			_, err := decConn.Write(rspSli[:])
			if err != nil {
				log.Fatal(err)
				return
			}

			//After send respond back, Delete Msg in the message Queue
			msgNo, err := strconv.ParseInt(bodySegs[2], 10, 64)
			if err != nil {
				log.Fatal(err)
				return
			}
			msgqueue.DelMsg(msgNo)
		}
	}
}

/*
* Process string message from client
* input: string net message
* output1: message header
* output2: message body
 */
func deal_msg(netmsg string) (string, string) {
	msg_fields := strings.Split(netmsg, "|")
	return msg_fields[0], msg_fields[1]
}

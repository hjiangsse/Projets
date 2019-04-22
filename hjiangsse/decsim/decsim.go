/*
* 报盘机模拟
 */
package main

import (
	"fmt"
	"github.com/hjiangsse/conf"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

//config file path
const TEC_BUS_MAP_CNF string = "./conf/tec_bus_map.txt"

//load global config; only read by go routines
var tec_bus_map map[string]string = conf.LoadPbuMap(TEC_BUS_MAP_CNF)

//time count channel
var time_chan chan time.Time = make(chan time.Time)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: PollClient PBU")
		return
	}

	//tec pbu check
	tecpbu := os.Args[1]
	if !checkValidTecPbu(tecpbu) {
		fmt.Fprintf(os.Stderr, "Invalid Tec Pbu!\n")
		os.Exit(1)
	}

	login_chan := make(chan string)
	conn, err := net.Dial("tcp", "localhost:8001")
	if err != nil {
		log.Fatal(err)
		return
	}

	defer conn.Close()

	go prcsCsRsp(conn, login_chan) //处理CS发回的响应

	go sendLogin(conn, tecpbu, login_chan)
	go startMsgSending(conn, tecpbu, login_chan)

	fmt.Scanln()
	fmt.Println("Done")
}

/*
* check if valid tec pbu
 */
func checkValidTecPbu(pbu string) bool {
	for _, val := range tec_bus_map {
		if val == pbu {
			return true
		}
	}
	return false
}

/*
* send login message to server
* login message is something like this: login|A0001
 */
func sendLogin(conn net.Conn, pbu string, log_chan chan string) {
	msg_slice := []byte("login|" + pbu + "|")
	_, err := conn.Write(msg_slice[:])
	if err != nil {
		log.Fatal(err)
		return
	}
}

/*
* sending messages to server
 */
func startMsgSending(conn net.Conn, tecpbu string, log_chan chan string) {
	var num = 1

	<-log_chan

	for {
		str_num := strconv.Itoa(num)

		msg_str := fmt.Sprintf("msg|%s,%s,This is msg%s|", tecpbu, getRandBusPbu(), str_num)

		msg_slice := []byte(msg_str)
		_, err := conn.Write(msg_slice[:])
		if err != nil {
			log.Fatal(err)
			return
		}

		time_chan <- time.Now()
		num++

		time.Sleep(100 * time.Microsecond)
	}
}

/*
* Process respond from CS
 */
func prcsCsRsp(conn net.Conn, log_chan chan string) {
	var buff [1024]byte

	for {
		n, err := conn.Read(buff[0:])
		if err != nil {
			//connection with CS fail
			log.Fatal(err)
			return
		}

		buffMsg := string(buff[0:n])
		buffSegs := strings.Split(buffMsg, "|")
		msgHead := buffSegs[0]

		switch msgHead {
		case "logrsp":
			msgSegs := strings.Split(buffSegs[1], ",")
			logRes := msgSegs[1]

			if logRes == "SUCCESS" {
				//如果登陆成功,开始发定单
				log_chan <- "login ok"
			}
		case "msgrsp":
			start := <-time_chan
			elapsed := time.Since(start)
			fmt.Printf("A message round took %s\n", elapsed)
		}
	}
}

func getRandBusPbu() string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	index := r.Intn(len(tec_bus_map))
	//get keys of the tec_bus_map
	keys := make([]string, 0, len(tec_bus_map))
	for k := range tec_bus_map {
		keys = append(keys, k)
	}

	return keys[index]
}

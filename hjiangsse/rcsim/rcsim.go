/*
* RC模拟器
 */
package main

import (
	"fmt"

	"log"
	"net"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("[Invalid parameters] usage: %s ip_addr port\n", os.Args[0])
		os.Exit(1)
	}

	ip_addr := os.Args[1]
	port := os.Args[2]

	addr := fmt.Sprintf("%s:%s", ip_addr, port)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn) //go routines handle a cs connection
	}
}

func handleConn(c net.Conn) {
	var buff [1024]byte

	for {
		n, err := c.Read(buff[0:])
		if err != nil { //net connection broken
			return
		}

		msg := string(buff[0:n])
		msgSegs := strings.Split(msg, "|")
		msgHead := msgSegs[0]
		msgBody := msgSegs[1]

		switch msgHead {
		case "login":
			rspMsg := fmt.Sprintf("%s|%s,%s", "logrsp", msgBody, "SUCCESS")
			rspSli := []byte(rspMsg)
			_, err := c.Write(rspSli[:])
			if err != nil {
				log.Fatal(err)
				return
			}
		case "msg":
			rspMsg := fmt.Sprintf("%s|%s rsp", "msgrsp", msgBody)
			fmt.Println(rspMsg)
			rspSli := []byte(rspMsg)
			_, err := c.Write(rspSli[:])
			if err != nil {
				log.Fatal(err)
				return
			}
		}
	}
}

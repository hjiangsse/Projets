package msgqueue

//package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu       sync.RWMutex
	msgNum   int64 = 0
	msgQueue       = make(map[int64]string)
)

func AddMsg(msg string) int64 {
	mu.Lock()
	defer mu.Unlock()

	msgNum++
	msgQueue[msgNum] = msg

	return msgNum
}

func DelMsg(msgNum int64) {
	mu.Lock()
	defer mu.Unlock()

	delete(msgQueue, msgNum)
}

func GetMsg(msgNum int64) string {
	mu.RLock()
	defer mu.RUnlock()

	return msgQueue[msgNum]
}

func TellAccount() {
	for {
		var count int64 = 0
		for _ = range msgQueue {
			count++
		}
		fmt.Printf("Num of Messages in msgQueue: %d\n", count)
		time.Sleep(1 * time.Second)
	}
}

/*
func main() {
	idx1 := AddMsg("This is the 1st msg!")
	idx2 := AddMsg("This is the 2st msg!")
	idx3 := AddMsg("This is the 3st msg!")
	fmt.Printf("idx1 = %d\n", idx1)
	fmt.Printf("idx2 = %d\n", idx2)
	fmt.Printf("idx3 = %d\n", idx3)
	DelMsg(idx2)
	fmt.Println(msgQueue)
}
*/

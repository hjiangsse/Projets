package rcpool

import (
	"fmt"
	"net"
	"sync"
	"time"
)

/*
* rc connection pool
* key: RC NODE NUM : RC01
* value: net.Conn
 */
var (
	mu      sync.RWMutex
	connMap = make(map[string]net.Conn)
)

/*
* Get conn hash map
 */
func GetAllConn() map[string]net.Conn {
	return connMap
}

/*
* Get entry in conn map
 */
func GetConnMap(rcnum string) net.Conn {
	mu.RLock()
	defer mu.RUnlock()

	return connMap[rcnum]
}

/*
* Add entry to conn map
 */
func AddConnMap(rcnum string, conn net.Conn) bool {
	mu.Lock()
	defer mu.Unlock()

	_, prs := connMap[rcnum]
	if prs == false {
		connMap[rcnum] = conn
		return true
	}

	return false
}

/*
* Delete entry in conn map
 */
func DelConnMap(rcnum string) bool {
	mu.Lock()
	defer mu.Unlock()

	_, prs := connMap[rcnum]
	if prs {
		delete(connMap, rcnum)
		return true
	}

	return false
}

/*
* Print the connection map
 */
func PrtConMap() {
	for {
		mu.RLock()
		fmt.Println(connMap)
		mu.RUnlock()
		time.Sleep(5 * time.Second)
	}
}

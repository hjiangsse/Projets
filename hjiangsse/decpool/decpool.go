package decpool

import (
	"fmt"
	"net"
	"sync"
	"time"
)

/*
* connection pool
* key: pbu
* value: net.Conn
 */
var (
	mu      sync.RWMutex //multiple reader, single writer
	connMap = make(map[string]net.Conn)
)

/*
* Add entry to conn map
 */
func AddConnMap(pbu string, conn net.Conn) bool {
	mu.Lock()
	defer mu.Unlock()

	_, prs := connMap[pbu]
	if !prs {
		connMap[pbu] = conn
		return true
	}
	return false
}

/*
* Get entry using tech pbu
 */
func GetConnMap(pbu string) net.Conn {
	mu.RLock()
	defer mu.RUnlock()

	return connMap[pbu]
}

/*
* Delete entry in conn map
 */
func DelConnMap(pbu string) bool {
	mu.Lock()
	defer mu.Unlock()

	_, prs := connMap[pbu]
	if prs {
		delete(connMap, pbu)
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

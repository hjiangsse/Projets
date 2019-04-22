package conf

//package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
* load pbu map from the config file
* the key is bus pbu and the value is tecpbu
* or the key is bus pbu and value is RC node number
* RC01=B0001,B0002,B0003
 */
func LoadPbuMap(conf_path string) map[string]string {
	pbumap := make(map[string]string)

	file, err := os.Open(conf_path)
	if err != nil {
		checkError(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		curr_line := scanner.Text()

		//skip comment lines
		if curr_line[0] != '#' {
			tec_pbu, bus_pbus := splitTecAndBus(curr_line)
			bus_pbu_slice := splitBusPbus(bus_pbus)

			for _, bus_pbu := range bus_pbu_slice {
				//test if one bus pbu belongs to multiple tec pbus
				if _, prs := pbumap[bus_pbu]; prs {
					err_msg := "one bus pbu blongs to multiple tec pbus.\n"
					fmt.Fprintf(os.Stderr, "Logical error: %s", err_msg)
					os.Exit(1)
				} else {
					pbumap[bus_pbu] = tec_pbu
				}
			}
		}
	}
	return pbumap
}

/*
* load rc node number to rc address hash
* key: RC01
* value: 127.0.0.1:9001
 */
func LoadRcToAddr(conf_path string) map[string]string {
	return revMap(LoadPbuMap(conf_path))
}

func splitTecAndBus(line string) (string, string) {
	segs := strings.Split(line, "=")
	return segs[0], segs[1]
}

func splitBusPbus(bus_pbus string) []string {
	busPbus := strings.Split(bus_pbus, ",")
	res := make([]string, len(busPbus))
	for i, val := range busPbus {
		res[i] = strings.TrimSpace(val)
	}
	return res
}

func revMap(input map[string]string) map[string]string {
	res := make(map[string]string)
	for k, v := range input {
		res[v] = k
	}
	return res
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

/*
func main() {
	pbu_map := LoadPbuMap("./pbumap.txt")
	fmt.Println(pbu_map)

	bus_to_rc := LoadPbuMap("./pburcmap.txt")
	fmt.Println(bus_to_rc)

	rc_to_addr := LoadRcToAddr("./rcaddr.txt")
	fmt.Println(rc_to_addr)

	//test tecpbu to busbpu check
	tecpbu := "A0001"
	buspbu := "B0002"
	if pbu_map[buspbu] == tecpbu {
		fmt.Println("OK!")
	} else {
		fmt.Println("Not OK!")
	}
}
*/

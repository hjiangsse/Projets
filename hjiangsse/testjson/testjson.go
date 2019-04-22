package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"os"
)

//one entry of config file
type CnfElem struct {
	TecPbu  string
	BizPbus []string
}

func main() {
	dat, err := ioutil.ReadFile("./tecbiz.json")
	if err != nil {
		fmt.Println("err: ", err)
	}

	var records []CnfElem
	if err = json.Unmarshal(dat, &records); err != nil {
		panic(err)
	}

	//fmt.Println("The number of records file: ", len(records))

	for _, elem := range records {
		fmt.Println(elem)
	}
	/*
				bolB, _ := json.Marshal(true)
				fmt.Println(string(bolB))

				intB, _ := json.Marshal(1)
				fmt.Println(string(intB))

				fltB, _ := json.Marshal(1.23)
				fmt.Println(string(fltB))

				strB, _ := json.Marshal("hjiang")
				fmt.Println(string(strB))

				sliB, _ := json.Marshal([]string{"apple", "orange", "green"})
				fmt.Println(string(sliB))

				mapB, _ := json.Marshal(map[string]string{"name": "hjiang", "sex": "man"})
				fmt.Println(string(mapB))

				//use json marshal a map, which key is a string, value is a slice of string
				mapD := map[string][]string{"T00001": {"B00001", "B00002", "B00003"},
					"T00002": {"B00003", "B00004", "B00005"}}
				mapE, _ := json.Marshal(mapD)
				fmt.Println(string(mapE))

			type ColorGroup struct {
				ID     int
				Name   string
				Colors []string
			}

			group := ColorGroup{
				ID:     1,
				Name:   "Reds",
				Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
			}

			b, err := json.Marshal(group)
			if err != nil {
				fmt.Println("error:", err)
			}
			fmt.Println(string(b))

			type CnfElem struct {
				TecPbu  string
				BizPbus []string
			}

			entry := CnfElem{
				TecPbu:  "A00001",
				BizPbus: []string{"B00001", "B00002", "B00003"},
			}

			c, err := json.Marshal(entry)
			if err != nil {
				fmt.Println("error:", err)
			}
			fmt.Println(string(c))

		entA := CnfElem{
			TecPbu:  "A00001",
			BizPbus: []string{"B00001", "B00002", "B00003"},
		}

		entB := CnfElem{
			TecPbu:  "A00002",
			BizPbus: []string{"B00004", "B00005", "B00006"},
		}

		entsli := []CnfElem{entA, entB}
		b, err := json.Marshal(entsli)
		if err != nil {
			fmt.Println("error", err)
		}
		fmt.Println(string(b))
	*/
}

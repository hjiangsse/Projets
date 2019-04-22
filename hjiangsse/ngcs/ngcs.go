package main

import (
	cscfg "./cscfg"
	"fmt"
	_ "github.com/sirupsen/logrus"
)

const TEC_BIZ_PATH = "./cscfg/tecbiz.json"
const RC_BIZ_PATH = "./cscfg/rcbiz.json"

func main() {
	bizToTecMap := make(map[string]string)
	if err := cscfg.LoadBizToTecMap(TEC_BIZ_PATH, &bizToTecMap); err != nil {
		fmt.Println(err)
	}
	fmt.Println(bizToTecMap)

	bizToRcMap := make(map[string]string)
	if err := cscfg.LoadBizToRcMap(RC_BIZ_PATH, &bizToRcMap); err != nil {
		fmt.Println(err)
	}
	fmt.Println(bizToRcMap)
}

package main

import (
	"fmt"
	csdat "github.com/hjiangsse/ngcs/csdat"
	_ "github.com/sirupsen/logrus"
)

func main() {
	bizToTecMap := make(map[string]string)
	if err := csdat.LoadBizToTecMap(&bizToTecMap); err != nil {
		fmt.Println(err)
	}
	fmt.Println(bizToTecMap)

	fmt.Println("\n")

	bizToRcMap := make(map[string]string)
	if err := csdat.LoadBizToRcMap(&bizToRcMap); err != nil {
		fmt.Println(err)
	}
	fmt.Println(bizToRcMap)
}

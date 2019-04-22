package cscnf

import (
	"fmt"
	"testing"
)

func Test_loadAllCnf_1(t *testing.T) {
	var confs csConfig
	if err := loadAllCnf(&confs); err != nil {
		fmt.Println(err)
		t.Error("load all config fail! Test Fail!")
	} else {
		t.Log("Test OK!")
	}
}

func Test_GetCsAddr(t *testing.T) {
	wantedAddr := "127.0.0.1:8001"
	getedAddr, err := GetCsAddr()
	if err != nil {
		t.Error("load cs config fail! Test Fail!")
	} else {
		if wantedAddr == getedAddr {
			t.Log("Test OK!")
		} else {
			t.Error("load cs config ok! But not wanted! Test Fail!")
		}
	}
}

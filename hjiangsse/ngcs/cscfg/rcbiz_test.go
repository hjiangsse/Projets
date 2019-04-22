package cscfg

import (
	"testing"
)

func Test_LoadRcBizConf_1(t *testing.T) {
	conf_path_1 := "./rcbiz.json"
	var confs []RcBizCnfElem
	if err := LoadRcBizConf(conf_path_1, &confs); err != nil {
		t.Error("valid tecbiz.json file, but test not pass!")
	} else {
		t.Log("valid tecbiz.json file, test pass!")
	}
}

/*
func Test_LoadRcBizConf_2(t *testing.T) {
	t.Error("Ha ha!")
}

func Test_LoadTecBizConf_2(t *testing.T) {
	conf_path_2 := "./tecbiz_wrong.json"
	var confs []CnfElem
	if err := LoadTecBizConf(conf_path_2, &confs); err != nil {
		t.Log("invalid tecbiz.json, return err, test pass!")
	} else {
		t.Error("invalid tecbiz.json, not return err, test failed!")
	}
}
*/

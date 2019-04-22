package cscfg

import (
	"testing"
)

func Test_loadTecBizConf_1(t *testing.T) {
	conf_path_1 := "./tecbiz.json"
	var confs []TecBizCnfElem
	if err := loadTecBizConf(conf_path_1, &confs); err != nil {
		t.Error("valid tecbiz.json file, but test not pass!")
	} else {
		t.Log("valid tecbiz.json file, test pass!")
	}
}

func Test_loadTecBizConf_2(t *testing.T) {
	conf_path_2 := "./tecbiz_wrong.json"
	var confs []TecBizCnfElem
	if err := loadTecBizConf(conf_path_2, &confs); err != nil {
		t.Log("invalid tecbiz.json, return err, test pass!")
	} else {
		t.Error("invalid tecbiz.json, not return err, test failed!")
	}
}

func Test_LoadBizToTecMap_1(t *testing.T) {
	bizTecMap := make(map[string]string)
	conf_path := "./tecbiz_multec.json"

	if err := LoadBizToTecMap(conf_path, &bizTecMap); err != nil {
		t.Log("bizpbu belong to multiple tecpbu in config file, err msg: ", err, "test OK!")
	} else {
		t.Error("Test failed!")
	}
}

package cscnf

import (
	"fmt"
	toml "github.com/BurntSushi/toml"
)

const csConfPath = "../conf/csconf.toml"

type csConfig struct {
	Titile string
	Csaddr addrInfo
	Rcaddr map[string]addrInfo
}

type addrInfo struct {
	Ip   string
	Port string
}

// Load all the configs, inner useage
func loadAllCnf(cscnf *csConfig) error {
	if _, err := toml.DecodeFile(csConfPath, cscnf); err != nil {
		return err
	}
	return nil
}

// Get Cs address
func GetCsAddr() (string, error) {
	var allConf csConfig
	if err := loadAllCnf(&allConf); err != nil {
		return "", err
	} else {
		csip := allConf.Csaddr.Ip
		csport := allConf.Csaddr.Port
		csaddr := fmt.Sprintf("%s:%s", csip, csport)
		return csaddr, nil
	}
}

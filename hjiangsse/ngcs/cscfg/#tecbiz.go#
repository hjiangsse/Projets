// Load tecpbu to bizpbu configure file
package cscfg

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

// Struct of one config entry
type TecBizCnfElem struct {
	TecPbu  string
	BizPbus []string
}

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

// Load config file, the result is a slice of CnfElem
func loadTecBizConf(confPath string, elms *[]TecBizCnfElem) error {
	// Test if the config file not exist
	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		log.Warn(err)
		return err
	}

	// Load file records into slice
	confs, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Warn(err)
		return err
	}

	if err = json.Unmarshal(confs, elms); err != nil {
		log.Warn(err)
		return err
	}

	return nil
}

// Load bizpbu-to-tecpbu hash, key is bizPbu and value is tecPbu
func LoadBizToTecMap(confPath string, bizToTec *map[string]string) error {
	var confElems []TecBizCnfElem
	if err := loadTecBizConf(confPath, &confElems); err != nil {
		log.Warn("Load TecBizConf Failed!")
		return err
	}

	for _, elem := range confElems {
		tecPbu := elem.TecPbu
		for _, biz := range elem.BizPbus {
			//the bizpbu to tecpbu have no duplicate, if one bizpbu blongs to multiple tecpbu, the config file is invalid
			_, prs := (*bizToTec)[biz]
			if prs {
				//the biz is already in the map, this is an error
				errmsg := fmt.Sprintf("%s belongs multiple tecpbus, this is a invalid config file!", biz)
				return errors.New(errmsg)
			} else {
				(*bizToTec)[biz] = tecPbu
			}
		}
	}

	return nil
}

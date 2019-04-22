// Load rcnum to bizpbu configure file
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
type RcBizCnfElem struct {
	RcName  string
	BizPbus []string
}

// Load config file, the result is a slice of CnfElem
func loadRcBizConf(confPath string, elms *[]RcBizCnfElem) error {
	// Test if the config file not exist
	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		log.Warn(err)
	}

	// Load file records into slice
	confs, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Warn(err)
	}

	if err = json.Unmarshal(confs, elms); err != nil {
		log.Warn(err)
	}

	return err
}

// Load bizpbu-to-rcnum hash, key is bizPbu and value is rcname
func LoadBizToRcMap(confPath string, bizToRc *map[string]string) error {
	var confElems []RcBizCnfElem
	if err := loadRcBizConf(confPath, &confElems); err != nil {
		log.Warn("Load TecBizConf Failed!")
		return err
	}

	for _, elem := range confElems {
		rcNam := elem.RcName
		for _, biz := range elem.BizPbus {
			//the bizpbu to rcname have no duplicate, if one bizpbu blongs to multiple rc, the config file is invalid
			_, prs := (*bizToRc)[biz]
			if prs {
				//the biz is already in the map, this is an error
				errmsg := fmt.Sprintf("%s belongs multiple Rcs, this is a invalid config file!", biz)
				return errors.New(errmsg)
			} else {
				(*bizToRc)[biz] = rcNam
			}
		}
	}

	return nil
}

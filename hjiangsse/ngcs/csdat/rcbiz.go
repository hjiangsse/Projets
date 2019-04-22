// Load rcnum to bizpbu configure file
package csdat

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

const rcToBizCnfFile = "data/rc_to_biz.json"

// Struct of one config entry
type RcBizCnfElem struct {
	RcName  string
	BizPbus []string
}

// Load config file, the result is a slice of CnfElem
func loadRcBizConf(elms *[]RcBizCnfElem) error {
	if _, err := os.Stat(rcToBizCnfFile); os.IsNotExist(err) {
		log.Warn(err)
	}

	confs, err := ioutil.ReadFile(rcToBizCnfFile)
	if err != nil {
		log.Warn(err)
	}

	if err = json.Unmarshal(confs, elms); err != nil {
		log.Warn(err)
	}

	return err
}

// Load bizpbu-to-rcnum hash, key is bizPbu and value is rcname
func LoadBizToRcMap(bizToRc *map[string]string) error {
	var confElems []RcBizCnfElem
	if err := loadRcBizConf(&confElems); err != nil {
		log.Warn("Load TecBizConf Failed!")
		return err
	}

	for _, elem := range confElems {
		rcNam := elem.RcName
		for _, biz := range elem.BizPbus {
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

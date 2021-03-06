package dl

import (
	log "github.com/sirupsen/logrus"

	"encoding/json"
)

func (self *Dl) Precompiling() {
	var err error
	var tmpList []json.RawMessage
	var tmpMap map[string]json.RawMessage

	// 如果是list
	if err = json.Unmarshal(self.AllStr, &tmpList); err == nil {
		if len(tmpList) < 1 {
			return
		}
		for tmpNum, tmpOne := range tmpList {
			data := &Dl{
				NodeIndex:  tmpNum,
				FatherNode: self,
				AllStr:     tmpOne,
			}
			data.Init()
			log.Debug("in Precompiling ", data.NodeName, self.SubNodeTree)
			self.SubNodeList = append(self.SubNodeList, data)
			data.Precompiling()
		}

		return
	}

	// 如果是map
	if err = json.Unmarshal(self.AllStr, &tmpMap); err == nil {
		log.Debug("get map", tmpMap)
		if len(tmpMap) < 1 {
			return
		}
		for tmpKey, tmpOne := range tmpMap {
			log.Debug("tmpKey:", tmpKey)
			log.Debug("tmpOne", string(tmpOne))
			data := &Dl{
				NodeName:   tmpKey,
				FatherNode: self,
				AllStr:     tmpOne,
			}
			data.Init()
			self.SubNodeTree[tmpKey] = data
			data.Precompiling()
			log.Debug(self.SubNodeTree)
		}
		return
	}

	if err = json.Unmarshal(self.AllStr, &self.TmpInterface); err == nil {
		return
	}

	panic(err)
}

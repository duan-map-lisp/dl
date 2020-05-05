package dl

import (
	log "github.com/sirupsen/logrus"
)

func (self *Dl) setCar() {
	Lambdas["car"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasNameForce("car")
		var err error
		var data *Dl
		var key string
		if len(self.SubNodeTree) >= 2 {
			if data, err = self.SubNodeGet("data"); err != nil {
				panic("'data' not found")
			}
			if data.CheckType() == "list" {
				tmpDataList := data.Call().(*Dl)
				if len(tmpDataList.SubNodeList) <= 0 {
					resI = nil
				} else {
					resI = tmpDataList.SubNodeList[0]
				}
			} else if data.CheckType() == "object" {
				if key, err = self.SubNodeGetSingleString("key"); err != nil {
					panic("'key' not found")
				}
				tmpDataObject := data.Call().(*Dl)
				var ok bool
				if resI, ok = tmpDataObject.SubNodeTree[key]; ok {
					return
				} else {
					resI = nil
					return
				}
				return
			} else if data.CheckType() == "null" {
				resI = nil
				return
			}
		} else if len(self.SubNodeList) == 2 {
			if data, err = self.SubNodeListGet(1); err != nil {
				panic("'data' not found")
			}
			if data.CheckType() == "list" || data.CheckType() == "null" {
				log.Info("???", data.String())
				log.Info("???", data.SubNodeList[0].String())

				tmpDataList := data.Call().(*Dl)
				if len(tmpDataList.SubNodeList) <= 0 {
					resI = nil
				} else {
					resI = tmpDataList.SubNodeList[0]
				}
			} else {
				panic("'car' format error")
			}
		} else if len(self.SubNodeList) == 3 {
			if data, err = self.SubNodeListGet(1); err != nil {
				panic("'data' not found")
			}
			if data.CheckType() == "object" || data.CheckType() == "null" {
				if key, err = self.SubNodeListGetSingleString(2); err != nil {
					panic("'key' not found")
				}
				tmpDataObject := data.Call().(*Dl)
				var ok bool
				if resI, ok = tmpDataObject.SubNodeTree[key]; ok {
					return
				} else {
					resI = nil
					return
				}
			} else {
				panic("'car' format error")
			}
		} else {
			panic("'car' format error")
		}

		return
	}
}

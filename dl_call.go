package dl

import (
	log "github.com/sirupsen/logrus"
)

func (self *Dl) Call() (resI interface{}) {
	// log.Debug("in Call ", self.TmpInterface)
	if self == nil {
		return
	}
	if self.TmpInterface != nil {
		resI = self.TmpInterface
		return
	}
	var err error
	var lambdaName string
	if len(self.SubNodeTree) >= 1 {
		if lambdaName, err = self.SubNodeGetSingleString("name"); err != nil {
			log.Debug(self.String)
			panic("'name' not found " + err.Error())
		}
	} else if len(self.SubNodeList) >= 1 {
		if lambdaName, err = self.SubNodeListGetSingleString(0); err != nil {
			log.Debug("......", self.String())
			panic("'name' not found " + err.Error())
		}
	} else {
		resI = nil
		return
	}

	if lambdaOne, ok := Lambdas[lambdaName]; ok {
		resI = lambdaOne(self)
	} else {
		panic("未定义的函数：" + lambdaName)
	}

	return
}

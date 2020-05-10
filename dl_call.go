package dl

import (
	log "github.com/sirupsen/logrus"
)

func (self *Dl) Call() (resI interface{}) {
	// log.Debug("in Call ", self.TmpInterface)
	switch tmpType := self.DataInterface.(type) {
	case map[string]*Dl:
	case []*Dl:
	default:
		resI = tmpType
		return
	}

	name, err := self.GetLambdasName()
	log.Info("get name is:", name)
	if err != nil {
		panic(err)
	}

	var lambdaName string
	lambdaNameI := self.GetSymbol(name)
	switch tmpLambdaName := lambdaNameI.(type) {
	case string:
		lambdaName = tmpLambdaName
	default:
		panic("symbol type not lambda")
	}

	if lambdaOne, ok := Lambdas[lambdaName]; ok {
		resI = lambdaOne(self)
	} else {
		panic("未定义的函数：" + lambdaName)
	}

	return
}

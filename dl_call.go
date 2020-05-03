package dl

import (
	log "github.com/sirupsen/logrus"
)

func (self *Dl) Call() (resI interface{}) {
	log.Debug("in Call ", self.TmpInterface)
	if self.TmpInterface != nil {
		resI = self.TmpInterface
		return
	}

	if len(self.TmpList) != 0 {
		resI = self.LambdaList()
		return
	}

	if len(self.TmpMap) != 0 {
		resI = self.LambdaMap()
		return
	}

	return
}

func (self *Dl) LambdaMap() (res interface{}) {
	var err error
	var lambdaName string
	lambdaName, err = self.SubNodeGetSingleString("name")
	if err != nil {
		panic(err)
	}

	lambdaOne := self.GetLambda(lambdaName)
	res = lambdaOne(self)
	return
}

func (self *Dl) LambdaList() (res interface{}) {
	var err error
	var lambdaName string
	lambdaName, err = self.SubNodeListGetSingleString(0)
	if err != nil {
		panic(err)
	}
	lambdaOne := self.GetLambda(lambdaName)
	res = lambdaOne(self)
	return
}

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
	var err error
	var lambdaName string
	if lambdaName, err = self.SubNodeGetSingleString("name"); err != nil {
		if lambdaName, err = self.SubNodeListGetSingleString(0); err != nil {
			resI = nil
			return
		}
	}
	lambdaOne := self.GetLambda(lambdaName)
	resI = lambdaOne(self)

	return
}

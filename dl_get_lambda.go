package dl

import (
	log "github.com/sirupsen/logrus"
)

func (self *Dl) GetLambda(lambdaName string) (res func(*Dl) interface{}) {
	log.Debug("in get lambda", self.NodeName, Lambdas)
	res, ok := Lambdas[lambdaName]
	if !ok {
		panic("未定义的函数：" + lambdaName)
	}

	return
}

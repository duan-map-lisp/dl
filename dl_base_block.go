package dl

import (
	log "github.com/sirupsen/logrus"
)

func (self *Dl) setBlock() {
	Lambdas["block"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasNameForce("block")
		// block无上限，只允许list模型
		if len(self.SubNodeTree) != 0 {
			panic("'block' must be list type")
		}
		for resIndex, resOne := range self.SubNodeList {
			if resIndex == 0 {
				continue
			}
			resI = resOne.Call()
			switch resTmp := resI.(type) {
			case *Dl:
				log.Info("block中间结果：", resTmp.String())
			default:
				log.Info("block中间结果：", resTmp)
			}
		}
		return
	}
}

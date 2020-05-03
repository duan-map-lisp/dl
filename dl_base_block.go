package dl

import (
	log "github.com/sirupsen/logrus"
)

func (self *Dl) setBlock() {
	Lambdas["block"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasNameForce("block")
		for _, resOne := range self.SubNodeList {
			self.BlockBreakFlag = false
			resI = resOne.Call()
			log.Debug("block 结果：", resI)
			if self.BlockBreakFlag {
				break
			}
		}
		return
	}
}

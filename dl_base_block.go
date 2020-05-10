package dl

import (
	log "github.com/sirupsen/logrus"
)

func (self *Dl) setBlock() {
	Lambdas["block"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasName("block")
		// block无上限，只允许list模型
		switch tmpType := self.DataInterface.(type) {
		case map[string]*Dl:
			panic("'block' must be list type")
		case []*Dl:
			for resIndex, resOne := range tmpType {
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
		default:
			panic("'block' format error")
		}

		return
	}
	self.Symbols["block"] = "block"
}

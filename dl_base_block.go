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
			var err error
			var flag string
			var res *Dl

			if resIndex == 0 {
				continue
			}
			resOneI := resOne.Call()

			switch resTmp := resOneI.(type) {
			case *Dl:
				if len(resTmp.SubNodeTree) >= 2 {
					if flag, err = resTmp.SubNodeGetSingleString("flag"); err != nil {
						panic("'flag' not found")
					}
					if res, err = resTmp.SubNodeGet("res"); err != nil {
						panic("'res' not found")
					}
				} else if len(resTmp.SubNodeList) == 2 {
					if flag, err = resTmp.SubNodeListGetSingleString(1); err != nil {
						panic("'flag' not found")
					}
					if res, err = resTmp.SubNodeListGet(2); err != nil {
						panic("'res' not found")
					}
				} else {
					panic("'get' format error")
				}
			default:
				panic("'block res' format error")
			}

			resI = res
			log.Info("block 中间结果：", res.String())
			if flag == "break" {
				log.Info("block break")
				break
			} else if flag == "continue" {
				continue
			}
		}
		return
	}
}

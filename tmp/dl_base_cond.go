package dl

import (
	log "github.com/sirupsen/logrus"
)

func (self *Dl) setCond() {
	Lambdas["cond"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasNameForce("cond")
		// cond无上限，只允许list模型
		if len(self.SubNodeTree) != 0 {
			panic("'cond' must be list type")
		}
		for resIndex, resOne := range self.SubNodeList {
			var err error
			var flag bool
			var res *Dl

			if resIndex == 0 {
				continue
			}
			resOneI := resOne.Call()

			switch resTmp := resOneI.(type) {
			case *Dl:
				if len(resTmp.SubNodeTree) >= 2 {
					if flag, err = resTmp.SubNodeGetSingleBool("flag"); err != nil {
						panic("'flag' not found")
					}
					if res, err = resTmp.SubNodeGet("res"); err != nil {
						panic("'res' not found")
					}
				} else if len(resTmp.SubNodeList) == 2 {
					if flag, err = resTmp.SubNodeListGetSingleBool(0); err != nil {
						panic("'flag' not found")
					}
					if res, err = resTmp.SubNodeListGet(1); err != nil {
						panic("'res' not found")
					}
				} else {
					panic("'get' format error")
				}
			default:
				panic("'cond res' format error")
			}

			resI = res
			log.Info("cond中间结果：", res.String())
			if flag {
				log.Info("cond break")
				break
			} else {
				continue
			}
		}
		return
	}
}

package dl

import (
	"crypto/sha512"
	"encoding/hex"
	log "github.com/sirupsen/logrus"
)

func (self *Dl) setLambda() {
	Lambdas["lambda"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasNameForce("lambda")
		// lambda操作比较复杂，list格式看的乱，只允许object格式
		if len(self.SubNodeList) != 0 {
			panic("'lambda' must be object type")
		} else if len(self.SubNodeTree) <= 1 {
			panic("'lambda' format error")
		}

		var err error
		var body *Dl
		if body, err = self.SubNodeGet("body"); err != nil {
			panic("'body' not found")
		}

		// 生成获得不重复的字符串
		var lambdaName string
		var has bool
		bodyStr := body.AllStr
		for {
			hashStr := sha512.Sum512(bodyStr)
			lambdaName = hex.EncodeToString(hashStr[:])

			if _, has = Lambdas[lambdaName]; has {
				bodyStr = []byte(lambdaName)
				continue
			} else {
				break
			}
		}
		Lambdas[lambdaName] = func(lambdaSelf *Dl) (lambdaRes interface{}) {
			log.Info("show args symbol:", lambdaSelf.Symbols)
			if lambdaSelf != nil {
				for symbolKey, symbolValue := range lambdaSelf.Symbols {
					// 传入的symbol在body内不能存在，防止外部参数覆盖lambda的内部symbol
					if _, ok := body.Symbols[symbolKey]; ok {
						panic("args symbol redefine in lambda")
					}
					body.Symbols[symbolKey] = symbolValue
				}
			}
			log.Debug("lambda begin call body", body.String())
			log.Debug("lambda begin call body SubNodeTree", body.SubNodeTree)
			lambdaRes = body.Call()
			return
		}

		resI = lambdaName

		return
	}
	return
}

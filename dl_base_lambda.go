package dl

import (
	"crypto/sha512"
	"encoding/hex"
	log "github.com/sirupsen/logrus"
)

func (self *Dl) setLambda() {
	Lambdas["lambda"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasNameForce("lambda")
		// lambda操作比较复杂，只允许object模型
		if len(self.SubNodeList) != 0 {
			panic("'lambda' must be object type")
		}
		var err error
		var body *Dl
		if body, err = self.SubNodeGet("body"); err != nil {
			panic("'body' not found")
		}

		// args[形参]实参
		var args *Dl
		if args, err = self.SubNodeGet("args"); err != nil {
			args = nil
			err = nil
		}

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
			if args != nil {
				for argsIndex, _ := range args.SubNodeList {
					argsSymbol, lambdaErr := args.SubNodeListGetSingleString(argsIndex)
					if lambdaErr != nil {
						panic("args type must []string")
					}
					log.Debug("lambdaSelf", lambdaSelf)
					log.Debug("lambdaSelf", lambdaSelf.Symbols)
					resSymbolOne := lambdaSelf.GetSymbol(argsSymbol)
					body.Symbols[argsSymbol] = resSymbolOne
				}
			}
			log.Debug("lambda begin call body", body)
			log.Debug("lambda begin call body SubNodeTree", body.SubNodeTree)
			lambdaRes = body.Call()
			return
		}

		resI = lambdaName

		return
	}
	return
}

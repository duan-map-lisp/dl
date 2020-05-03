package dl

import (
	"crypto/sha512"
	"encoding/hex"
	log "github.com/sirupsen/logrus"
)

func (self *Dl) setLambda() {
	Lambdas["lambda"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasName("lambda")
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

		hashStr := sha512.Sum512(body.AllStr)
		lambdaName := hex.EncodeToString(hashStr[:])
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

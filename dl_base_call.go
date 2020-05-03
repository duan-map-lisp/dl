package dl

import (
	log "github.com/sirupsen/logrus"
)

func (self *Dl) setCall() {
	Lambdas["call"] = func(self *Dl) (resI interface{}) {
		log.Debug("in Call", self.TmpInterface)
		self.CheckLambdasNameForce("call")
		var err error
		var lambda string
		if lambda, err = self.SubNodeGetSingleString("lambda"); err != nil {
			panic("'lambda' not found")
		}
		var args *Dl
		lambdaArgsDl := &Dl{
			FatherNode: self,
		}
		lambdaArgsDl.Init()
		if args, err = self.SubNodeGet("args"); err != nil {
			args = nil
			err = nil
		}
		if args != nil {
			log.Debug(args.SubNodeTree)
			for argsFormal, _ := range args.SubNodeTree {
				argsSymbol, lambdaErr := args.SubNodeGetSingleString(argsFormal)
				if lambdaErr != nil {
					log.Debug(lambdaErr)
					panic("args type must map[string]string")
				}
				resSymbolOne := self.GetSymbol(argsSymbol)
				lambdaArgsDl.Symbols[argsFormal] = resSymbolOne
			}
		}
		lambdaFunc, ok := Lambdas[lambda]
		if !ok {
			panic("lambda '" + lambda + "' not found")
		}
		resI = lambdaFunc(lambdaArgsDl)

		return
	}
	return
}

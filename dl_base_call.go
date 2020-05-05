package dl

import (
	log "github.com/sirupsen/logrus"
)

func (self *Dl) setCall() {
	Lambdas["call"] = func(self *Dl) (resI interface{}) {
		// log.Debug("in Call", self.TmpInterface)
		self.CheckLambdasNameForce("call")
		var err error
		var lambda string
		if len(self.SubNodeTree) >= 2 {
			if lambda, err = self.SubNodeGetSingleString("lambda"); err != nil {
				panic("'lambda' not found")
			}
		} else if len(self.SubNodeList) >= 2 {
			if lambda, err = self.SubNodeListGetSingleString(1); err != nil {
				panic("'lambda' not found")
			}
		} else {
			panic("'call' format error")
		}
		lambdaFunc, ok := Lambdas[lambda]
		if !ok {
			panic("lambda '" + lambda + "' not found")
		}

		// args传入参数，给lambda所在作用域添加symbol参数
		var args *Dl
		if len(self.SubNodeTree) >= 3 {
			if args, err = self.SubNodeGet("args"); err != nil {
				args = nil
			}
		} else if len(self.SubNodeList) == 3 {
			if args, err = self.SubNodeListGet(2); err != nil {
				args = nil
			}
		} else {
			args = nil
		}

		if args != nil {
			if len(args.SubNodeList) != 0 {
				panic("call 'args' type must object")
			}
			log.Debug(args.SubNodeTree)
			for argsSymbol, argsValue := range args.SubNodeTree {
				args.Symbols[argsSymbol] = argsValue.Call()
			}
		}

		resI = lambdaFunc(args)

		return
	}
	return
}

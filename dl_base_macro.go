package dl

func (self *Dl) setMacro() {
	Lambdas["macro"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasNameForce("macro")

		if len(self.SubNodeTree) != 0 {
			self.SubNodeTree["name"].TmpInterface = "lambda"
		} else if len(self.SubNodeList) != 0 {
			self.SubNodeList[0].TmpInterface = "lambda"
		}

		var err error
		var symbol string
		if symbol, err = self.SubNodeGetSingleString("symbol"); err != nil {
			panic(err)
		}

		lambdaFunc, ok := Lambdas["lambda"]
		if !ok {
			panic("lambda not found")
		}
		lambdaName := lambdaFunc(self)

		self.FatherNode.Symbols[symbol] = lambdaName
		resI = lambdaName
		return
	}
	return
}

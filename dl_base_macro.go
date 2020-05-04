package dl

func (self *Dl) setMacro() {
	Lambdas["macro"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasNameForce("macro")
		// macro实体是lambda
		// lambda操作比较复杂，只允许object模型
		if len(self.SubNodeList) != 0 {
			panic("'lambda' must be object type")
		}
		self.SubNodeTree["name"].TmpInterface = "lambda"

		var err error
		var symbol string
		if symbol, err = self.SubNodeGetSingleString("symbol"); err != nil {
			err = nil
			symbol = ""
		}

		lambdaFunc, ok := Lambdas["lambda"]
		if !ok {
			panic("lambda not found")
		}
		lambdaName := lambdaFunc(self)

		if symbol != "" {
			self.FatherNode.Symbols[symbol] = lambdaName
		} else {
			resI = lambdaName
		}
		return
	}
	return
}

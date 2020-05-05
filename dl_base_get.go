package dl

func (self *Dl) setGet() {
	Lambdas["get"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasNameForce("get")
		var err error
		var symbol string
		if len(self.SubNodeTree) >= 2 {
			if symbol, err = self.SubNodeGetSingleString("symbol"); err != nil {
				panic("'symbol' not found")
			}
		} else if len(self.SubNodeList) == 2 {
			if symbol, err = self.SubNodeListGetSingleString(1); err != nil {
				panic("'symbol' not found")
			}
		} else {
			panic("'get' format error")
		}
		resI = self.GetSymbol(symbol)

		return
	}
}

package dl

func (self *Dl) setGet() {
	Lambdas["get"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasNameForce("get")
		var err error
		var symbol string
		if symbol, err = self.SubNodeGetSingleString("symbol"); err != nil {
			panic("'symbol' not found")
		}
		resI = self.GetSymbol(symbol)

		return
	}
}

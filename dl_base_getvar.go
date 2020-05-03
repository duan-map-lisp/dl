package dl

func (self *Dl) setGetvar() {
	Lambdas["getvar"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasName("getvar")
		var err error
		var symbol string
		if symbol, err = self.SubNodeGetSingleString("symbol"); err != nil {
			panic("'symbol' not found")
		}
		resI = self.GetSymbol(symbol)

		return
	}
}

package dl

func checkQuote(self *Dl) (res bool) {
	if len(self.SubNodeTree) != 0 {
		res = false
		return
	}
	if len(self.SubNodeList) != 0 {
		res = false
		return
	}
	res = true

	return
}

func (self *Dl) setQuote() {
	Lambdas["quote"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasNameForce("quote")
		var err error
		var data *Dl

		if len(self.SubNodeTree) >= 2 {
			if data, err = self.SubNodeGet("data"); err != nil {
				panic("'data' not found")
			}
		} else if len(self.SubNodeList) == 2 {
			if data, err = self.SubNodeListGet(1); err != nil {
				panic("'data' not found")
			}
		} else {
			panic("quote format error")
		}

		resI = data
		return
	}
}

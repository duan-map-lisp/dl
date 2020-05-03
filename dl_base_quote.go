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

		if len(self.SubNodeTree) != 0 {
			if data, err = self.SubNodeGet("data"); err != nil {
				panic(err)
			}
			if checkQuote(data) {
				resI = data.Call()
			} else {
				resI = data
			}
			return
		}

		if len(self.SubNodeList) == 2 {
			if data, err = self.SubNodeListGet(1); err != nil {
				panic(err)
			}
			if checkQuote(data) {
				resI = data.Call()
			} else {
				resI = data
			}
			return
		}

		panic("quote format error")
	}
}

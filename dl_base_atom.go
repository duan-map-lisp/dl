package dl

func checkAtom(self *Dl) (res bool) {
	tmpData := self.Call()
	switch tmpRes := tmpData.(type) {
	case *Dl:
		if len(tmpRes.SubNodeTree) != 0 {
			res = false
			return
		}
		if len(tmpRes.SubNodeList) != 0 {
			res = false
			return
		}
	default:
		res = true
	}

	return
}

func (self *Dl) setAtom() {
	Lambdas["atom"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasNameForce("atom")
		var err error
		var data *Dl

		if len(self.SubNodeTree) != 0 {
			if data, err = self.SubNodeGet("data"); err != nil {
				panic(err)
			}
			resI = checkAtom(data)
			return
		}

		if len(self.SubNodeList) == 2 {
			if data, err = self.SubNodeListGet(1); err != nil {
				panic(err)
			}
			resI = checkAtom(data)
			return
		}

		panic("atom format error")
	}
}

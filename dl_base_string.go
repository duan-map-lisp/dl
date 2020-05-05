package dl

func (self *Dl) setString() {
	Lambdas["string"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasNameForce("string")
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
			panic("'get' format error")
		}

		if data == nil {
			resI = "null"
		} else {
			tmpData := data.Call()
			switch dataTmp := tmpData.(type) {
			case *Dl:
				resI = dataTmp.String()
			default:
				resI = dataTmp
			}
		}

		return
	}
}

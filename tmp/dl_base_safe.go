package dl

func (self *Dl) setSafe() {
	Lambdas["safe"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasNameForce("safe")

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
			panic("'safe' format error")
		}

		if len(self.FatherNode.SubNodeTree) != 0 {
			self.FatherNode.SubNodeTree[self.NodeName] = data
		}
		if len(self.FatherNode.SubNodeList) != 0 {
			self.FatherNode.SubNodeList[self.NodeIndex] = data
		}
		return
	}
	return
}

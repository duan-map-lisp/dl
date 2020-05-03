package dl

func (self *Dl) setSafe() {
	Lambdas["safe"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasNameForce("safe")

		var err error
		var data *Dl
		if data, err = self.SubNodeGet("data"); err != nil {
			panic(err)
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

package dl

func (self *Dl) setBlock() {
	self.Lambdas["block"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasName("block")
		for _, resOne := range self.SubNodeList {
			self.BlockBreakFlag = false
			resI = resOne.Run()
			if self.BlockBreakFlag {
				break
			}
		}
		return
	}
}

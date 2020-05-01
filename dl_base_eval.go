package dl

func (self *Dl) setEval() {
	self.Lambdas["eval"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasName("eval")
		var err error
		var data []byte
		if data, err = self.SubNodeGetSliceByte("data"); err != nil {
			panic("'data' not found")
		}

		resI = (&Dl{
			FatherNode:  self,
			AllStr:      data,
			SubNodeTree: map[string]*Dl{},
			Lambdas:     map[string]func(*Dl) interface{}{},
			Symbols:     map[string]interface{}{},
		}).Eval()
		return
	}
}

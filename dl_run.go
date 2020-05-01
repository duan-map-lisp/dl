package dl

import (
// "fmt"
)

func (self *Dl) Run() (resI interface{}) {
	// fmt.Println ("in Run", self.TmpInterface)
	if self.TmpInterface != nil {
		resI = self.TmpInterface
		return
	}

	if len(self.TmpMap) != 0 {
		resI = self.LambdaMap()
		return
	}

	if len(self.TmpList) != 0 {
		resI = self.LambdaList()
		return
	}

	return
}

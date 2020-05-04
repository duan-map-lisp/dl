package dl

func CheckEq(src *Dl, dest *Dl) (ok bool) {
	srcType := src.CheckType()
	destType := dest.CheckType()
	if srcType != destType {
		ok = false
		return
	}
	if srcType == "object" || srcType == "list" {
		ok = false
		return
	}
	if src.Call() != dest.Call() {
		ok = false
		return
	}
	ok = true
	return
}

func (self *Dl) setEq() {
	Lambdas["eq"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasNameForce("eq")
		// eq无上限，只允许list模型
		if len(self.SubNodeTree) != 0 {
			panic("'eq' must be list type")
		}
		if len(self.SubNodeList) < 3 {
			panic("'eq' symbol must more than 3")
		}
		var firstNode *Dl
		for resIndex, resOne := range self.SubNodeList {
			if resIndex == 0 {
				continue
			}
			if resIndex == 1 {
				firstNode = resOne
				continue
			}
			if CheckEq(firstNode, resOne) {
				resI = true
				continue
			} else {
				resI = false
				return
			}
		}
		return
	}
}

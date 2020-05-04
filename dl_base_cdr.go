package dl

func (self *Dl) setCdr() {
	Lambdas["cdr"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasNameForce("cdr")
		var err error
		var data *Dl
		var key string
		if len(self.SubNodeTree) >= 2 {
			if data, err = self.SubNodeGet("data"); err != nil {
				panic("'data' not found")
			}
			if data.CheckType() == "list" {
				tmpDataList := data.Call().(*Dl)
				if len(tmpDataList.SubNodeList) <= 1 {
					resI = nil
				} else {
					resI = tmpDataList.SubNodeList[1:]
				}
			} else if data.CheckType() == "object" {
				if key, err = self.SubNodeGetSingleString("key"); err != nil {
					panic("'key' not found")
				}
				tmpDataObject := data.Call().(*Dl)
				delete(tmpDataObject.SubNodeTree, key)
				resI = tmpDataObject
				return
			}
		} else if len(self.SubNodeList) == 2 {
			if data, err = self.SubNodeListGet(1); err != nil {
				panic("'data' not found")
			}
			if data.CheckType() == "list" {
				tmpDataList := data.Call().(*Dl)
				if len(tmpDataList.SubNodeList) <= 1 {
					resI = nil
				} else {
					resI = tmpDataList.SubNodeList[1:]
				}
			} else {
				panic("'cdr' format error")
			}
		} else if len(self.SubNodeList) == 3 {
			if data, err = self.SubNodeListGet(1); err != nil {
				panic("'data' not found")
			}
			if data.CheckType() == "object" {
				if key, err = self.SubNodeListGetSingleString(2); err != nil {
					panic("'key' not found")
				}
				tmpDataObject := data.Call().(*Dl)
				delete(tmpDataObject.SubNodeTree, key)
				resI = tmpDataObject
				return
			} else {
				panic("'cdr' format error")
			}
		} else {
			panic("'cdr' format error")
		}

		return
	}
}

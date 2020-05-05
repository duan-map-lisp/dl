package dl

func (self *Dl) setAppend() {
	Lambdas["append"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasNameForce("append")
		var err error
		var insert *Dl
		var data *Dl
		var key string
		if len(self.SubNodeTree) >= 3 {
			if insert, err = self.SubNodeGet("insert"); err != nil {
				panic("'insert' not found")
			}
			if data, err = self.SubNodeGet("data"); err != nil {
				panic("'data' not found")
			}
			if data.CheckType() == "list" || data.CheckType() == "null" {
				tmpDataList := data.Call().(*Dl)
				tmpRes := &Dl{
					FatherNode: self,
				}
				tmpRes.Init()
				tmpInsert := insert.Call()
				switch tmpInsertRes := tmpInsert.(type) {
				case *Dl:
					tmpRes.SubNodeList = append(tmpDataList.SubNodeList, tmpInsertRes)
				default:
					tmpRes.SubNodeList = append(tmpDataList.SubNodeList, &Dl{
						FatherNode:   self,
						TmpInterface: tmpInsertRes,
					})
				}
				resI = tmpRes
				return
			} else if data.CheckType() == "object" || data.CheckType() == "null" {
				if key, err = self.SubNodeGetSingleString("key"); err != nil {
					panic("'key' not found")
				}
				tmpDataObject := data.Call().(*Dl)
				tmpRes := &Dl{
					FatherNode: self,
				}
				tmpRes.Init()
				tmpRes.SubNodeTree = tmpDataObject.SubNodeTree
				tmpRes.SubNodeTree[key] = insert
				resI = tmpRes
				return
			}
		} else if len(self.SubNodeList) == 3 {
			if insert, err = self.SubNodeListGet(1); err != nil {
				panic("'insert' not found")
			}
			if data, err = self.SubNodeListGet(2); err != nil {
				panic("'data' not found")
			}
			if data.CheckType() == "list" || data.CheckType() == "null" {
				tmpDataList := data.Call().(*Dl)
				tmpRes := &Dl{
					FatherNode: self,
				}
				tmpRes.Init()
				tmpInsert := insert.Call()
				switch tmpInsertRes := tmpInsert.(type) {
				case *Dl:
					tmpRes.SubNodeList = append(tmpDataList.SubNodeList, tmpInsertRes)
				default:
					tmpRes.SubNodeList = append(tmpDataList.SubNodeList, &Dl{
						FatherNode:   self,
						TmpInterface: tmpInsertRes,
					})
				}
				resI = tmpRes
				return
			} else {
				panic("'append' format error")
			}
		} else if len(self.SubNodeList) == 4 {
			if data, err = self.SubNodeListGet(1); err != nil {
				panic("'data' not found")
			}
			if data.CheckType() == "object" || data.CheckType() == "null" {
				if key, err = self.SubNodeListGetSingleString(3); err != nil {
					panic("'key' not found")
				}
				tmpDataObject := data.Call().(*Dl)
				tmpRes := &Dl{
					FatherNode: self,
				}
				tmpRes.Init()
				tmpRes.SubNodeTree = tmpDataObject.SubNodeTree
				tmpRes.SubNodeTree[key] = insert
				resI = tmpRes
				return
			} else {
				panic("'append' format error")
			}
		} else {
			panic("'append' format error")
		}

		return
	}
}

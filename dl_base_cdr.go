package dl

func CopyMap(src map[string]*Dl, dest map[string]*Dl) {
	for key, value := range src {
		dest[key] = value
	}
}

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
			if data.CheckType() == "list" || data.CheckType() == "null" {
				tmpDataList := data.Call().(*Dl)
				if len(tmpDataList.SubNodeList) <= 1 {
					resI = nil
				} else {
					tmpNode := &Dl{
						FatherNode: tmpDataList.FatherNode,
					}
					tmpNode.Init()
					tmpNode.SubNodeList = tmpDataList.SubNodeList[1:]
					resI = tmpNode
				}
			} else if data.CheckType() == "object" || data.CheckType() == "null" {
				if key, err = self.SubNodeGetSingleString("key"); err != nil {
					panic("'key' not found")
				}
				tmpDataObject := data.Call().(*Dl)
				tmpNodeMap := make(map[string]*Dl)
				CopyMap(tmpDataObject.SubNodeTree, tmpNodeMap)

				delete(tmpNodeMap, key)
				tmpNode := &Dl{
					FatherNode: tmpDataObject.FatherNode,
				}
				tmpNode.Init()
				tmpNode.SubNodeTree = tmpNodeMap
				resI = tmpNode
				return
			}
		} else if len(self.SubNodeList) == 2 {
			if data, err = self.SubNodeListGet(1); err != nil {
				panic("'data' not found")
			}
			if data.CheckType() == "list" || data.CheckType() == "null" {
				tmpDataList := data.Call().(*Dl)
				if len(tmpDataList.SubNodeList) <= 1 {
					resI = nil
				} else {
					tmpNode := &Dl{
						FatherNode: tmpDataList.FatherNode,
					}
					tmpNode.Init()
					tmpNode.SubNodeList = tmpDataList.SubNodeList[1:]
					resI = tmpNode
				}
			} else {
				panic("'cdr' format error")
			}
		} else if len(self.SubNodeList) == 3 {
			if data, err = self.SubNodeListGet(1); err != nil {
				panic("'data' not found")
			}
			if data.CheckType() == "object" || data.CheckType() == "null" {
				if key, err = self.SubNodeListGetSingleString(2); err != nil {
					panic("'key' not found")
				}
				tmpDataObject := data.Call().(*Dl)
				tmpNodeMap := make(map[string]*Dl)
				CopyMap(tmpDataObject.SubNodeTree, tmpNodeMap)

				delete(tmpNodeMap, key)
				tmpNode := &Dl{
					FatherNode: tmpDataObject.FatherNode,
				}
				tmpNode.Init()
				tmpNode.SubNodeTree = tmpNodeMap
				resI = tmpNode
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

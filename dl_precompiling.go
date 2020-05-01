package dl

import (
	"fmt"

	"encoding/json"
)

func (self *Dl) Precompiling () {
	var err error
	// 如果是list
	if err = json.Unmarshal (self.AllStr, &self.TmpList); err == nil {
		self.AllStr = []byte ("")
		if len (self.TmpList) < 1 {
			return
		}
		for tmpNum, tmpOne := range self.TmpList {
			data := Dl {
				NodeIndex: tmpNum,
				FatherNode: self,
				AllStr: tmpOne,
				SubNodeTree: map[string]*Dl {},
				Lambdas: map[string]func (*Dl) (interface{}) {},
			}
			fmt.Println ("in Precompiling ", data.NodeName, self.SubNodeTree)
			self.SubNodeList = append (self.SubNodeList, &data)
			(&data).Precompiling ()
		}

		return
	}

	// 如果是map
	if err = json.Unmarshal (self.AllStr, &self.TmpMap); err == nil {
		fmt.Println ("get map", self.TmpMap)
		self.AllStr = []byte ("")
		if len (self.TmpMap) < 1 {
			return
		}
		for tmpKey, tmpOne := range self.TmpMap {
			fmt.Println ("tmpKey:", tmpKey)
			fmt.Println ("tmpOne", string (tmpOne))
			data := Dl {
				NodeName: tmpKey,
				FatherNode: self,
				AllStr: tmpOne,
				SubNodeTree: map[string]*Dl {},
				Lambdas: map[string]func (*Dl) (interface{}) {},
			}
			self.SubNodeTree[tmpKey] = &data
			(&data).Precompiling ()
			fmt.Println (self.SubNodeTree)
		}
		return
	}

	if err = json.Unmarshal (self.AllStr, &self.TmpInterface); err == nil {
		self.AllStr = []byte ("")
		return
	}

	panic (err)
}

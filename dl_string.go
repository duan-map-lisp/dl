package dl

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func (self *Dl) String() (res string) {
	tmp, _ := json.Marshal(self)
	res = string(tmp)
	return
}

func (self *Dl) StringNodeTree() (res string) {
	if self.FatherNode != nil {
		switch tmpType := self.FatherNode.DataInterface.(type) {
		case []*Dl:
			for index, value := range tmpType {
				if index == 0 {
					res += "\nGet NodeIndex '0': "
					res += fmt.Sprint(value.String())
				}
				if value != self {
					continue
				}
				res += "\nNodeIndex: "
				res += strconv.Itoa(index)
				if value.FatherNode != nil {
					res += self.FatherNode.StringNodeTree()
				} else {
					res += "\nGet NodeName: root\n"
				}
				return
			}
		case map[string]*Dl:
			for key, value := range tmpType {
				if key == "name" {
					res += "\nGet NodeName 'name': "
					res += fmt.Sprint(value.String())
				}
				if value != self {
					continue
				}
				res += "\nNodeName: "
				res += key
				if value.FatherNode != nil {
					res += self.FatherNode.StringNodeTree()
				} else {
					res += "\nGet NodeName: root\n"
				}
				return
			}
		default:
			panic("???")
		}

		res += self.FatherNode.StringNodeTree()
	} else {
		res += "\nGet NodeName: root\n"
		return
	}
	return
}

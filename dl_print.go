package dl

import (
	"encoding/json"
)

func AddStringObject(root map[string]interface{}, key string, value *Dl) (res map[string]interface{}) {
	if len(value.SubNodeTree) > 0 {
		node := map[string]interface{}{}
		for nodeKey, nodeValue := range value.SubNodeTree {
			AddStringObject(node, nodeKey, nodeValue)
		}
		root[key] = &node
	} else if len(value.SubNodeList) > 0 {
		node := []interface{}{}
		for _, value := range value.SubNodeList {
			AddStringList(node, value)
		}
		root[key] = &node
	} else {
		root[key] = value.TmpInterface
	}
	res = root
	return
}

func AddStringList(root []interface{}, value *Dl) (res []interface{}) {
	if len(value.SubNodeTree) > 0 {
		node := map[string]interface{}{}
		for nodeKey, nodeValue := range value.SubNodeTree {
			AddStringObject(node, nodeKey, nodeValue)
		}
		root = append(root, &node)
	} else if len(value.SubNodeList) > 0 {
		node := []interface{}{}
		for _, value := range value.SubNodeList {
			AddStringList(node, value)
		}
		root = append(root, &node)
	} else {
		root = append(root, value.TmpInterface)
	}
	res = root

	return
}

func (self *Dl) String() (res string) {
	var rootI interface{}
	if len(self.SubNodeTree) > 0 {
		root := map[string]interface{}{}
		for key, value := range self.SubNodeTree {
			root = AddStringObject(root, key, value)
		}
		rootI = root
	} else if len(self.SubNodeList) > 0 {
		root := []interface{}{}
		for _, value := range self.SubNodeList {
			root = AddStringList(root, value)
		}
		rootI = root
	} else {
		rootI = self.TmpInterface
	}

	if resByte, err := json.Marshal(rootI); err != nil {
		panic(err)
	} else {
		res = string(resByte)
	}
	return
}

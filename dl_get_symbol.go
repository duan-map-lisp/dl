package dl

import (
// log "github.com/sirupsen/logrus"
)

func (self *Dl) GetSymbol(symbolName string) (res interface{}) {
	var ok bool
	res, ok = self.Symbols[symbolName]
	if !ok {
		if self.FatherNode == nil {
			res = nil
		} else {
			res = self.FatherNode.GetSymbol(symbolName)
		}
	}

	return
}

// 查找symbol所在的节点
func (self *Dl) GetSymbolNode(symbolName string) (resNode *Dl) {
	_, ok := self.Symbols[symbolName]
	if !ok {
		if self.FatherNode == nil {
			resNode = nil
		} else {
			resNode = self.FatherNode.GetSymbolNode(symbolName)
		}
	} else {
		resNode = self
	}

	return
}

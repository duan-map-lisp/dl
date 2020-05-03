package dl

import (
	log "github.com/sirupsen/logrus"
)

func (self *Dl) GetSymbol(symbolName string) (res interface{}) {
	log.Debug("in get symbol", self.Symbols)
	res, ok := self.Symbols[symbolName]
	if !ok {
		if self.FatherNode == nil {
			panic("undefind symbol:" + symbolName)
		} else {
			res = self.FatherNode.GetSymbol(symbolName)
		}
	}

	return
}

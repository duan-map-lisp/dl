package dl

import (
	log "github.com/sirupsen/logrus"
)

func (self *Dl) setDel() {
	Lambdas["del"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasNameForce("del")
		var err error
		var symbol string
		if len(self.SubNodeTree) >= 2 {
			if symbol, err = self.SubNodeGetSingleString("symbol"); err != nil {
				panic("'symbol' not found")
			}
		} else if len(self.SubNodeList) == 2 {
			if symbol, err = self.SubNodeListGetSingleString(1); err != nil {
				panic("'symbol' not found")
			}
		} else {
			panic("'del' format error")
		}

		resNode := self.GetSymbolNode(symbol)
		if resNode == nil {
			return
		}
		log.Debug("del resNode ", resNode.String())
		log.Debug("del resNode Symbols ", resNode.Symbols)
		delete(resNode.Symbols, symbol)

		return
	}
}

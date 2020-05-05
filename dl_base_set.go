package dl

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

func (self *Dl) setSet() {
	Lambdas["set"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasNameForce("set")
		var err error
		var symbol string
		var value *Dl
		if len(self.SubNodeTree) >= 3 {
			if symbol, err = self.SubNodeGetSingleString("symbol"); err != nil {
				panic("'symbol' not found")
			}
			if value, err = self.SubNodeGet("value"); err != nil {
				panic("'value' not found")
			}
		} else if len(self.SubNodeList) == 3 {
			if symbol, err = self.SubNodeListGetSingleString(1); err != nil {
				panic("'symbol' not found")
			}
			if value, err = self.SubNodeListGet(2); err != nil {
				panic("'value' not found")
			}
		} else {
			panic("'set' format error")
		}

		resNode := self.GetSymbolNode(symbol)
		log.Debug("set resNode ", resNode.String())
		log.Debug("set resNode Symbols ", resNode.Symbols)
		resData := resNode.GetSymbol(symbol)
		if checkQuote(value) {
			if err = json.Unmarshal(value.AllStr, &resData); err != nil {
				panic(err)
			}
			resNode.Symbols[symbol] = resData
		} else {
			panic("'set' must use base data type")
		}

		return
	}
}

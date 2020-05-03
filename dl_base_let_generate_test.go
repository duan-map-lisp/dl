package dl

import (
	log "github.com/sirupsen/logrus"
)

func (self *Dl) setLet() {
	Lambdas["let"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasNameForce("let")
		if self.FatherNode == nil {
			panic("let不可能是root节点")
		}

		var err error
		var symbol string
		if symbol, err = self.SubNodeGetSingleString("symbol"); err != nil {
			panic("'symbol' not found")
		}
		var symbol_type string
		if symbol_type, err = self.SubNodeGetSingleString("type"); err != nil {
			panic("'type' not found")
		}
		switch symbol_type {

			{{$all_types := MkSlice "int8" "int16" "int32" "int64" "uint8" "uint16" "uint32" "uint64" "int" "uint" "rune" "byte" "uintptr" "string"}}
			{{/* $format_types := MkSlice "single" "slice" "map" */}}
			{{$format_types := MkSlice "single"}}
			{{range $_, $type_base := $all_types}}
			{{range $_, $format_one := $format_types}}
			{{$type_one := GetFormat $type_base $format_one}}

		case "{{$type_one}}":
			if _, ok := self.FatherNode.Symbols[symbol]; ok {
				panic ("redefine val " + symbol)
			}
			if self.FatherNode.Symbols[symbol], err = self.SubNodeGet{{CoverSnakeCaseToPascalCase $format_one}}{{CoverSnakeCaseToPascalCase $type_base}} ("value"); err != nil {
				panic (err)
			}

			{{end}}
			{{end}}
		case "lambda":
			if _, ok := self.FatherNode.Symbols[symbol]; ok {
				panic ("redefine val " + symbol)
			}
			if self.FatherNode.Symbols[symbol], err = self.SubNodeGetSingleString ("value"); err != nil {
				panic (err)
			}
		case "array":
			if _, ok := self.FatherNode.Symbols[symbol]; ok {
				panic ("redefine val " + symbol)
			}
			var tmpNode *Dl
			if tmpNode, err = self.SubNodeGet ("value"); err != nil {
				panic (err)
			}
			if len (tmpNode.TmpMap) != 0 {
				panic ("value not array type")
			}
			self.FatherNode.Symbols[symbol] = tmpNode
		case "object":
			if _, ok := self.FatherNode.Symbols[symbol]; ok {
				panic ("redefine val " + symbol)
			}
			var tmpNode *Dl
			if tmpNode, err = self.SubNodeGet ("value"); err != nil {
				panic (err)
			}
			if len (tmpNode.TmpList) != 0 {
				panic ("value not object type")
			}
			self.FatherNode.Symbols[symbol] = tmpNode

		default:
			panic ("error var type" + symbol_type)
		}
		log.Debug ("self.FatherNode.Symbols", self.FatherNode.Symbols)
		return
	}
	return
}

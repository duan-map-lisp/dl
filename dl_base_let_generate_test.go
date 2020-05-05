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
		var symbol_type string
		if len(self.SubNodeTree) >= 4 {
			if symbol, err = self.SubNodeGetSingleString("symbol"); err != nil {
				panic("'symbol' not found")
			}
			if symbol_type, err = self.SubNodeGetSingleString("type"); err != nil {
				panic("'type' not found")
			}
		} else if len(self.SubNodeList) == 4 {
			if symbol, err = self.SubNodeListGetSingleString(1); err != nil {
				panic("'symbol' not found")
			}
			if symbol_type, err = self.SubNodeListGetSingleString(2); err != nil {
				panic("'type' not found")
			}
		} else {
			panic("'let' format error")
		}

		switch symbol_type {

			{{$all_types := MkSlice "int8" "int16" "int32" "int64" "uint8" "uint16" "uint32" "uint64" "float32" "float64" "int" "uint" "rune" "byte" "uintptr" "string"}}
			{{/* $format_types := MkSlice "single" "slice" "map" */}}
			{{$format_types := MkSlice "single"}}
			{{range $_, $type_base := $all_types}}
			{{range $_, $format_one := $format_types}}
			{{$type_one := GetFormat $type_base $format_one}}

		case "{{$type_one}}":
			if _, ok := self.FatherNode.Symbols[symbol]; ok {
				panic ("redefine val " + symbol)
			}
			if len (self.SubNodeTree) >= 4 {
				if self.FatherNode.Symbols[symbol], err = self.SubNodeGet{{CoverSnakeCaseToPascalCase $format_one}}{{CoverSnakeCaseToPascalCase $type_base}} ("value"); err != nil {
					panic (err)
				}
			} else if len (self.SubNodeList) == 4 {
				if self.FatherNode.Symbols[symbol], err = self.SubNodeListGet{{CoverSnakeCaseToPascalCase $format_one}}{{CoverSnakeCaseToPascalCase $type_base}} (3); err != nil {
					panic (err)
				}
			} else {
				panic("'let' format error")
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
		case "list":
			if _, ok := self.FatherNode.Symbols[symbol]; ok {
				panic ("redefine val " + symbol)
			}
			var tmpNode *Dl
			if len (self.SubNodeTree) >= 4 {
				if tmpNode, err = self.SubNodeGet ("value"); err != nil {
					panic (err)
				}
			} else if len (self.SubNodeList) == 4 {
				if tmpNode, err = self.SubNodeListGet (3); err != nil {
					panic (err)
				}
			} else {
				panic("'let' format error")
			}
			tmpRes := tmpNode.Call ()
			switch resTmp := tmpRes.(type) {
			case *Dl:
				if len (resTmp.SubNodeTree) != 0 {
					panic ("value not object type")
				}
			default:
				panic ("value not object type")
			}
			self.FatherNode.Symbols[symbol] = tmpRes
		case "object":
			if _, ok := self.FatherNode.Symbols[symbol]; ok {
				panic ("redefine val " + symbol)
			}
			var tmpNode *Dl
			if len (self.SubNodeTree) >= 4 {
				if tmpNode, err = self.SubNodeGet ("value"); err != nil {
					panic (err)
				}
			} else if len (self.SubNodeList) == 4 {
				if tmpNode, err = self.SubNodeListGet (3); err != nil {
					panic (err)
				}
			} else {
				panic("'let' format error")
			}

			tmpRes := tmpNode.Call ()
			switch resTmp := tmpRes.(type) {
			case *Dl:
				if len (resTmp.SubNodeList) != 0 {
					panic ("value not object type")
				}
			default:
				panic ("value not object type")
			}
			self.FatherNode.Symbols[symbol] = tmpRes

		default:
			panic ("error var type " + symbol_type)
		}
		log.Debug ("self.FatherNode.Symbols ", self.FatherNode.Symbols)
		return
	}
	return
}

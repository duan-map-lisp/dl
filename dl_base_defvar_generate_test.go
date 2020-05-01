package dl

import (
	"fmt"
)

func (self *Dl) setDefvar() {
	self.Lambdas["defvar"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasName("defvar")
		if self.FatherNode == nil {
			panic("defvar不可能是root节点")
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
			{{$format_types := MkSlice "single" "slice" "map"}}
			{{range $_, $type_base := $all_types}}
			{{range $_, $format_one := $format_types}}
			{{$type_one := GetFormat $type_base $format_one}}

		case "{{$type_one}}":
			if self.FatherNode.Symbols[symbol], err = self.SubNodeGet{{CoverSnakeCaseToPascalCase $format_one}}{{CoverSnakeCaseToPascalCase $type_base}} ("value"); err != nil {
				panic (err)
			}

			{{end}}
			{{end}}

		default:
			panic ("error var type" + symbol_type)
		}
		fmt.Println ("self.FatherNode.Symbols", self.FatherNode.Symbols)
		return
	}
	return
}

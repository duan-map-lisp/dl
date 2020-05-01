package dl

import (
	"strconv"
	"encoding/json"
)

{{$all_types := MkSlice "int8" "int16" "int32" "int64" "uint8" "uint16" "uint32" "uint64" "int" "uint" "rune" "byte" "uintptr" "string"}}
{{$format_types := MkSlice "single" "slice" "map"}}
{{range $_, $type_base := $all_types}}
{{range $_, $format_one := $format_types}}
{{$type_one := GetFormat $type_base $format_one}}

func (self *Dl) SubNodeGet{{CoverSnakeCaseToPascalCase $format_one}}{{CoverSnakeCaseToPascalCase $type_base}} (key string) (res {{$type_one}}, err error) {
	var value *Dl
	if value, err = self.SubNodeCheckKey (key); err != nil {
		return
	}
	resI := value.Run ()
	if resI == nil {
		panic ("get value is nil")
	}
	switch resTmp := resI.(type) {
	case {{$type_one}}:
		res = resTmp
		return
	}

	var tmpRes {{$type_one}}
	if err = json.Unmarshal(value.AllStr, &tmpRes); err != nil {
		panic (key + " type not {{$type_one}}: " + err.Error ())
	}
	res = tmpRes
	return
}

func (self *Dl) SubNodeListGet{{CoverSnakeCaseToPascalCase $format_one}}{{CoverSnakeCaseToPascalCase $type_base}} (index int) (res {{$type_one}}, err error) {
	if len(self.SubNodeList) <= index {
		panic ("index out range")
	}

	value := self.SubNodeList[index]
	if resI := value.Run (); resI == nil {
		panic ("get value is nil")
	}
	var tmpRes {{$type_one}}
	if err = json.Unmarshal(value.AllStr, &tmpRes); err != nil {
		panic (strconv.Itoa (index) + " type not {{$type_one}}: " + err.Error ())
	}
	res = tmpRes
	return
}

{{end}}
{{end}}

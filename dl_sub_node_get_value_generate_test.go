package dl

import (
	"errors"
	"encoding/json"
)

func (self *Dl) SubNodeCheckKey(key string) (value *Dl, err error) {
	var ok bool
	if value, ok = self.SubNodeTree[key]; !ok {
		err = errors.New("sub node key " + key + " not found")
		return
	}
	return
}

{{$all_types := MkSlice "int8" "int16" "int32" "int64" "uint8" "uint16" "uint32" "uint64" "float32" "float64" "int" "uint" "rune" "byte" "uintptr" "string" "bool"}}
{{/* $format_types := MkSlice "single" "slice" "map" */}}
{{$format_types := MkSlice "single"}}
{{range $_, $type_base := $all_types}}
{{range $_, $format_one := $format_types}}
{{$type_one := GetFormat $type_base $format_one}}

func (self *Dl) SubNodeGet{{CoverSnakeCaseToPascalCase $format_one}}{{CoverSnakeCaseToPascalCase $type_base}} (key string) (res {{$type_one}}, err error) {
	var value *Dl
	if value, err = self.SubNodeCheckKey (key); err != nil {
		return
	}
	resI := value.Call ()
	if resI == nil {
		err = errors.New ("get value is nil")
		return
	}

	for {
		switch resTmp := resI.(type) {
		case *Dl:
			value = resTmp
			resI = value.Call ()
			continue
		case {{$type_one}}:
			res = resTmp
			return
		default:
			break
		}
		break
	}

	var tmpRes {{$type_one}}
	if err = json.Unmarshal (value.AllStr, &tmpRes); err != nil {
		return
	}
	res = tmpRes
	return
}

func (self *Dl) SubNodeListGet{{CoverSnakeCaseToPascalCase $format_one}}{{CoverSnakeCaseToPascalCase $type_base}} (index int) (res {{$type_one}}, err error) {
	if len(self.SubNodeList) <= index {
		err = errors.New ("index out range")
		return
	}

	value := self.SubNodeList[index]
	resI := value.Call ()
	if resI == nil {
		err = errors.New ("get value is nil")
		return
	}

	for {
		switch resTmp := resI.(type) {
		case *Dl:
			value = resTmp
			resI = value.Call ()
			continue
		case {{$type_one}}:
			res = resTmp
			return
		default:
			break
		}
		break
	}

	var tmpRes {{$type_one}}
	if err = json.Unmarshal (value.AllStr, &tmpRes); err != nil {
		return
	}
	res = tmpRes
	return
}

{{end}}
{{end}}

func (self *Dl) SubNodeGet (key string) (res *Dl, err error) {
	if res, err = self.SubNodeCheckKey (key); err != nil {
		return
	}
	return
}

func (self *Dl) SubNodeListGet (index int) (res *Dl, err error) {
	if len(self.SubNodeList) <= index {
		err = errors.New ("index out range")
		return
	}

	res = self.SubNodeList[index]
	return
}

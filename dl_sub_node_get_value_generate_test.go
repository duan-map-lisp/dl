package dl

import (
	"errors"
	"encoding/json"
)

func (self *Dl) SubNodeCheckKey(key string) (value *Dl, err error) {
	switch tmpType := self.DataInterface.(type) {
	case map[string]*Dl:
		var ok bool
		if value, ok = tmpType[key]; !ok {
			err = errors.New("sub node key " + key + " not found")
			return
		}
	default:
		err = errors.New("type not object:" + self.String ())
	}
	return
}

func (self *Dl) SubNodeCheckIndex (index int) (value *Dl, err error) {
	switch tmpType := self.DataInterface.(type) {
	case []*Dl:
		if len (tmpType) <= index {
			err = errors.New ("index out range")
			return
		}
		value = tmpType[index]
	default:
		err = errors.New("type not list:" + self.String ())
	}
	return
}

// {{$all_types := MkSlice "int8" "int16" "int32" "int64" "uint8" "uint16" "uint32" "uint64" "float32" "float64" "int" "uint" "rune" "byte" "uintptr" "string" "bool"}}

{{$all_types := MkSlice "int8" "int16" "int32" "int64" "uint8" "uint16" "uint32" "uint64" "float32" "float64" "int" "uint" "rune" "byte" "uintptr" "string" "bool"}}
{{range $_, $type_one := $all_types}}
func (self *Dl) SubNodeGet{{CoverSnakeCaseToPascalCase $type_one}} (key string) (res {{$type_one}}, err error) {
	var value *Dl
	if value, err = self.SubNodeCheckKey (key); err != nil {
		return
	}
	if value == nil {
		err = errors.New ("get value is nil")
		return
	}

	switch resTmp := value.DataInterface.(type) {
	case {{$type_one}}:
		res = resTmp
		return
	case json.Number:
		if err = json.Unmarshal ([]byte(resTmp.String ()), &res); err != nil {
			return
		}
	case map[string]*Dl,[]*Dl :
		resI := value.Call ()
		switch resITmp := resI.(type) {
		case {{$type_one}}:
			res = resITmp
			return
		default:
			err = errors.New ("error type:" + value.String())
			return
		}
	default:
		err = errors.New ("error type:" + value.String())
		return
	}

	return
}

func (self *Dl) SubNodeListGet{{CoverSnakeCaseToPascalCase $type_one}} (index int) (res {{$type_one}}, err error) {
	var value *Dl
	if value, err = self.SubNodeCheckIndex (index); err != nil {
		return
	}

	if value == nil {
		err = errors.New ("get value is nil")
		return
	}

	switch resTmp := value.DataInterface.(type) {
	case {{$type_one}}:
		res = resTmp
		return
	case json.Number:
		if err = json.Unmarshal ([]byte(resTmp.String ()), &res); err != nil {
			return
		}
	case map[string]*Dl,[]*Dl :
		resI := value.Call ()
		switch resITmp := resI.(type) {
		case {{$type_one}}:
			res = resITmp
			return
		default:
			err = errors.New ("error type" + value.String())
			return
		}
	default:
		err = errors.New ("error type" + value.String())
		return
	}

	return
}
{{end}}

// get node
func (self *Dl) SubNodeGet (key string) (res *Dl, err error) {
	if res, err = self.SubNodeCheckKey (key); err != nil {
		return
	}
	return
}

func (self *Dl) SubNodeListGet (index int) (res *Dl, err error) {
	if res, err = self.SubNodeCheckIndex (index); err != nil {
		return
	}
	return
}

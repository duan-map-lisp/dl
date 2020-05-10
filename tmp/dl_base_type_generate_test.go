package dl

func (self *Dl) CheckType() (res string) {
	tmpData := self.Call()
	if tmpData == nil {
		res = "null"
		return
	}
	switch tmpRes := tmpData.(type) {
{{$all_types := MkSlice "int8" "int16" "int32" "int64" "uint8" "uint16" "uint32" "uint64" "float32" "float64" "int" "uint" "uintptr" "string" "bool"}}
{{range $_, $type_one:= $all_types}}
	case {{$type_one}}:
		res = "{{$type_one}}"
		return
{{end}}
	case *Dl:
		if len(tmpRes.SubNodeTree) != 0 {
			res = "object"
			return
		}
		if len(tmpRes.SubNodeList) != 0 {
			res = "list"
			return
		}
	default:
		res = "null"
		return
	}

	return
}

func (self *Dl) setType() {
	Lambdas["type"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasNameForce("type")
		var err error
		var data *Dl

		if len(self.SubNodeTree) >= 2 {
			// object模型的type
			if data, err = self.SubNodeGet("data"); err != nil {
				panic(err)
			}
			resI = data.CheckType()
		} else if len(self.SubNodeList) == 2 {
			// list模型的type
			if data, err = self.SubNodeListGet(1); err != nil {
				panic(err)
			}
			resI = data.CheckType()
		} else {
			panic("atom format error")
		}

		return
	}
}

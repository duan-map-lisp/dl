package dl

import (
	"encoding/json"
)

func getSliceByte(dataNode *Dl) (data []byte) {
	var err error
	resNodeI := dataNode.Call()
	if resNodeI == nil {
		panic("get resNodeI is nil")
	}
	switch resTmp := resNodeI.(type) {
	case []byte:
		data = resTmp
		return
	}

	var tmpRes []byte
	if err = json.Unmarshal(dataNode.AllStr, &tmpRes); err != nil {
		panic("data type not []byte: " + err.Error())
	}
	data = tmpRes

	return
}

func (self *Dl) setEval() {
	Lambdas["eval"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasName("eval")
		var err error
		var dataNode *Dl
		if dataNode, err = self.SubNodeGet("data"); err != nil {
			panic("'data' not found")
		}
		data := getSliceByte(dataNode)

		evalSubNode := &Dl{
			FatherNode: self,
			AllStr:     data,
		}
		evalSubNode.Init()
		evalSubNode.Precompiling()
		resI = evalSubNode.Call()

		return
	}
}

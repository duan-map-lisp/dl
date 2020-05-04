package dl

import (
	"encoding/json"
	// log "github.com/sirupsen/logrus"
)

func GetSliceByte(dataNode *Dl) (data []byte) {
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
	// eval一定执行期，一定是顶层，宏不能继承给下一级eval
	Lambdas["eval"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasNameForce("eval")
		var err error
		var dataNode *Dl
		if len(self.SubNodeTree) >= 2 {
			if dataNode, err = self.SubNodeGet("data"); err != nil {
				panic("'data' not found")
			}
		} else if len(self.SubNodeList) == 2 {
			if dataNode, err = self.SubNodeListGet(1); err != nil {
				panic("'data' not found")
			}
		} else {
			panic("'eval' format error")
		}
		data := GetSliceByte(dataNode)

		evalSubNode := &Dl{
			FatherNode: self,
			AllStr:     data,
		}
		evalSubNode.Init()

		// 读取解析所有字符串解析成算法树
		evalSubNode.Precompiling()
		GenerateFlag = true
		for {
			// 如果宏被处理过，再处理一次
			if !GenerateFlag {
				// 如果宏已经不再被处理，预处理期结束
				break
			}
			GenerateFlag = false
			// 处理宏
			evalSubNode.Generate()
			// 处理正则展开宏
			evalSubNode.GenerateRegexp()
		}
		// 进入执行其前把预处理期垃圾回收一下
		evalSubNode.CleanGenerate()

		resI = evalSubNode.Call()

		return
	}
}

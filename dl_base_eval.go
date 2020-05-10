package dl

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

func (self *Dl) setEval() {
	// eval一定执行期，一定是顶层，宏不能继承给下一级eval
	Lambdas["eval"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasName("eval")
		var err error
		var data string
		switch tmpType := self.DataInterface.(type) {
		case map[string]*Dl:
			if data, err = self.SubNodeGetString("data"); err != nil {
				panic("'data' not found")
			}
		case []*Dl:
			if len(tmpType) != 2 {
				panic("'eval' format error")
			}
			if data, err = self.SubNodeListGetString(1); err != nil {
				panic("'data' not found")
			}
		default:
			panic("'eval' format error")
		}

		test := Dl{
			FatherNode: self,
		}
		evalSubNode := &test
		evalSubNode.Init()
		if err = json.Unmarshal([]byte(data), &test); err != nil {
			panic(err)
		}

		// 读取解析所有字符串解析成算法树
		log.Info("eval precompiling root：", evalSubNode.String())
		GenerateFlag = true
		for {
			// 如果宏被处理过，再处理一次
			if !GenerateFlag {
				// 如果宏已经不再被处理，预处理期结束
				break
			}
			GenerateFlag = false
			// 处理宏
			// evalSubNode.Generate()
			// 处理正则展开宏
			// evalSubNode.GenerateRegexp()
		}
		log.Info("eval precompiling end")

		// 进入执行期前把预处理期垃圾回收一下
		log.Info("begin to clean generate...")
		// CleanGenerate(evalSubNode)
		log.Info("end to clean generate...")

		// 计入执行期
		log.Info("begin to running...", evalSubNode.String())
		resI = evalSubNode.Call()
		log.Info("end running...")

		return
	}
	self.Symbols["eval"] = "eval"
}

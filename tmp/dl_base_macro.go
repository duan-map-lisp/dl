package dl

import (
	log "github.com/sirupsen/logrus"
)

func (self *Dl) setMacro() {
	Lambdas["macro"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasNameForce("macro")
		// macro实体是lambda
		// lambda操作比较复杂，只允许object模型
		if len(self.SubNodeList) != 0 {
			panic("'lambda' must be object type")
		} else if len(self.SubNodeTree) <= 1 {
			panic("'lambda' format error")
		}
		self.SubNodeTree["name"].TmpInterface = "lambda"

		var err error
		var symbol string
		if symbol, err = self.SubNodeGetSingleString("symbol"); err != nil {
			err = nil
			symbol = ""
		}

		lambdaFunc, ok := Lambdas["lambda"]
		if !ok {
			panic("lambda not found")
		}
		lambdaName := lambdaFunc(self)

		if symbol != "" {
			self.FatherNode.Symbols[symbol] = lambdaName
		}

		// macro定义后需要自我销毁，否则下一轮macro又会被执行
		if len(self.FatherNode.SubNodeTree) != 0 {
			// 如果父节点是object类型，delete销毁当前节点。
			delete(self.FatherNode.SubNodeTree, self.NodeName)
		} else if len(self.FatherNode.SubNodeList) != 0 {
			// 如果父节点是list类型的，将当前位置置为null，不能挪切片位置。
			// 挪切片位置会打乱list的顺序位置，这里请注意！！！
			// 所以定义宏在object中无影响，在list中会留下坑。可能导致执行期代码与预计不符。
			log.Info("end macro:", self.NodeIndex, self.String())
			self.FatherNode.SubNodeList[self.NodeIndex] = nil
		}

		resI = lambdaName
		return
	}
	return
}

package dl

import (
	log "github.com/sirupsen/logrus"
)

func (self *Dl) setRemacro() {
	Lambdas["remacro"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasNameForce("remacro")
		if self.FatherNode == nil {
			panic("remacro不可能是root节点")
		}

		var err error
		var regexp string
		var macro string
		if len(self.SubNodeTree) >= 3 {
			if regexp, err = self.SubNodeGetSingleString("regexp"); err != nil {
				panic("'regexp' not found")
			}
			if macro, err = self.SubNodeGetSingleString("macro"); err != nil {
				panic("'macro' not found")
			}
		} else if len(self.SubNodeList) == 3 {
			if regexp, err = self.SubNodeListGetSingleString(1); err != nil {
				panic("'regexp' not found")
			}
			if macro, err = self.SubNodeListGetSingleString(2); err != nil {
				panic("'macro' not found")
			}
		} else {
			panic("'remacro' format error")
		}

		if _, ok := self.FatherNode.RegexpMacros[regexp]; ok {
			panic("redefine remacro " + regexp)
		}
		self.FatherNode.RegexpMacros[regexp] = macro

		// remacro定义后需要自我销毁，否则下一轮macro又会被执行
		if len(self.FatherNode.SubNodeTree) != 0 {
			// 如果父节点是object类型，delete销毁当前节点。
			delete(self.FatherNode.SubNodeTree, self.NodeName)
		} else if len(self.FatherNode.SubNodeList) != 0 {
			// 如果父节点是list类型的，将当前位置置为null，不能挪切片位置。
			// 挪切片位置会打乱list的顺序位置，这里请注意！！！
			// 所以定义宏在object中无影响，在list中会留下坑。可能导致执行期代码与预计不符。
			log.Info("end macro:", self.NodeIndex, self.String())
			self.FatherNode.SubNodeList[self.NodeIndex] = nil
		} else {
			panic("")
		}

		return
	}
}

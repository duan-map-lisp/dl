package dl

import (
// log "github.com/sirupsen/logrus"
)

func (self *Dl) setLoad() {
	// load一定是预处理期
	Lambdas["load"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasNameForce("load")
		var err error
		var dataNode *Dl
		if dataNode, err = self.SubNodeGet("data"); err != nil {
			panic("'data' not found")
		}
		data := GetSliceByte(dataNode)

		if len(self.FatherNode.SubNodeTree) != 0 {
			loadSubNode := &Dl{
				FatherNode: self,
				AllStr:     data,
			}
			loadSubNode.Init()

			// 读取解析所有字符串解析成算法树
			loadSubNode.Precompiling()
			resI = loadSubNode
		} else if len(self.FatherNode.SubNodeList) != 0 {
			loadSubNode := &Dl{
				FatherNode: self,
				AllStr:     data,
			}
			loadSubNode.Init()

			// 读取解析所有字符串解析成算法树
			loadSubNode.Precompiling()
			resI = loadSubNode
		}

		return
	}
}

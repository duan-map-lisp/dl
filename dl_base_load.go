package dl

import (
// log "github.com/sirupsen/logrus"
)

func (self *Dl) setLoad() {
	// load一定是预处理期，否则返回的就是加载文件的数据正文，object或者list了
	Lambdas["load"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasNameForce("load")
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
			panic("'load' format error")
		}
		data := GetSliceByte(dataNode)

		loadSubNode := &Dl{
			FatherNode: self,
			AllStr:     data,
		}
		loadSubNode.Init()

		// 读取解析所有字符串解析成算法树
		loadSubNode.Precompiling()
		resI = loadSubNode

		return
	}
}

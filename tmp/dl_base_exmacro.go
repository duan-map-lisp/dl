package dl

import (
	log "github.com/sirupsen/logrus"
)

func (self *Dl) setExmacro() {
	Lambdas["exmacro"] = func(self *Dl) (resI interface{}) {
		// log.Debug("in Exmacro", self)
		self.CheckLambdasNameForce("exmacro")
		// 宏其实就是lambda，宏展开其实是在调用call，改个名字直接执行
		if len(self.SubNodeTree) != 0 {
			self.SubNodeTree["name"].TmpInterface = "call"
		} else if len(self.SubNodeList) != 0 {
			self.SubNodeList[0].TmpInterface = "call"
		}

		callFunc, ok := Lambdas["call"]
		if !ok {
			panic("call not found")
		}
		log.Info("begin exmacro self", self.String())
		// 拿到宏展开的结果，exRes
		exRes := callFunc(self)
		log.Debug("end exmacro self", self.String())
		log.Debug("res exmacro self", exRes.(*Dl).String())

		// 如果父节点为object格式，按照NodeName，用结果将自己覆盖了
		if len(self.FatherNode.SubNodeTree) != 0 {
			log.Debug("self.NodeName ", self.NodeName)
			switch resTmp := exRes.(type) {
			case *Dl:
				log.Info("res exmacro self", resTmp.String())
				resTmp.FatherNode = self.FatherNode
				self.FatherNode.SubNodeTree[self.NodeName] = resTmp
			default:
				newNode := &Dl{
					NodeName:   self.NodeName,
					FatherNode: self.FatherNode,
				}
				newNode.Init()
				newNode.TmpInterface = resTmp
				self.FatherNode.SubNodeTree[self.NodeName] = newNode
			}
		} else if len(self.FatherNode.SubNodeList) != 0 {
			log.Debug("self.NodeIndex ", self.NodeIndex)
			switch resTmp := exRes.(type) {
			case *Dl:
				log.Info("res exmacro self", resTmp.String())
				resTmp.FatherNode = self.FatherNode
				self.FatherNode.SubNodeList[self.NodeIndex] = resTmp
			default:
				newNode := &Dl{
					NodeIndex:  self.NodeIndex,
					FatherNode: self.FatherNode,
				}
				newNode.Init()
				newNode.TmpInterface = resTmp
				self.FatherNode.SubNodeList[self.NodeIndex] = newNode
			}
		} else {
			panic("")
		}
		log.Debug("exmacro end fatherNode: ", self.FatherNode)

		return
	}
	return
}

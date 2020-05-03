package dl

import (
	log "github.com/sirupsen/logrus"
)

func (self *Dl) setExmacro() {
	Lambdas["exmacro"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasNameForce("exmacro")
		log.Debug("in Exmacro", self)
		if len(self.SubNodeTree) != 0 {
			self.SubNodeTree["name"].TmpInterface = "call"
		} else if len(self.SubNodeList) != 0 {
			self.SubNodeList[0].TmpInterface = "call"
		}

		log.Debug("in Exmacro Lambdas", Lambdas)
		log.Debug("in Exmacro Symbols", self.Symbols)
		log.Debug("in Exmacro Father Symbols", self.FatherNode.Symbols)

		callFunc, ok := Lambdas["call"]
		if !ok {
			panic("call not found")
		}
		exRes := callFunc(self)
		log.Debug("exmacro self", self)
		if len(self.FatherNode.SubNodeTree) != 0 {
			log.Debug("self.NodeName ", self.NodeName)
			switch resTmp := exRes.(type) {
			case *Dl:
				resTmp.FatherNode = self.FatherNode
				self.FatherNode.SubNodeTree[self.NodeName] = resTmp
			default:
				newNode := &Dl{
					NodeName:     self.NodeName,
					FatherNode:   self.FatherNode,
					TmpInterface: resTmp,
				}
				newNode.Init()
				self.FatherNode.SubNodeTree[self.NodeName] = newNode
			}
		}
		if len(self.FatherNode.SubNodeList) != 0 {
			log.Debug("self.NodeIndex ", self.NodeIndex)
			switch resTmp := exRes.(type) {
			case *Dl:
				resTmp.FatherNode = self.FatherNode
				self.FatherNode.SubNodeList[self.NodeIndex] = resTmp
				log.Debug("?????????????", resTmp)
				log.Debug("?????????????", self.FatherNode)
			default:
				newNode := &Dl{
					NodeName:     self.NodeName,
					FatherNode:   self.FatherNode,
					TmpInterface: resTmp,
				}
				newNode.Init()
				self.FatherNode.SubNodeList[self.NodeIndex] = newNode
			}
		}

		return
	}
	return
}

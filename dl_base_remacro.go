package dl

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

		if _, ok := RegexpMacros[regexp]; ok {
			panic("redefine remacro " + regexp)
		}
		RegexpMacros[regexp] = macro

		if _, ok := self.FatherNode.Symbols[regexp]; ok {
			panic("redefine remacro " + regexp)
		}
		self.FatherNode.Symbols[regexp] = macro

		if len(self.FatherNode.SubNodeTree) != 0 {
			newNode := &Dl{
				NodeName:   self.NodeName,
				FatherNode: self.FatherNode,
			}
			newNode.Init()
			self.FatherNode.SubNodeTree[self.NodeName] = newNode
		} else if len(self.FatherNode.SubNodeList) != 0 {
			newNode := &Dl{
				NodeName:   self.NodeName,
				FatherNode: self.FatherNode,
			}
			newNode.Init()
			self.FatherNode.SubNodeList[self.NodeIndex] = newNode
		} else {
			panic("")
		}

		return
	}
}

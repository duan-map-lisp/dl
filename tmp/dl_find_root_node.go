package dl

func (self *Dl) FindRootNode() (res *Dl) {
	if self.FatherNode != nil {
		res = self.FatherNode.FindRootNode()
	} else {
		res = self
	}

	return
}

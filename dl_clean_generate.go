package dl

func (self *Dl) CleanGenerate() {
	for _, value := range self.SubNodeTree {
		value.Symbols = map[string]interface{}{}
	}
	for _, value := range self.SubNodeList {
		value.Symbols = map[string]interface{}{}
	}
	Lambdas = map[string]func(*Dl) interface{}{}
	RegexpMacros = map[string]string{}
	self.SetBaseFunc()
}

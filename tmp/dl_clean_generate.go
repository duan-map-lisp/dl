package dl

func CleanGenerate(self *Dl) {
	for _, value := range self.SubNodeTree {
		if value == nil {
			continue
		}
		value.Symbols = map[string]interface{}{}
		value.RegexpMacros = map[string]string{}
		CleanGenerate(value)
	}
	for _, value := range self.SubNodeList {
		if value == nil {
			continue
		}
		value.Symbols = map[string]interface{}{}
		value.RegexpMacros = map[string]string{}
		CleanGenerate(value)
	}
	Lambdas = map[string]func(*Dl) interface{}{}
	self.SetBaseFunc()
}

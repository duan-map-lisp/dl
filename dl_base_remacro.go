package dl

func (self *Dl) setRemacro() {
	Lambdas["remacro"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasNameForce("remacro")

		var err error
		var regexp string
		if regexp, err = self.SubNodeGetSingleString("regexp"); err != nil {
			panic(err)
		}

		var macro string
		if macro, err = self.SubNodeGetSingleString("macro"); err != nil {
			panic(err)
		}

		if _, ok := RegexpMacros[regexp]; ok {
			panic("redefine remacro " + regexp)
		}
		RegexpMacros[regexp] = macro

		return
	}
}

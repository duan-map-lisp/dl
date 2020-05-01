package dl

import (
	"plugin"
)

func (self *Dl) setPlugin() {
	self.Lambdas["plugin"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasName("plugin")
		var err error
		var path string
		if path, err = self.SubNodeGetSingleString("path"); err != nil {
			panic("'path' not found")
		}

		var tmpSo *plugin.Plugin
		if tmpSo, err = plugin.Open(path); err != nil {
			panic(err)
		}

		funcName, nameErr := tmpSo.Lookup("Load")
		if nameErr != nil {
			panic(nameErr)
		}

		funcName.(func(*Dl))(self.FatherNode)
		resI = true

		return
	}
}

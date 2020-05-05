package dl

import (
	"plugin"
)

func (self *Dl) setPlugin() {
	Lambdas["plugin"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasNameForce("plugin")
		var err error
		var filepath string
		var pkg string
		if len(self.SubNodeTree) >= 2 {
			if filepath, err = self.SubNodeGetSingleString("filepath"); err != nil {
				panic("'filepath' not found")
			}
			if pkg, err = self.SubNodeGetSingleString("package"); err != nil {
				pkg = ""
			}
		} else if len(self.SubNodeList) == 2 {
			if filepath, err = self.SubNodeListGetSingleString(1); err != nil {
				panic("'filepath' not found")
			}
			pkg = ""
		} else if len(self.SubNodeList) == 3 {
			if filepath, err = self.SubNodeListGetSingleString(1); err != nil {
				panic("'filepath' not found")
			}
			if pkg, err = self.SubNodeListGetSingleString(2); err != nil {
				panic("'package' not found")
			}
		} else {
			panic("'plugin' format error")
		}

		var tmpSo *plugin.Plugin
		if tmpSo, err = plugin.Open(filepath); err != nil {
			panic(err)
		}

		funcName, nameErr := tmpSo.Lookup("Load")
		if nameErr != nil {
			panic(nameErr)
		}

		funcName.(func(*Dl, string))(self, pkg)

		return
	}
}

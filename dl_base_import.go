package dl

import (
	"io/ioutil"
	"os"
)

func (self *Dl) setImport() {
	Lambdas["import"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasNameForce("import")
		var err error
		var filepath string
		if len(self.SubNodeTree) >= 2 {
			if filepath, err = self.SubNodeGetSingleString("filepath"); err != nil {
				panic("'filepath' not found")
			}
		} else if len(self.SubNodeList) == 2 {
			if filepath, err = self.SubNodeListGetSingleString(1); err != nil {
				panic("'filepath' not found")
			}
		} else {
			panic("'import' format error")
		}

		var file *os.File
		var res []byte

		if file, err = os.Open(filepath); err != nil {
			panic(err)
		}

		if res, err = ioutil.ReadAll(file); err != nil {
			panic(err)
		}

		resI = res
		return
	}
}

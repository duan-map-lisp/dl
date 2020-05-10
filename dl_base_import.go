package dl

import (
	"io/ioutil"
	"os"
)

func (self *Dl) setImport() {
	Lambdas["import"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasName("import")
		var err error
		var filepath string
		switch tmpType := self.DataInterface.(type) {
		case map[string]*Dl:
			if filepath, err = self.SubNodeGetString("filepath"); err != nil {
				panic("'filepath' not found")
			}
		case []*Dl:
			if len(tmpType) != 2 {
				panic("'import' format error")
			}
			if filepath, err = self.SubNodeListGetString(1); err != nil {
				panic("'filepath' not found")
			}
		default:
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

		resI = string(res)
		return
	}
	self.Symbols["import"] = "import"
}

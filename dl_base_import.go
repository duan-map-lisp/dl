package dl

import (
	"io/ioutil"
	"os"
)

func (self *Dl) setImport() {
	Lambdas["import"] = func(self *Dl) (resI interface{}) {
		self.CheckLambdasName("import")
		var err error
		var filename string
		if filename, err = self.SubNodeGetSingleString("filename"); err != nil {
			panic(err)
		}

		var file *os.File
		var res []byte

		if file, err = os.Open(filename); err != nil {
			panic(err)
		}

		if res, err = ioutil.ReadAll(file); err != nil {
			panic(err)
		}

		resI = res
		return
	}
}

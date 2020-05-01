package dl

import (
	"os"
	"io/ioutil"
	"encoding/json"
	"plugin"
)

type Dl struct {
	AllStr []byte

	NodeName string
	NodeIndex int
	FatherNode *Dl
	SubNodeTree map[string]*Dl
	SubNodeList []*Dl

	TmpList []json.RawMessage
	TmpMap map[string]json.RawMessage
	TmpInterface interface{}

	Lambdas map[string]func (*Dl) (interface{})
	Symbols map[string]interface{}
}

func (self *Dl) Init () {
	// 加载包含文件字符串
	self.Lambdas = map[string]func (*Dl) (interface{}) {}
	self.Lambdas["import"] = func (self *Dl) (resI interface{}) {
		var err error
		var name string
		if name, err = self.SubNodeGetString ("name"); err != nil {
			panic (err)
		}
		if name != "import" {
			panic ("function name not import")
		}

		var filename string
		if filename, err = self.SubNodeGetString ("filename"); err != nil {
			panic (err)
		}

		var file *os.File
		var res []byte

		if file, err = os.Open (filename); err != nil {
			panic (err)
		}

		if res, err = ioutil.ReadAll (file); err != nil {
			panic (err)
		}

		resI = res
		return
	}
	self.Lambdas["eval"] = func (self *Dl) (resI interface{}) {
		var err error
		var name string
		if name, err = self.SubNodeGetString ("name"); err != nil {
			panic (err)
		}
		if name != "eval" {
			panic ("function name not eval")
		}
		var data []byte
		if data, err = self.SubNodeGetBytes ("data"); err != nil {
			panic ("'data' not found")
		}

		resI = (&Dl {
			FatherNode: self,
			AllStr: data,
			SubNodeTree: map[string]*Dl {},
			Lambdas: map[string]func (*Dl) (interface{}) {},
		}).Eval ()
		return
	}
	self.Lambdas["plugin"] = func (self *Dl) (resI interface{}) {
		var err error
		var name string
		if name, err = self.SubNodeGetString ("name"); err != nil {
			panic (err)
		}
		if name != "plugin" {
			panic ("function name not plugin")
		}
		var path string
		if path, err = self.SubNodeGetString ("path"); err != nil {
			panic ("'path' not found")
		}

		var tmpSo *plugin.Plugin
		if tmpSo, err = plugin.Open (path); err != nil {
			panic (err)
		}

		funcName, nameErr := tmpSo.Lookup ("Load")
		if nameErr != nil {
			panic (nameErr)
		}

		funcName.(func (*Dl)) (self.FatherNode)
		resI = true

		return
	}
	self.Lambdas["block"] = func (self *Dl) (resI interface{}) {
		var err error
		var name string
		if name, err = self.SubNodeListGetString(0); err != nil {
			panic (err)
		}
		if name != "block" {
			panic ("function name not block")
		}

		for _, resOne := range self.SubNodeList {
			resI = resOne.Run ()
		}
		return
	}

	return
}

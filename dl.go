package dl

import (
	"encoding/json"
)

type Dl struct {
	AllStr []byte

	NodeName    string
	NodeIndex   int
	FatherNode  *Dl
	SubNodeTree map[string]*Dl
	SubNodeList []*Dl

	TmpList      []json.RawMessage
	TmpMap       map[string]json.RawMessage
	TmpInterface interface{}

	Lambdas map[string]func(*Dl) interface{}
	Symbols map[string]interface{}

	BlockBreakFlag bool
}

func (self *Dl) Init() {
	// 加载包含文件字符串
	self.Lambdas = map[string]func(*Dl) interface{}{}
	self.setImport()
	self.setEval()
	self.setPlugin()
	self.setBlock()
	self.setDefvar()

	return
}

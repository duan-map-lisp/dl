package dl

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

var Lambdas map[string]func(*Dl) interface{}
var RegexpMacros map[string]string
var GenerateFlag bool

func init() {
	Lambdas = map[string]func(*Dl) interface{}{}
	RegexpMacros = map[string]string{}
	// log.SetLevel(log.DebugLevel)
	log.SetLevel(log.InfoLevel)
}

type Dl struct {
	AllStr []byte

	NodeName    string
	NodeIndex   int
	FatherNode  *Dl
	SubNodeTree map[string]*Dl
	SubNodeList []*Dl
	SubNodeTmp  *Dl

	TmpList      []json.RawMessage
	TmpMap       map[string]json.RawMessage
	TmpInterface interface{}

	Symbols map[string]interface{}
}

func (self *Dl) Init() {
	// 加载包含文件字符串
	self.SubNodeTree = map[string]*Dl{}
	self.SubNodeList = []*Dl{}
	self.SubNodeTmp = nil
	self.TmpInterface = nil
	self.Symbols = map[string]interface{}{}

	return
}

func (self *Dl) SetBaseFunc() {
	// 打开一个文件按照[]byte格式读取内容
	self.setImport()
	// eval解释执行[]byte格式的代码
	self.setEval()
	// load加载[]byte格式的代码按照object返回
	self.setLoad()
	// 加载一个第三方中间件插件
	self.setPlugin()
	// 顺序执行代码，返回最后一行的执行结果，如果遇到break为true时中断并返回执行结果
	self.setBlock()
	// 定义一个变量
	self.setLet()
	// 获取一个变量的值
	self.setGet()
	// 执行一个定义过的lambda
	self.setCall()
	// 注册一个lambda函数
	self.setLambda()

	// 注册一个macro函数
	self.setMacro()
	// 注册一个正则macro函数
	self.setRemacro()
	// 展开一个定义过的macro
	self.setExmacro()
	// 正则宏保护函数
	self.setSafe()

	// 标准lisp操作函数
	self.setQuote()
	self.setType()
	self.setEq()
	self.setCar()
	self.setCdr()
	self.setCons()
}

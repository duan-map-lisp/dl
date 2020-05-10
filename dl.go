package dl

import (
	"bytes"
	"encoding/json"
	"os"

	log "github.com/sirupsen/logrus"
)

var Lambdas map[string]func(*Dl) interface{}
var GenerateFlag bool

func init() {
	Lambdas = map[string]func(*Dl) interface{}{}
	// log.SetLevel(log.DebugLevel)
	log.SetLevel(log.InfoLevel)
}

type Dl struct {
	FatherNode *Dl
	SubNodeTmp *Dl

	DataInterface interface{}

	Symbols      map[string]interface{}
	RegexpMacros map[string]string
}

func (self *Dl) MarshalJSON() (res []byte, err error) {
	res, err = json.Marshal(self.DataInterface)
	return
}

func (self *Dl) UnmarshalJSON(data []byte) (err error) {
	// 如果是list
	dataReader := bytes.NewReader(data)
	jsonDecode := json.NewDecoder(dataReader)
	jsonDecode.UseNumber()

	var tmpList []*Dl
	if err = jsonDecode.Decode(&tmpList); err == nil {
		for index, _ := range tmpList {
			tmpList[index].FatherNode = self
		}
		self.DataInterface = tmpList
		log.Debug("decode list json ok:", self.DataInterface)
		return
	}

	dataReader.Seek(0, os.SEEK_SET)
	var tmpObject map[string]*Dl
	if err = jsonDecode.Decode(&tmpObject); err == nil {
		for key, _ := range tmpObject {
			tmpObject[key].FatherNode = self
		}
		self.DataInterface = tmpObject
		log.Debug("decode object json ok:", self.DataInterface)
		return
	}

	dataReader.Seek(0, os.SEEK_SET)
	if err = jsonDecode.Decode(&self.DataInterface); err == nil {
		log.Debug("decode json ok:", self.DataInterface)
		switch tmpType := self.DataInterface.(type) {
		case json.Number:
			log.Debug("json.Number", tmpType)
		}
		return
	}

	panic(err)
	return
}

func (self *Dl) Init() {
	// 加载包含文件字符串
	self.Symbols = map[string]interface{}{}
	self.RegexpMacros = map[string]string{}

	return
}

func (self *Dl) SetBaseFunc() {
	// 打开一个文件按照[]byte格式读取内容
	self.setImport()
	// eval解释执行[]byte格式的代码
	self.setEval()
	// 顺序执行到最后，返回最后一条的值
	self.setBlock()
	/*
		// load加载[]byte格式的代码按照object返回
		self.setLoad()
		// 加载一个第三方中间件插件
		self.setPlugin()
		// 定义一个变量
		self.setLet()
		// 获取一个变量的值
		self.setGet()
		// 设置一个变量的值，类型自动推导，必须声明过，不会自动声明变量
		self.setSet()
		// 删除一个变量，会一直向上找到最近的一个symbol删除
		self.setDel()
		// 注册一个lambda函数
		self.setLambda()
		// 执行一个定义过的lambda
		self.setCall()
		// 将一个结点转换为string类型的json字符串返回。
		self.setString()

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
		self.setAppend()
		self.setCond()
	*/
}

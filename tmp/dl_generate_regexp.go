package dl

import (
	"encoding/json"
	"regexp"

	log "github.com/sirupsen/logrus"
)

func (self *Dl) GenerateRegexp() {
	if self == nil {
		return
	}
	if self.TmpInterface != nil {
		return
	}

	var ok bool
	var lambdaNameNode *Dl
	var lambdaName string
	if lambdaNameNode, ok = self.SubNodeTree["name"]; ok {
		switch tmpRes := lambdaNameNode.TmpInterface.(type) {
		case string:
			lambdaName = tmpRes
		}
	}
	if len(self.SubNodeList) > 0 {
		switch tmpRes := self.SubNodeList[0].TmpInterface.(type) {
		case string:
			lambdaName = tmpRes
		}
	}

	if lambdaName == "safe" {
		return
	}

	self.GenerateRegexpSubNode()
}

func (self *Dl) GenerateRegexpOne(subStrsArray [][]string, macroName string) {
	// 正则宏展开成exmacro形态。构造预计的json代码。
	tmpNode := struct {
		Name   string                 `json:"name"`
		Args   map[string]interface{} `json:"args"`
		Lambda map[string]string      `json:"lambda"`
	}{
		Name: "call",
		Args: map[string]interface{}{
			"reargs": []interface{}{
				"quote",
				subStrsArray,
			},
		},
		Lambda: map[string]string{
			"name":   "get",
			"symbol": macroName,
		},
	}

	// 生成json字符串，传入临时节点
	tmpAllStr, err := json.Marshal(tmpNode)
	if err != nil {
		panic(err)
	}
	log.Info("remacro json str is:", string(tmpAllStr))
	newNode := &Dl{
		NodeName:   self.NodeName,
		FatherNode: self,
		AllStr:     tmpAllStr,
	}

	newNode.Init()
	newNode.Precompiling()

	// 得到展开后的节点，覆盖到自己所在的位置
	if len(self.FatherNode.SubNodeTree) != 0 {
		self.FatherNode.SubNodeTree[self.NodeName] = newNode
	} else if len(self.FatherNode.SubNodeList) != 0 {
		self.FatherNode.SubNodeList[self.NodeIndex] = newNode
	} else {
		panic("")
	}

	log.Info("remacro begin call fatherNode is:", self.FatherNode.String())
	log.Info("remacro begin call self is:", self.String())
	log.Info("remacro begin call newNode is:", newNode.String())
	remacroEndRes := newNode.Call()
	log.Info("remacro end call self is:", self.String())

	// 得到的展开后的宏，展开的结果，覆盖到自己所在的位置
	if len(self.FatherNode.SubNodeTree) != 0 {
		log.Debug("self.NodeName ", self.NodeName)
		switch resTmp := remacroEndRes.(type) {
		case *Dl:
			log.Info("object res exmacro self", resTmp.String())
			resTmp.FatherNode = self.FatherNode
			self.FatherNode.SubNodeTree[self.NodeName] = resTmp
		default:
			newNode := &Dl{
				NodeName:   self.NodeName,
				FatherNode: self.FatherNode,
			}
			newNode.Init()
			newNode.TmpInterface = resTmp
			self.FatherNode.SubNodeTree[self.NodeName] = newNode
		}
	} else if len(self.FatherNode.SubNodeList) != 0 {
		log.Debug("self.NodeIndex ", self.NodeIndex)
		switch resTmp := remacroEndRes.(type) {
		case *Dl:
			log.Info("list res exmacro self", resTmp.String())
			resTmp.FatherNode = self.FatherNode
			self.FatherNode.SubNodeList[self.NodeIndex] = resTmp
		default:
			newNode := &Dl{
				NodeIndex:  self.NodeIndex,
				FatherNode: self.FatherNode,
			}
			newNode.Init()
			newNode.TmpInterface = resTmp
			self.FatherNode.SubNodeList[self.NodeIndex] = newNode
		}
	} else {
		panic("")
	}

	return
}

func (self *Dl) GenerateRegexpSubNode() {
	for _, value := range self.SubNodeTree {
		if value == nil {
			continue
		}
		var tmpString string
		if err := json.Unmarshal(value.AllStr, &tmpString); err != nil {
			err = nil
			continue
		}

		var regexpStr string
		var macroName string
		var nextNode *Dl
		var nextIndex int
		nextNode = self
		nextIndex = -1

		for {
			// 树结构从当前节点向上遍历，拿一个作用域内的正则匹配出来。
			regexpStr, macroName, nextNode, nextIndex = GetRegexpOne(nextNode, nextIndex)
			if nextNode == nil {
				break
			}

			re := regexp.MustCompile(regexpStr)
			subStrsArray := re.FindAllStringSubmatch(tmpString, -1)
			log.Debug("正则获取 ", subStrsArray)
			if len(subStrsArray) == 0 {
				continue
			}

			// 正则匹配成功，触发正则宏展开，标记还不能离开预处理期
			GenerateFlag = true

			// 进入正则宏展开
			value.GenerateRegexpOne(subStrsArray, macroName)
		}
	}

	for _, value := range self.SubNodeList {
		if value == nil {
			continue
		}
		var tmpString string
		if err := json.Unmarshal(value.AllStr, &tmpString); err != nil {
			err = nil
			continue
		}

		var regexpStr string
		var macroName string
		var nextNode *Dl
		var nextIndex int
		nextNode = self
		nextIndex = -1

		for {
			// 树结构从当前节点向上遍历，拿一个作用域内的正则匹配出来。
			regexpStr, macroName, nextNode, nextIndex = GetRegexpOne(nextNode, nextIndex)
			if nextNode == nil {
				break
			}

			re := regexp.MustCompile(regexpStr)
			subStrsArray := re.FindAllStringSubmatch(tmpString, -1)
			log.Debug("正则获取 ", subStrsArray)
			if len(subStrsArray) == 0 {
				continue
			}

			// 正则匹配成功，触发正则宏展开，标记还不能离开预处理期
			GenerateFlag = true

			// 进入正则宏展开
			value.GenerateRegexpOne(subStrsArray, macroName)
		}
	}
}

func GetRegexpOne(node *Dl, index int) (regexpStr string, macroName string, res *Dl, resIndex int) {
	if node == nil {
		regexpStr = ""
		macroName = ""
		res = nil
		resIndex = index
		return
	}
	tmpIndex := 0
	for tmpRegexpStr, tmpMacroName := range node.RegexpMacros {
		if tmpIndex <= index {
			tmpIndex++
			continue
		} else {
			regexpStr = tmpRegexpStr
			macroName = tmpMacroName
			res = node
			resIndex = tmpIndex
		}
		return
	}

	if node.FatherNode == nil {
		regexpStr = ""
		macroName = ""
		res = nil
	} else {
		regexpStr, macroName, res, resIndex = GetRegexpOne(node.FatherNode, -1)
	}

	return
}

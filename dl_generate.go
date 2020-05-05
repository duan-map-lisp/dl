package dl

import (
	log "github.com/sirupsen/logrus"
)

func (self *Dl) Generate() {
	// log.Debug("in Generate ", self.TmpInterface)
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

	if (lambdaName == "macro") || (lambdaName == "exmacro" || (lambdaName == "remacro")) {
		log.Info("get macro ", lambdaName)
		GenerateFlag = true
		_ = self.Call()
	} else {
		self.GenerateSubNode()
	}

	return
}

func (self *Dl) GenerateSubNode() {
	for _, tmpSubNode := range self.SubNodeTree {
		tmpSubNode.Generate()
	}

	for _, tmpSubNode := range self.SubNodeList {
		tmpSubNode.Generate()
	}
}

package dl

import (
	log "github.com/sirupsen/logrus"
)

func (self *Dl) Generate() {
	log.Debug("in Generate ", self.TmpInterface)
	if self.TmpInterface != nil {
		return
	}
	var err error
	var lambdaName string
	if lambdaName, err = self.SubNodeGetSingleString("name"); err != nil {
		if lambdaName, err = self.SubNodeListGetSingleString(0); err != nil {
			self.GenerateSubNode()
			return
		}
	}

	if (lambdaName == "macro") || (lambdaName == "exmacro" || (lambdaName == "remacro")) {
		GenerateFlag = true
		log.Debug("macro nodeName ", self.NodeName)
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

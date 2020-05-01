package dl

import (
	"fmt"
	// "encoding/json"
)

func (self *Dl) LambdaMap () (res interface{}) {
	var err error
	var lambdaName string
	lambdaName, err = self.SubNodeGetString ("name")
	if err != nil {
		panic (err)
	}

	lambdaOne := self.GetLambda (lambdaName)
	res = lambdaOne (self)
	return
}

func (self *Dl) GetLambda (lambdaName string) (res func(*Dl) (interface{})) {
	fmt.Println ("in get lambda", self.NodeName, self.Lambdas)
	res, ok := self.Lambdas[lambdaName]
	if !ok {
		if self.FatherNode == nil {
			panic ("未定义的函数：" + lambdaName)
		} else {
			res = self.FatherNode.GetLambda (lambdaName)
		}
	}

	return
}

func (self *Dl) LambdaList () (res interface{}) {
	var err error
	var lambdaName string
	lambdaName, err = self.SubNodeListGetString (0)
	if err != nil {
		panic (err)
	}
	lambdaOne := self.GetLambda (lambdaName)
	res = lambdaOne (self)
	return
}

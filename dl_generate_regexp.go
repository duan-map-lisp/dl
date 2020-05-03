package dl

import (
	"encoding/json"
	"regexp"
)

func (self *Dl) GenerateRegexp() {
	var err error
	var lambdaName string
	if lambdaName, err = self.SubNodeGetSingleString("name"); err != nil {
		if lambdaName, err = self.SubNodeListGetSingleString(0); err != nil {
			self.GenerateSubNode()
			return
		}
	}

	if lambdaName == "safe" {
		return
	}

	for key, value := range self.SubNodeTree {
		var tmpRes string
		if err = json.Unmarshal(value.AllStr, &tmpRes); err != nil {
			err = nil
			continue
		}

		for regexpStr, macroName := range RegexpMacros {
			re := regexp.MustCompile(regexpStr)
			subStrsArray := re.FindAllStringSubmatch(tmpRes, -1)
			if len(subStrsArray) == 0 {
				continue
			}

			GenerateFlag = true
			resI := self.GenerateRegexpOne(subStrsArray, macroName)
			switch resTmp := resI.(type) {
			case *Dl:
				self.SubNodeTree[key] = resTmp
			default:
				newNode := &Dl{
					NodeName:     key,
					FatherNode:   self,
					TmpInterface: resTmp,
				}
				newNode.Init()
				self.SubNodeTree[key] = newNode
			}
			break
		}
	}

	for index, value := range self.SubNodeList {
		var tmpRes string
		if err = json.Unmarshal(value.AllStr, &tmpRes); err != nil {
			err = nil
			continue
		}

		for regexpStr, macroName := range RegexpMacros {
			re := regexp.MustCompile(regexpStr)
			subStrsArray := re.FindAllStringSubmatch(tmpRes, -1)
			if len(subStrsArray) == 0 {
				continue
			}

			GenerateFlag = true
			resI := self.GenerateRegexpOne(subStrsArray, macroName)
			switch resTmp := resI.(type) {
			case *Dl:
				self.SubNodeList[index] = resTmp
			default:
				newNode := &Dl{
					NodeIndex:    index,
					FatherNode:   self,
					TmpInterface: resTmp,
				}
				newNode.Init()
				self.SubNodeList[index] = newNode
			}
			break
		}
	}
}

func (self *Dl) GenerateRegexpOne(subStrsArray [][]string, macroName string) (resI interface{}) {
	newNode := &Dl{
		FatherNode: self,
		AllStr:     []byte(`{"name":"exmacro", "args": {"reargs":[]}, "lambda": {"name": "get", "symbol":""}}`),
	}
	newNode.Init()
	newNode.Precompiling()

	var ok bool
	var tmpArgs *Dl
	tmpArgs, ok = newNode.SubNodeTree["args"]
	if !ok {
		panic("")
	}
	var tmpArgsReargs *Dl
	tmpArgsReargs, ok = tmpArgs.SubNodeTree["reargs"]
	if !ok {
		panic("")
	}
	for _, subStrs := range subStrsArray {
		subStrsOne := &Dl{
			FatherNode: tmpArgsReargs,
		}
		subStrsOne.Init()
		for _, subStr := range subStrs {
			subStrOne := &Dl{
				FatherNode:   subStrsOne,
				TmpInterface: subStr,
			}
			subStrOne.Init()
			subStrsOne.SubNodeList = append(subStrsOne.SubNodeList, subStrOne)
		}

		tmpArgsReargs.SubNodeList = append(tmpArgsReargs.SubNodeList, subStrsOne)
	}

	var tmpLambda *Dl
	tmpLambda, ok = newNode.SubNodeTree["lambda"]
	if !ok {
		panic("")
	}
	var tmpLambdaSymbol *Dl
	tmpLambdaSymbol, ok = tmpLambda.SubNodeTree["symbol"]
	if !ok {
		panic("")
	}
	tmpLambdaSymbol.TmpInterface = macroName

	resI = newNode.Call()

	return
}

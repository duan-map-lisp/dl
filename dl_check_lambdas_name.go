package dl

import (
	"errors"
)

func (self *Dl) CheckLambdasNameForce(checkName string) {
	var err error
	if err = self.CheckLambdasName(checkName); err != nil {
		panic(err)
	}
}

func (self *Dl) CheckLambdasName(checkName string) (err error) {
	var name string
	if len(self.SubNodeTree) >= 1 {
		if name, err = self.SubNodeGetSingleString("name"); err != nil {
			err = errors.New("not found lambda name")
			return
		}
	} else if len(self.SubNodeList) >= 1 {
		if name, err = self.SubNodeListGetSingleString(0); err != nil {
			err = errors.New("not found lambda name")
			return
		}
	} else {
		err = errors.New("not found lambda name")
		return
	}

	if name != checkName {
		err = errors.New("lambda name not " + checkName)
	}

	return
}

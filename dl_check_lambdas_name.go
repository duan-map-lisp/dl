package dl

import (
	"errors"
)

func (self *Dl) CheckLambdasName(checkName string) {
	var err error
	var name string
	if name, err = self.GetLambdasName(); err != nil {
		panic(err)
	}
	if name != checkName {
		panic("lambda name not " + checkName)
	}
}

func (self *Dl) GetLambdasName() (name string, err error) {
	switch self.DataInterface.(type) {
	case map[string]*Dl:
		if name, err = self.SubNodeGetString("name"); err != nil {
			err = errors.New("not found lambda name")
			return
		}
	case []*Dl:
		if name, err = self.SubNodeListGetString(0); err != nil {
			err = errors.New("not found lambda name")
			return
		}
	default:
		err = errors.New("not found lambda name")
	}
	return
}

package dl

import (
	"errors"
)

func (self *Dl) SubNodeCheckKey (key string) (value *Dl, err error) {
	var ok bool
	if value, ok = self.SubNodeTree[key]; !ok {
		err = errors.New ("sub node key " + key + " not found")
		return
	}
	return
}

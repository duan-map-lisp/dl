package dl

func (self *Dl) CheckLambdasName(checkName string) {
	var err error
	var name string
	if name, err = self.SubNodeGetSingleString("name"); err != nil {
		if name, err = self.SubNodeListGetSingleString(0); err != nil {
			panic("not found lambda name")
		}
	}
	if name != checkName {
		panic("lambda name not " + checkName)
	}

	return
}

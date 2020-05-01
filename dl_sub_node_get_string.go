package dl

func (self *Dl) SubNodeGetString (key string) (res string, err error) {
	var value *Dl
	if value, err = self.SubNodeCheckKey (key); err != nil {
		return
	}
	resI := value.Run ()
	switch resTmp := resI.(type) {
	case string:
		res = resTmp
	default:
		panic (key + " type not string")
	}
	return
}

func (self *Dl) SubNodeListGetString (index int) (res string, err error) {
	if len(self.SubNodeList) <= index {
		panic ("index out range")
	}

	value := self.SubNodeList[index]
	resI := value.Run ()
	switch resTmp := resI.(type) {
	case string:
		res = resTmp
	default:
		panic ("type not string")
	}
	return
}

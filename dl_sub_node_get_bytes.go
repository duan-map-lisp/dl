package dl

func (self *Dl) SubNodeGetBytes (key string) (res []byte, err error) {
	var value *Dl
	if value, err = self.SubNodeCheckKey (key); err != nil {
		return
	}
	resI := value.Run ()
	switch resTmp := resI.(type) {
	case []byte:
		res = resTmp
	default:
		panic (key + " type not bytes")
	}
	return
}

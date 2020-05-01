package dl

import (
	// "fmt"
	// "encoding/json"
)

func (self *Dl) Eval () (resI interface{}) {
	// self.PrecompilingRegexp ()
	self.Precompiling ()
	resI = self.Run ()

	return
}

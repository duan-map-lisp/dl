package main

import (
	"fmt"
	"os"
	"flag"

	"github.com/duan-map-lisp/dl"
)

func main () {
	var filename string
	flag.StringVar (&filename, "f", "", "文件路径")
	defer flag.PrintDefaults ()
	if len (os.Args) < 2 {
		return
	}
	flag.Parse ()

	if filename == "" {
		filename = os.Args[len (os.Args) - 1]
	}

	fmt.Println ("dl filename: ", filename)


	test := dl.Dl {
		NodeName: "root",
		AllStr: []byte (`{"name": "eval", "data": {"name": "import", "filename": "` + filename + `"}}`),
		SubNodeTree: map[string]*dl.Dl {},
		Lambdas: map[string]func (*dl.Dl) (interface{}) {},
		Symbols: map[string]interface{} {},
	}
	(&test).Init ()
	res := (&test).Eval ()
	fmt.Println ("结果：", res)
}

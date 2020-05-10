package main

import (
	"fmt"
	"os"
	"flag"
	"encoding/json"

	"github.com/duan-map-lisp/dl"
)

func main () {
	var filepath string
	flag.StringVar (&filepath, "f", "", "文件路径")
	defer flag.PrintDefaults ()
	if len (os.Args) < 2 {
		return
	}
	flag.Parse ()

	if filepath == "" {
		filepath = os.Args[len (os.Args) - 1]
	}

	fmt.Println ("dl filepath: ", filepath)

	AllStr , err := json.Marshal (struct {
		Name string `json:"name"`
		Data interface {} `json:"data"`
	} {
		Name: "eval",
		Data: struct {
			Name string `json:"name"`
			Filepath string `json:"filepath"`
		}{
			Name: "import",
			Filepath: filepath,
		},
	})
	if err != nil {
		panic ("cover eval file error")
	}

	var test dl.Dl
	if err = json.Unmarshal (AllStr, &test); err != nil {
		panic (err)
	}
	(&test).Init ()
	(&test).SetBaseFunc ()
	evalFuncSymbol := (&test).GetSymbol ("eval")
	if evalFuncSymbol == nil {
		panic ("eval not found")
	}
	evalFunc, ok := dl.Lambdas[evalFuncSymbol.(string)]
	if !ok {
		panic ("eval not found")
	}
	res := evalFunc (&test)
	switch resTmp := res.(type) {
	case *dl.Dl:
		fmt.Println ("结果：", resTmp.String ())
	default:
		fmt.Println ("结果：", res)
	}
	return
}

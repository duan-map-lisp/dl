package main

import (
	"os"
	"strings"
	"regexp"
	"text/template"
	"path"
)

func main () {
	if len (os.Args) != 2 {
		panic (os.Args)
	}

	t, err := template.New (path.Base (os.Args[1])).Funcs (template.FuncMap {
		"MkSlice": MkSlice,
		"MkMap": MkMap,
		"CoverSnakeCaseToPascalCase": CoverSnakeCaseToPascalCase,
		"GetFormat": GetFormat,
	}).ParseFiles (os.Args[1])
	if err != nil {
		panic (err)
	}

	if ok, err := regexp.MatchString ("^.*_generate_test.go$", os.Args[1]); err != nil || !ok {
		panic (err)
	}

	var outFile *os.File
	if outFile, err = os.OpenFile(os.Args[1][:len (os.Args[1]) - 17] + "_generate_drop.go", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644); err != nil {
		panic (err)
	}

	err = t.Execute (outFile, nil)
}

func MkSlice (args ...interface{}) []interface{} {
	return args
}

func MkMap (args ...interface{}) map[string]interface{} {
	if len(args) % 2 != 0 {
		panic("bad makemap")
	}

	m := make(map[string]interface{})
	for i := 0; i < len(args); i += 2 {
		switch key := args[i].(type) {
		case string:
			m[key] = args[i + 1]
		}
	}
	return m
}

func CoverSnakeCaseToPascalCase (inStr string) (outStr string) {
	tmpArray := strings.Split (inStr, "_")
	outStr = ""
	for _, tmpOne := range tmpArray {
		tmpOneByte := []byte(tmpOne)
		tmpOneByte[0] = strings.ToUpper (string (tmpOneByte[0]))[0]
		outStr += string (tmpOneByte)
	}

	return
}

func GetFormat (inStr string, format string) (outStr string) {
	if ("single" == format) {
		outStr = inStr
	} else if ("slice" == format) {
		outStr = "[]" + inStr
	} else if ("map" == format) {
		outStr = "map[string]" + inStr
	}

	return
}

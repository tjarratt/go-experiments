package main

import(
	"fmt"
	"os"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	if len(os.Args) != 2 {
		println("expected exactly one arg: path to a file to inspect")
		os.Exit(1)
	}

	fileSet := token.NewFileSet()
	rootNode, err := parser.ParseFile(fileSet, os.Args[1], nil, 0)
	if err != nil {
		panic(err.Error())
	}

	ast.Inspect(rootNode, func(node ast.Node) bool {
		if node == nil {
			return false
		}

		fmt.Printf("%#v\n", node)
		return true
	})

	return
}

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		println("usage: dump_ast  path/to/a/file.go")
		os.Exit(1)
	}

	fileSet := token.NewFileSet()
	rootNode, err := parser.ParseFile(fileSet, os.Args[1], nil, 0)
	if err != nil {
		panic(err.Error())
	}

	ast.Inspect(rootNode, func(node ast.Node) bool {
		if node == nil {
			return true
		}

		switch node := node.(type) {
		case *ast.Ident:
			if node.Obj != nil {
				fmt.Printf("%p %#v\n%p %#v\n", node, node, node.Obj, node.Obj)
			} else {
				fmt.Printf("%p %#v\n", node, node)
			}
		default:
			fmt.Printf("%p %#v\n", node, node)
		}

		return true
	})

	return
}

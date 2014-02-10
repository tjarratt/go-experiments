package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"io/ioutil"
	"os"
)

func main() {
	file := &ast.File{}
	packageName := &ast.Ident{Name: "foo"}
	file.Package = 1
	file.Name = packageName

	genDecl := &ast.GenDecl{
		TokPos: 14,
		Tok:    85,
		Lparen: 0,
		Rparen: 0,
	}
	genDecl.Specs = append(genDecl.Specs, createDescribe())
	file.Decls = append(file.Decls, genDecl)

	if len(os.Args) > 1 && os.Args[1] == "write" {
		var buffer bytes.Buffer
		fileSet := token.NewFileSet()
		err := format.Node(&buffer, fileSet, file)
		if err != nil {
			println("whoops!", err.Error())
			os.Exit(1)
		}

		ioutil.WriteFile("test_underscore_describe.go", buffer.Bytes(), 0600)
		println("done!")
	} else {
		ast.Inspect(file, func(node ast.Node) bool {
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
	}
}

func createDescribe() *ast.ValueSpec {
	fieldList := &ast.FieldList{Opening: 48, Closing: 49}
	funcType := &ast.FuncType{Params: fieldList, Func: 44}
	funcLit := &ast.FuncLit{
		Type: funcType,
		Body: &ast.BlockStmt{Lbrace: 51, Rbrace: 54},
	}
	description := &ast.BasicLit{Kind: 9, Value: `"something cool"`}
	describe := &ast.Ident{Name: "Describe"}
	callExpr := &ast.CallExpr{
		Fun:    describe,
		Args:   []ast.Expr{description, funcLit},
		Lparen: 30,
		Rparen: 55,
	}
	underscoreObj := &ast.Object{Kind: 4, Name: "_", Data: 0}
	underscore := &ast.Ident{NamePos: 18, Name: "_", Obj: underscoreObj}
	valueSpec := &ast.ValueSpec{
		Names:  []*ast.Ident{underscore},
		Values: []ast.Expr{callExpr},
	}

	underscoreObj.Decl = valueSpec
	return valueSpec
}

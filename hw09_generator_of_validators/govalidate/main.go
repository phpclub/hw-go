package main

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
)

var (
	ErrMissingParam = errors.New("missing param")
	// ErrFailedReadDir  = errors.New("failed reading directory")
	// ErrFailedReadFile = errors.New("failed read file")
)

const NeedArgs = 2

func parseAst(originalf string) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, originalf, nil, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	//ast.Print(fset, f)
	//return
	for _, decl := range f.Decls {
		if typeDecl, ok := decl.(*ast.GenDecl); ok && token.TYPE == typeDecl.Tok {
			printType(typeDecl, fset)
		}
	}
}

func printType(typeDecl *ast.GenDecl, fset *token.FileSet) {
	for _, spec := range typeDecl.Specs {
		//fmt.Printf("%+v ", typeDecl.Tok)
		fmt.Printf("type") //typeDecl.Tok)
		fmt.Printf(" %s ", spec.(*ast.TypeSpec).Name)
		if ident, ok := spec.(*ast.TypeSpec).Type.(*ast.Ident); ok {
			fmt.Printf("  %s \n", ident.Name)
		}
		if structDecl, ok := spec.(*ast.TypeSpec).Type.(*ast.StructType); ok {
			printStructure(structDecl, fset)
		}
		fmt.Printf("\n")
	}
}

func printStructure(structDecl *ast.StructType, fset *token.FileSet) {
	fmt.Printf(" struct {\n")
	for _, fld := range structDecl.Fields.List {
		fmt.Printf("\t %s", fld.Names[0])
		if ident, ok := fld.Type.(*ast.Ident); ok {
			fmt.Printf(" %s ", ident.Name)
		}
		if AT, ok := fld.Type.(*ast.ArrayType); ok {
			if AT.Len == nil {
				fmt.Printf(" []%s", AT.Elt)
			} else {
				fmt.Printf(" [%s]%s", AT.Len.(*ast.BasicLit).Value, AT.Elt)
			}
		}
		if fld.Tag != nil {
			var tagNameBuf bytes.Buffer
			err := printer.Fprint(&tagNameBuf, fset, fld.Tag)
			if err != nil {
				log.Fatalf("failed printing %s", err)
			}
			fmt.Printf(" %s", tagNameBuf.String())
		}
		print("\n")
	}
	fmt.Printf("}\n")
}

func main() {
	if len(os.Args) < NeedArgs {
		println(ErrMissingParam.Error())
		return
	}
	parseAst(os.Args[1])
}

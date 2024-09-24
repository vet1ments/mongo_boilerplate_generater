package structparser

import (
	"boilerplate/internal/utils"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

type StructInfo struct {
	Name   string
	Fields []*FieldInfo
}

type FieldInfo struct {
	Name string
	Type string
}

func ParseStructs(filename string) ([]*StructInfo, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	var structs []*StructInfo
	ast.Inspect(node, func(n ast.Node) bool {
		if genDecl, ok := n.(*ast.GenDecl); ok {
			//fmt.Println(genDecl.Doc.Text())
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					if structType, ok := typeSpec.Type.(*ast.StructType); ok {
						structName := typeSpec.Name.Name

						if genDecl.Doc != nil {
							fmt.Println(strings.HasPrefix(strings.ToLower(genDecl.Doc.Text()), "embed"))
						}

						struct_ := &StructInfo{
							Name: structName,
						}

						for _, field := range structType.Fields.List {
							field_ := &FieldInfo{
								Type: utils.TypeToString(field.Type),
							}
							for _, fieldName := range field.Names {
								field_.Name = fieldName.Name
							}
							struct_.Fields = append(struct_.Fields, field_)
						}
						structs = append(structs, struct_)
					}
				}
			}

		}

		return true
	})

	return structs, nil
}

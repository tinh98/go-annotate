package goan

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
)

func GenerateSwaggerAnnotations(file string) error {
	// Parse the file and generate an AST.
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	for _, decl := range node.Decls {

		funcDecl, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}

		// Check if the function already has a comment.
		if funcDecl.Doc != nil {
			comments := funcDecl.Doc.List
			if len(comments) > 0 && strings.HasPrefix(comments[0].Text, "// @") {
				//continue // The function already has a swagger annotation.
			}
		}

		var methodName, recvType string
		if funcDecl.Recv != nil {
			for _, field := range funcDecl.Recv.List {
				switch t := field.Type.(type) {
				case *ast.Ident:
					recvType = fmt.Sprintf("%v", t)
				case *ast.StarExpr:
					recvType = fmt.Sprintf("*%v", t.X)
				}
			}
		}

		methodName = funcDecl.Name.Name

		// Get the function parameters and return type.
		params := []string{}
		for _, param := range funcDecl.Type.Params.List {
			paramType := fmt.Sprintf("%v", param.Type)
			paramName := param.Names[0].Name
			params = append(params, fmt.Sprintf("%s %s", paramName, paramType))
		}

		// Generate the swagger annotation.
		var swaggerComments bytes.Buffer
		fmt.Fprintf(&swaggerComments, "// %s\n", methodName)
		fmt.Fprintf(&swaggerComments, "// @Summary TODO\n")
		fmt.Fprintf(&swaggerComments, "// @Description TODO\n")
		fmt.Fprintf(&swaggerComments, "// @Tags TODO\n")
		fmt.Fprintf(&swaggerComments, "// @Accept json\n")
		fmt.Fprintf(&swaggerComments, "// @Produce json\n")

		if recvType != "" {
			fmt.Fprintf(&swaggerComments, "// @Param %s path string true \"%s ID\"\n", strings.ToLower(recvType), recvType)
		}

		for _, param := range params {
			parts := strings.Split(param, " ")
			fmt.Fprintf(&swaggerComments, "// @Param %s body %s true \"%s\"\n", parts[0], parts[1], parts[0])
		}

		for _, field := range funcDecl.Type.Params.List {
			paramType := getType(field.Type)
			paramName := ""
			if len(field.Names) > 0 {
				paramName = field.Names[0].Name
			}

			if paramName != "" {
				fmt.Fprintf(&swaggerComments, "// @Param %s query %s true \"%s\"\n", paramName, paramType, paramName)
			}
		}

		//returnType := getType(funcDecl.Type.Results.List[0].Type)
		returnType := ""
		if funcDecl.Type.Results != nil {
			returnType = fmt.Sprintf("%v", funcDecl.Type.Results.List[0].Type)
		}

		if returnType != "" {
			fmt.Fprintf(&swaggerComments, "// @Success 200 {object} %s\n", returnType)
		}

		fmt.Fprintf(&swaggerComments, "// @Router /%s/%s [get]\n", strings.ToLower(recvType), getMethodName(methodName))
		commentGroup := &ast.CommentGroup{}
		commentGroup.List = append(commentGroup.List, &ast.Comment{Text: swaggerComments.String()})

		funcDecl.Doc = commentGroup
	}

	// Write the modified code back to the file.
	outFile, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	if err := format.Node(outFile, fset, node); err != nil {
		log.Fatal(err)
	}
	//fmt.Println(buf.String())
	fmt.Printf("%s has been updated with swagger annotations.\n", file)

	return nil
}

func getType(expr ast.Expr) string {
	switch expr.(type) {
	case *ast.Ident:
		return expr.(*ast.Ident).Name
	case *ast.ArrayType:
		return fmt.Sprintf("[]%s", getType(expr.(*ast.ArrayType).Elt))
	case *ast.StarExpr:
		return fmt.Sprintf("*%s", getType(expr.(*ast.StarExpr).X))
	default:
		return ""
	}
}

func getMethodName(funcName string) string {
	replacements := []string{"Get", "Post", "Push", "Delete"}
	for _, r := range replacements {
		funcName = strings.ReplaceAll(funcName, r, "")
	}
	return funcName
}

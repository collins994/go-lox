package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: generate_ast.exe <file base name>")
		return
	}

		var err error = defineAst(os.Args[1], []string{
			"Binary : left Expr , operator Token , right Expr",
			"Grouping : expression Expr",
			"Literal : value Object",
			"Unary : operator Token , right Expr",
		})

		if err != nil {
			log.Fatal(err)
		}
	}

	func defineAst(basename string, types []string) error {
		var fileName = fmt.Sprintf(".\\code\\%v.go", basename)
		var builder strings.Builder
		var imports = []string{"fmt"};

		// package name and imports
		builder.WriteString("package main \n\nimport (\n")
		for _, package_name := range imports {
			builder.WriteString(fmt.Sprintf("\t\"%v\"\n", package_name))
		}
		builder.WriteString(")\n\n")

		// Expr interface type
		builder.WriteString("type Expr interface{}\n");

		for _, typ := range types {
			var s = strings.Split(typ, ":")
			var name = s[0];
			var fields = s[1];
			builder.WriteString(fmt.Sprintf("type %v struct {\n", name));
			for _, field := range strings.Split(fields, ",") {
				builder.WriteString(fmt.Sprintf("\t%v\n", field))
			}
			builder.WriteString("}\n\n")
		}

	err := os.WriteFile(fileName, []byte(builder.String()), 0644) // 0644 - read/write permissions for owner
	if err != nil {
		return err
	}

	return nil
}

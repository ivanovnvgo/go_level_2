//Go. Уровень 2. Урок 7. Рефлексия и кодогенерация.
//Написать кодогенератор под какую-нибудь задачу.
//Задача следующая: есть список констант типа Color,
//необходимо получить массив, содержащий в себе все константы Color.
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
	"text/template"
)

var constListTmpl = `//an array containing all constants of the Color type 
//is formed from the list of constants of the Color type
package {{.Package}}

type {{.Name}}s []{{.Name}}
func (c {{.Name}}s)List() []{{.Name}} {
	return []{{.Name}}{{"{"}}{{.List}}{{"}"}}
}
`

func main() {
	var source = `
package main

type Color int

const (
	Green Color = iota
	Red
	Blue
	Black
)
	`
	//получаем ast (abstract syntax tree) файла с исходным кодом на языке go.AST
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", []byte(source), 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	typeName := "Color"         //тип констант для которых будет создан список
	typ := ""                   //для запоминания последнего определенного типа в ast
	consts := make([]string, 0) //массив для сохранения найденных констант
	for _, decl := range f.Decls {
		//массив с определениями типов, переменных, констант, функций и т.п.
		switch decl := decl.(type) {
		case *ast.GenDecl:
			switch decl.Tok {
			case token.CONST: //нам интересны только константы
				for _, spec := range decl.Specs {
					vspec := spec.(*ast.ValueSpec) //отсюда мы получим наименование константы
					if vspec.Type == nil && len(vspec.Values) > 0 {
						//случай определения константы как "X = 1"
						//такая константа не имеет типа и может быть пропущена
						//это может означать, что был начат новый блок определения const
						typ = ""
						continue
					}
					if vspec.Type != nil {
						//"const Green Color" - запоминаем тип константы
						if ident, ok := vspec.Type.(*ast.Ident); ok {
							typ = ident.Name
						} else {
							continue
						}

					}
					if typ == typeName {
						//тип константы совпадает с искомым, запоминаем имя константы в массив consts
						consts = append(consts, vspec.Names[0].Name)
					}
				}
			}
		}
	}

	templateData := struct {
		Package string
		Name    string
		List    string
	}{
		Package: "main",
		Name:    typeName,
		List:    strings.Join(consts, ", "),
	}
	t := template.Must(template.New("const-list").Parse(constListTmpl))

	if t.Execute(os.Stdout, templateData) != nil {
		fmt.Println(err)
	}
}

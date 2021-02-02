// text/template is a useful text generating tool.
// Related examples: http://golang.org/pkg/text/template/#pkg-examples
package main

import (
    "os"
    "text/template"
    "fmt"
)

var templateFuncs = template.FuncMap{
	"rpad": rpad,
}
// rpad adds padding to the right of a string.
func rpad(s string, padding int) string {
	template := fmt.Sprintf("%%-%ds", padding)
	return fmt.Sprintf(template, s)
}


func check(err error) {
    if err != nil {
        panic(err)
    }
}
func main() {
    t := template.New("help")
    t.Funcs(templateFuncs)
    t = template.Must(t.Parse("{{rpad .Name .NamePadding}}\n"))

    cmd := struct {
        Name string,
        NamePadding int,
    }{
        "abc",
        5
    }
    check(t.Execute(os.Stdout, cmd))
}


// https://golang.org/pkg/text/template/#hdr-Nested_template_definitions
package main

import (
    "os"
    "text/template"
)


func check(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    t := template.Must(template.New("t1").
        Parse(`{{define "T1"}}ONE{{end}}
{{define "T2"}}TWO{{end}}
{{define "T3"}}{{template "T1"}} {{template "T2"}}{{end}}
{{template "T3"}}`))
    check(t.Execute(os.Stdout, nil))
}


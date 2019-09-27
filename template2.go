// https://stackoverflow.com/questions/11123865/format-a-go-string-without-printing
package main

import(
	"text/template"
	"bytes"
)

func main(){

	const emailTmpl = `Hi {{.Name}}!

	Your account is ready, your user name is: {{.UserName}}

	You have the following roles assigned:
	{{range $i, $r := .Roles}}{{if ne $i 0}}, {{end}}{{.}}{{end}}`

	data := map[string]interface{}{
		"Name":     "Bob",
		"UserName": "bob92",
		"Roles":    []string{"dbteam", "uiteam", "tester"},
	}

	t := template.Must(template.New("email").Parse(emailTmpl))
	buf := &bytes.Buffer{}
	if err := t.Execute(buf, data); err != nil {
		panic(err)
	}
	s := buf.String()

	println(s)
}

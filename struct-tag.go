// https://play.golang.org/p/0SJT8I6gGR5
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name_field"`
	Age  int
}

func main() {
	user := &User{"John Doe The Fourth", 20}

	field, ok := reflect.TypeOf(user).Elem().FieldByName("Name")
	if !ok {
		panic("Field not found")
	}
	fmt.Println(getStructTag(field))
}

func getStructTag(f reflect.StructField) string {
	return string(f.Tag)
}

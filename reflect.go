// https://stackoverflow.com/questions/23350173/how-do-you-loop-through-the-fields-in-a-golang-struct-to-get-and-set-values-in-a?rq=1

package main

import (
	"fmt"
	"reflect"
)

func main() {
	type T struct {
		A int
		B string
	}
	t := &T{23, "skidoo"}
	t2 := &T{}
	Merge(t, t2)
	//s := reflect.ValueOf(&t).Elem()
	//s2 := reflect.ValueOf(&t2).Elem()
	//typeOfT := s.Type()
	fmt.Println("t=", t)
	fmt.Println("t2=", t2)
}

func Merge(src interface{}, target interface{}) {
	s := reflect.ValueOf(src).Elem()
	t := reflect.ValueOf(target).Elem()

	typeOfSrc := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		f2 := t.Field(i)

		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfSrc.Field(i).Name, f.Type(), f.Interface())
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfSrc.Field(i).Name, f2.Type(), f2.Interface())

		f2.Set(reflect.Value(f))
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfSrc.Field(i).Name, f2.Type(), f2.Interface())

	}

	fmt.Println("src=", )
	fmt.Println("target=", target)
}

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
	type T2 struct {
		B string
		A int
	}
	t := &T{23, "skidoo"}
	t2 := &T2{}
	Merge(t, t2)
	fmt.Println("src =", t)
	fmt.Println("target =", t2)
}

// Merge field with same type and name from source struct to target struct 
func Merge(src interface{}, target interface{}) {
	s := reflect.ValueOf(src).Elem()
	t := reflect.ValueOf(target).Elem()

	typeOfSrc := s.Type()
	typeOfTarget := t.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		for j := 0; j < t.NumField(); j++ {
			f2 := t.Field(j)

			if typeOfSrc.Field(i).Name == typeOfTarget.Field(j).Name {
				if f.Type() == f2.Type() {
					f2.Set(reflect.Value(f))
				}
			}
		}
	}

}

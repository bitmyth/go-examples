package main

import (
	"fmt"
	"go.marzhillstudios.com/pkg/go-html-transform/html/transform"
)


func main() {
	tree, _ := h5.New(rdr)
	t := transform.New(Tree)
	t.Apply(CopyAnd(myModifiers...), "li.menuitem")
	t.Apply(Replace(Text("my new text"), "a")
	fmt.Printf("%v",t)
}

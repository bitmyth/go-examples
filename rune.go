// https://www.asciitable.xyz/golang-program-convert-string-character-to-ascii-code/
package main

import(
	"fmt"
)

func main(){
	c := 'A' // rune (characters in Go are represented using `rune` data type)
	asciiValue := int(c)

	fmt.Printf("Ascii Value of %c = %d\n", c, asciiValue)

	asciiValue = 97
	character := rune(asciiValue)

	fmt.Printf("Character corresponding to Ascii Code %d = %c\n", asciiValue, character)

}

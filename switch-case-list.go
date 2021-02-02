package main

func main(){
	if WhiteSpace(32){
		println("is white space")
	} else {
		println("is not white space")
	}

}
func WhiteSpace(c rune) bool {
    switch c {
    case ' ', '\t', '\n', '\f', '\r':
        return true
    }
    return false
}

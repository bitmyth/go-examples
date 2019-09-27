package main
import "math"

func main(){
	var (
		a int64=2
		b int =2
	)
	println("a == b ",a==int64(b))
	//	println("a == b ",a==b) // invalid operation: a == b (mismatched types int64 and int)
	total := int(math.Ceil(float64(8 / 3)))
	println(total)
}

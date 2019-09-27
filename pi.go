package main

func main() {
	odd := 1.0
	result := 1.0
	sign := 1.0
	for i := 0; i < 100000000; i++ {
		odd += 2
		sign = -sign
		result += sign * (1 / odd)
	}
	println(result * 4)
}

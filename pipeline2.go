package main

func gen(nums []int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

func sq(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

func odd(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            if n%2 != 0 {
                out <- n
            }
        }
        close(out)
    }()
    return out
}

func sum(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        var sum = 0
        for n := range in {
            sum += n
        }
        out <- sum
        close(out)
    }()
    return out
}

type GenFunc func([]int) <-chan int
type PipeFunc func(<-chan int) <-chan int

func pipeline(nums []int, gen GenFunc, pipeFns ... PipeFunc) <-chan int {
    ch := gen(nums)
    for i := range pipeFns {
        ch = pipeFns[i](ch)
    }
    return ch
}

func main() {
    var nums = []int{1, 2, 3, 4, 5, 6, 7, 8}

    // first version
    for n := range sum(sq(odd(gen(nums)))) {
        println(n)
    }

    // second version
    for n := range pipeline(nums, gen, odd, sq, sum) {
        println(n)
    }
}

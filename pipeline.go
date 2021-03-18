package main

import (
    "math"
    "sync"
)

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

func is_prime(value int) bool {
    for i := 2; i <= int(math.Floor(float64((value / 2)))); i++ {
        if value%i == 0 {
            return false
        }
    }
    return value > 1
}

func prime(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            if is_prime(n) {
                out <- n
            }
        }
        close(out)
    }()
    return out
}

func makeRange(min, max int) []int {
    a := make([]int, max-min+1)
    for i := range a {
        a[i] = min + i
    }
    return a
}

func main() {
    nums := makeRange(1, 100)
    in := gen(nums)

    const n = 5
    var chans [n]<-chan int

    for i := range chans {
        chans[i] = sum(prime(in))
    }

    for n := range sum(merge(chans[:])) {
        println(n)
    }
}

func merge(cs []<-chan int) <-chan int {
    var wg sync.WaitGroup
    out := make(chan int)

    // Start an output goroutine for each input channel in cs.  output
    // copies values from c to out until c is closed, then calls wg.Done.
    output := func(c <-chan int) {
        for n := range c {
            out <- n
        }
        wg.Done()
    }
    wg.Add(len(cs))
    for _, c := range cs {
        go output(c)
    }

    // Start a goroutine to close out once all the output goroutines are
    // done.  This must start after the wg.Add call.
    go func() {
        wg.Wait()
        close(out)
    }()
    return out
}

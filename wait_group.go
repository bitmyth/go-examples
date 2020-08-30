package main

func main(){
    wg :=sync.WaitGroup{}
    requests := make(chan int, 5)

    for i := 1; i <= 5; i++ {
        wg.Add(1)
        requests <- i
    }
    close(requests)

    for i := 1; i <= 5; i++ {
        go func() {
            req := <-requests
            wg.Done()
            fmt.Println("request", req, time.Now())

        }()
    }

    wg.Wait()
}

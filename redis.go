package main

import (
    "fmt"
	//"github.com/garyburd/redigo/redis"
	"github.com/gomodule/redigo/redis"
)

func main() {
    c, err := redis.Dial("tcp", "localhost:6379")
    if err != nil {
        fmt.Println("conn redis failed,", err)
        return
    }

    defer c.Close()
    _, err = c.Do("Set", "abc", 100)
    if err != nil {
        fmt.Println(err)
        return
    }

    r, err := redis.Int(c.Do("Get", "abc"))
    if err != nil {
        fmt.Println("get abc failed,", err)
        return
    }

    fmt.Println(r)
}

func Scan(){
    c, err := redis.Dial("tcp", "localhost:6379")
    if err != nil {
        fmt.Println("Connect to redis error", err)
        return
    }
    count := 0
    iter := 0

    var keys []string

    for {
        arr, err := redis.Values(c.Do("SCAN", iter, "MATCH", "*", "COUNT", 10000))
        if err != nil {
            log.Fatal(err)
        } else {
            iter, _ = redis.Int(arr[0], nil)
            keys, _ = redis.Strings(arr[1], nil)
        }
        count += len(keys)

        log.Printf("iter:%v", iter)
        log.Printf("count:%v", iter)

        // SCAN is a cursor based iterator. This means that at every call of the command,
        // the server returns an updated cursor that the user needs to use as the cursor argument in the next call.
        // An iteration starts when the cursor is set to 0, and terminates when the cursor returned by the server is 0.
        // The following is an example of SCAN iteration:
        if iter == 0 {
            break
        }

        for _, val := range keys {
            log.Printf("val :%v", val)
        }
    }
}

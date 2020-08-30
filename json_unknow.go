package main

import (
    "encoding/json"
    "fmt"
    "strings"
)

func UnknownJson(data string) {
    if data != `` {
        r := strings.NewReader(data)
        dec := json.NewDecoder(r)
        switch data[0] {
        case 91:
            // "[" 开头的Json
            param := []interface{}{}
            dec.Decode(&param)
            fmt.Println(param)
        case 123:
            // "{" 开头的Json
            param := make(map[string]interface{})
            dec.Decode(&param)
            fmt.Println(param)
        }
    }
}

func main() {
    UnknownJson(`{"a":1}`)
    UnknownJson(`[{"a":1},{"b":2}]`)
}

package main

import (
    "bufio"
    "fmt"
    "os/exec"
)

func main() {

    cmd := exec.Command("sh", "-c", `for number in {0..10}; do echo "$number ";sleep 1; done;`)
    pipe, _ := cmd.StdoutPipe()
    if err := cmd.Start(); err != nil {
        // handle error
    }
    reader := bufio.NewReader(pipe)
    line, err := reader.ReadString('\n')
    for err == nil {
        fmt.Println(line)
        line, err = reader.ReadString('\n')
    }
    cmd.Wait()
}

// https://zupzup.org/go-port-forwarding/

// nc -l 0.0.0.0 8888
// go run main.go --target 127.0.0.1:8888 --port 1337
// nc localhost 1337

package main

import (
    "flag"
    "fmt"
    "io"
    "log"
    "net"
    "os"
    "os/signal"
)
// First off, we use the flag package to parse the incoming target and port parameters.

var (
    target string
    port   int
)

func init() {
    flag.StringVar(&target, "target", "", "target (<host>:<port>)")
    flag.IntVar(&port, "port", 1337, "port")
}

func main() {
    flag.Parse()
    println(target)
   // Then, we start a server on the given port, so clients can connect to our service (1337 and 1338 in the above example).

    incoming, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
    if err != nil {
        log.Fatalf("could not start server on %d: %v", port, err)
    }
    fmt.Printf("server running on %d\n", port)
// We listen for incoming connections and accept the first one, creating our client. This version only works for one connected client, but it wouldn’t be hard to change this to a multi-client solution by creating a goroutine for each connected client.

    client, err := incoming.Accept()
    if err != nil {
        log.Fatal("could not accept client connection", err)
    }
    defer client.Close()
    fmt.Printf("client '%v' connected!\n", client.RemoteAddr())
// Finally, we connect to the specified target server (nginx and the remote service in the example).

    target, err := net.Dial("tcp", target)
    if err != nil {
        log.Fatal("could not connect to target", err)
    }
    defer target.Close()
    fmt.Printf("connection to server %v established!\n", target.RemoteAddr())
// Alright, we have a server which accepts incoming connections and is connected to the target server. Now we need a way to move the traffic between client and target.


    signals := make(chan os.Signal, 1)
    stop := make(chan bool)
    signal.Notify(signals, os.Interrupt)
    go func() {
        for _ = range signals {
            fmt.Println("\nReceived an interrupt, stopping...")
                stop <- true
        }
    }()

// Luckily, both client and target are of type net.Conn and satisfy both the io.Reader and io.Writer interfaces. So, passing the incoming traffic to the outgoing connection and vice versa is trivial using io.Copy.

    go func() { io.Copy(target, client) }()
    go func() { io.Copy(client, target) }()
// If we would start the program like this however, it would instantly run through and quit after the first connection. For this purpose, we introduce a stop channel, which blocks at the end of the program and reacts to an os.Signal.

// This way the application will continue running until e.g.: CTRL+C is pressed.

    <-stop
}

// And that’s it.

// If we run this application as specified in the beginning of the post, with one local and one remote part and navigate to http://localhost:1338 we can see that the traffic is passed through successfully.
// 
// Although this is just a fairly useless HTTP proxy at this state, there are lots of things one could accomplish by extending this simple idea.
// 
// For example, you could monitor, multiplex, filter, manipulate or encrypt traffic… endless possibilities! :)
// 
// Conclusion
// This was another nice little example made possible by the power of Go’s Reader and Writer interfaces.
// 
// Also, this example only uses a bare minimum of functionality, all from the standard library, without being overly verbose, which speaks for Go’s networking primitives.

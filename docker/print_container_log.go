// https://docs.docker.com/engine/api/sdk/examples/
package main

import (
    "context"
    "io"
    "os"

    "github.com/docker/docker/api/types"
    "github.com/docker/docker/client"
)

func main() {
    ctx := context.Background()
    cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    if err != nil {
        panic(err)
    }

    options := types.ContainerLogsOptions{ShowStdout: true}
    
    out, err := cli.ContainerLogs(ctx, "e27cba78623bc30918", options)
    if err != nil {
        panic(err)
    }

    io.Copy(os.Stdout, out)
}

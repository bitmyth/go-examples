package main

import (
    "context"
    "fmt"

    "github.com/docker/docker/api/types"
    "github.com/docker/docker/api/types/container"
    "github.com/docker/docker/client"
)

func main() {
    ctx := context.Background()
    cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    if err != nil {
        panic(err)
    }

    createResp, err := cli.ContainerCreate(ctx, &container.Config{
        Image: "alpine",
        Cmd:   []string{"touch", "/helloworld"},
    }, nil, nil, nil,"")
    if err != nil {
        panic(err)
    }

    if err := cli.ContainerStart(ctx, createResp.ID, types.ContainerStartOptions{}); err != nil {
        panic(err)
    }

    statusCh, errCh := cli.ContainerWait(ctx, createResp.ID, container.WaitConditionNotRunning)
    select {
    case err := <-errCh:
        if err != nil {
            panic(err)
        }
    case <-statusCh:
    }

    commitResp, err := cli.ContainerCommit(ctx, createResp.ID, types.ContainerCommitOptions{Reference: "helloworld"})
    if err != nil {
        panic(err)
    }

    fmt.Println(commitResp.ID)
}

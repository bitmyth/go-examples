// https://docs.docker.com/engine/api/sdk/#sdk-and-api-quickstart
package main

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/docker/go-connections/nat"
)

func main() {
    ctx := context.Background()
    cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    if err != nil {
		println(err.Error())
        panic(err)
    }

	reader, err := cli.ImagePull(ctx, "docker.io/library/mysql:5.7", types.ImagePullOptions{})
    if err != nil {
        panic(err)
    }
    io.Copy(os.Stdout, reader)

	hostConfig := &container.HostConfig{
        PortBindings: map[nat.Port][]nat.PortBinding{
            "3306/tcp": {
                {"localhost", "3306"},
            },
        },
    }

    resp, err := cli.ContainerCreate(ctx, &container.Config{
        Image: "mysql:5.7",
        Cmd:   []string{"--character-set-server=utf8mb4", "--default-time-zone=+08:00"},
        Tty:   false,
        Env:   []string{"MYSQL_ROOT_PASSWORD=123", "MYSQL_DATABASE=accounts"},
    }, hostConfig, nil, nil, "")

    if err != nil {
        panic(err)
    }

    if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
        panic(err)
    }

    statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
    select {
    case err := <-errCh:
        if err != nil {
            panic(err)
        }
    case <-statusCh:
    }

    out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
    if err != nil {
        panic(err)
    }

    stdcopy.StdCopy(os.Stdout, os.Stderr, out)
}

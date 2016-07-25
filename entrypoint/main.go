package main

import (
	"fmt"
	"net/http"
	"github.com/bfirsh/go-dcgi"
	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/container"
	"golang.org/x/net/context"
	"github.com/docker/engine-api/types/filters"
)


func main() {
	cli, err := client.NewClient("unix:///var/run/docker.sock", "v1.23", nil, nil)
	if err != nil {
		panic(err)
	}

// Create a network if not created.

networkExists := false

networkResource, err := cli.NetworkList(context.Background(),  types.NetworkListOptions {
Filters: filters.NewArgs(),
})
for _, val := range networkResource {
	if val.Name == "serverlessdockervotingapp_default" {
		networkExists = true
		fmt.Println("Network already created!")
	}
}
var networkId string;

if networkExists == false {
	networkResponse, err := cli.NetworkCreate(context.Background(), "serverlessdockervotingapp_default", types.NetworkCreate{
			CheckDuplicate: true,
			Driver:         "bridge",
			EnableIPv6:     true,
			Internal:       true,
			Options: map[string]string{
				"opt-key": "opt-value",
			},
		})
		if err != nil {
			panic(err)
		}
	networkId = networkResponse.ID
}

hostConfig := &container.HostConfig{
	NetworkMode: container.NetworkMode(networkId),
	Binds:       []string{"/var/run/docker.sock:/var/run/docker.sock"},
}

//Create a db container if not created

containerList, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
		Size:   true,
		All:    true,
		Since:  "container",
		Filter: filters.NewArgs(),
})

dbExists := false
for _, val := range containerList {
	if len(val.Names) == 1 {
		if val.Names[0] == "/db" {
			dbExists = true
			fmt.Printf("Container db exists\n")
		}
	}
}

if dbExists == false {
	containerResponse, err := cli.ContainerCreate(context.Background(), &container.Config{Hostname: "db", Image: "postgres:9.4"}, hostConfig, nil, "db")
	if err != nil {
		panic(err)
	}
	err = cli.ContainerStart(context.Background(), containerResponse.ID, types.ContainerStartOptions{CheckpointID: "checkpoint_id"})
	if err != nil {
		panic(err)
	}
}


http.Handle("/vote/", &dcgi.Handler{
	Image:      "bfirsh/serverless-vote",
	Client:     cli,
	HostConfig: hostConfig,
	Root:       "/vote", // strip /vote from all URLs
})
http.Handle("/result/", &dcgi.Handler{
	Image:      "bfirsh/serverless-result",
	Client:     cli,
	HostConfig: hostConfig,
	Root:       "/result",
})
fmt.Println(http.ListenAndServe(":8080", nil))

}

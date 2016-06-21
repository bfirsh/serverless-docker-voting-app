package main

import (
	"net/http"

	"github.com/bfirsh/go-dcgi"
	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types/container"
)

func main() {
	cli, err := client.NewClient("unix:///var/run/docker.sock", "v1.23", nil, nil)
	if err != nil {
		panic(err)
	}

	hostConfig := &container.HostConfig{
		NetworkMode: "serverlessdockervotingapp_default",
		Binds:       []string{"/var/run/docker.sock:/var/run/docker.sock"},
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
	http.ListenAndServe(":80", nil)
}

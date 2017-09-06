package docker

import (
	"context"
	"strings"
	"sync/atomic"
	"time"

	prompt "github.com/c-bata/go-prompt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func newContainerList(cli *client.Client) (*containerList, error) {
	return &containerList{
		cli: cli,
	}, nil
}

type containerList struct {
	list          atomic.Value
	lastFetchedAt time.Time

	cli *client.Client
}

func (c *containerList) fetchContainerList() {
	containers, err := c.cli.ContainerList(context.Background(), types.ContainerListOptions{})

	if err != nil {
		panic(err)
	}

	c.list.Store(containers)
}

func (c *containerList) getContainerSuggestions() []prompt.Suggest {
	go c.fetchContainerList()
	containers, ok := c.list.Load().([]types.Container)
	if !ok || len(containers) == 0 {
		return []prompt.Suggest{}
	}
	s := make([]prompt.Suggest, len(containers))
	for i, container := range containers {
		name := container.ID

		if len(container.Names) > 0 {
			name = strings.TrimLeft(container.Names[0], "/")
		}

		s[i] = prompt.Suggest{
			Text:        name,
			Description: container.Image,
		}
	}
	return s
}
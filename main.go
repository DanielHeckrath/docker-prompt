package main

import (
	"fmt"

	"github.com/c-bata/go-prompt"

	"github.com/DanielHeckrath/docker-prompt/docker"
	"github.com/docker/docker/client"
)

var (
	version  string
	revision string
)

func main() {
	cli, err := client.NewEnvClient()

	if err != nil {
		panic(err)
	}

	completer, err := docker.NewCompleter(cli)

	if err != nil {
		panic(err)
	}

	fmt.Printf("docker-prompt %s (rev-%s)\n", version, revision)
	fmt.Println("Please use `exit` or `Ctrl-D` to exit this program..")
	defer fmt.Println("Bye!")
	p := prompt.New(
		docker.Executor,
		completer.Complete,
		prompt.OptionTitle("docker-prompt: interactive docker client"),
		prompt.OptionPrefix(">>> "),
		prompt.OptionInputTextColor(prompt.Yellow),
	)
	p.Run()
}
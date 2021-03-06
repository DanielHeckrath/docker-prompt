package docker

import (
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
)

type Completer struct {
	containers *containerList
}

func NewCompleter(cli *client.Client) (*Completer, error) {
	containers, err := newContainerList(cli)

	if err != nil {
		return nil, errors.Wrap(err, "unable to create container list")
	}

	return &Completer{
		containers: containers,
	}, nil
}

func (c *Completer) Complete(d prompt.Document) []prompt.Suggest {
	if d.TextBeforeCursor() == "" {
		return []prompt.Suggest{}
	}
	args := strings.Split(d.TextBeforeCursor(), " ")
	// w := d.GetWordBeforeCursor()

	// If PIPE is in text before the cursor, returns empty suggestions.
	for i := range args {
		if args[i] == "|" {
			return []prompt.Suggest{}
		}
	}

	// If word before the cursor starts with "-", returns CLI flag options.
	// if strings.HasPrefix(w, "-") {
	// 	return optionCompleter(args, strings.HasPrefix(w, "--"))
	// }

	return c.completeArguments(excludeOptions(args))
}

var commands = []prompt.Suggest{
	{Text: "attach", Description: "Attach local standard input, output, and error streams to a running container"},
	{Text: "build", Description: "Build an image from a Dockerfile"},
	{Text: "commit", Description: "Create a new image from a container's changes"},
	{Text: "cp", Description: "Copy files/folders between a container and the local filesystem"},
	{Text: "deploy", Description: "Deploy a new stack or update an existing stack"},
	{Text: "diff", Description: "Inspect changes to files or directories on a container's filesystem"},
	{Text: "events", Description: "Get real time events from the server"},
	{Text: "exec", Description: "Run a command in a running container"},
	{Text: "export", Description: "Export a container's filesystem as a tar archive"},
	{Text: "history", Description: "Show the history of an image"},
	{Text: "images", Description: "List images"},
	{Text: "import", Description: "Import the contents from a tarball to create a filesystem image"},
	{Text: "info", Description: "Display system-wide information"},
	{Text: "inspect", Description: "Return low-level information on Docker objects"},
	{Text: "kill", Description: "Kill one or more running containers"},
	{Text: "load", Description: "Load an image from a tar archive or STDIN"},
	{Text: "login", Description: "Log in to a Docker registry"},
	{Text: "logout", Description: "Log out from a Docker registry"},
	{Text: "logs", Description: "Fetch the logs of a container"},
	{Text: "pause", Description: "Pause all processes within one or more containers"},
	{Text: "port", Description: "List port mappings or a specific mapping for the container"},
	{Text: "ps", Description: "List containers"},
	{Text: "pull", Description: "Pull an image or a repository from a registry"},
	{Text: "push", Description: "Push an image or a repository to a registry"},
	{Text: "rename", Description: "Rename a container"},
	{Text: "restart", Description: "Restart one or more containers"},
	{Text: "rm", Description: "Remove one or more containers"},
	{Text: "rmi", Description: "Remove one or more images"},
	{Text: "run", Description: "Run a command in a new container"},
	{Text: "save", Description: "Save one or more images to a tar archive (streamed to STDOUT by default)"},
	{Text: "search", Description: "Search the Docker Hub for images"},
	{Text: "start", Description: "Start one or more stopped containers"},
	{Text: "stats", Description: "Display a live stream of container(s) resource usage statistics"},
	{Text: "stop", Description: "Stop one or more running containers"},
	{Text: "tag", Description: "Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE"},
	{Text: "top", Description: "Display the running processes of a container"},
	{Text: "unpause", Description: "Unpause all processes within one or more containers"},
	{Text: "update", Description: "Update configuration of one or more containers"},
	{Text: "version", Description: "Show the Docker version information"},
	{Text: "wait", Description: "Block until one or more containers stop, then print their exit codes"},

	// Custom command.
	{Text: "exit", Description: "Exit this program"},
}

func (c *Completer) completeArguments(args []string) []prompt.Suggest {
	if len(args) <= 1 {
		return prompt.FilterHasPrefix(commands, args[0], true)
	}

	first := args[0]
	second := args[1]
	switch first {
	case "logs":
		if len(args) == 2 {
			return prompt.FilterContains(c.containers.getContainerSuggestions(), second, true)
		}
	default:
		return []prompt.Suggest{}
	}
	return []prompt.Suggest{}
}
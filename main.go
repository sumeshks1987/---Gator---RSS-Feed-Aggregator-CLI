package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: gator <command> [args]")
		os.Exit(1)
	}

	cmd := command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	commands := registerCommands()

	handler, ok := commands[cmd.Name]
	if !ok {
		fmt.Println("unknown command")
		os.Exit(1)
	}

	state, err := NewState()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := handler(state, cmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

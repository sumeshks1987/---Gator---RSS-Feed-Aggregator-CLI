package main

import (
	"fmt"
	"log"
	"os"

	"rss/internal/web"
)

func main() {
	state, err := NewState()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Start web server
	if len(os.Args) > 1 && os.Args[1] == "serve" {
		server := web.New(state.DB)
		log.Fatal(server.Start(":8080"))
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("usage: gator <command>")
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

	if err := handler(state, cmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

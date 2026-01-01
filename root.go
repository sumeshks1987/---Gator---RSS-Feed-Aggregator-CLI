package main

type command struct {
	Name string
	Args []string
}

func registerCommands() map[string]func(*State, command) error {
	return map[string]func(*State, command) error{
		"register": Register,
		"login":    Login,
		"reset":    Reset,
		"agg":      Agg,
		"feeds":    Feeds,

		// Middleware-protected commands
		"addfeed":   middlewareLoggedIn(AddFeed),
		"follow":    middlewareLoggedIn(Follow),
		"following": middlewareLoggedIn(Following),
		"unfollow":  middlewareLoggedIn(Unfollow),
		"browse":    middlewareLoggedIn(Browse),
	}
}

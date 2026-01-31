package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Manhnguyen981024/blog-aggregator/internal/config"
	"github.com/Manhnguyen981024/blog-aggregator/internal/database"
	"github.com/Manhnguyen981024/blog-aggregator/internal/handler"
	_ "github.com/lib/pq"
)

func main() {
	conf, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	db, err := sql.Open("postgres", conf.DbURL)
	if err != nil {
		log.Fatal("error reading config")
	}
	defer db.Close()
	dbQueries := database.New(db)

	state := config.State{
		DB:     dbQueries,
		Config: &conf,
	}

	commands := config.Commands{
		Cmds: map[string]func(*config.State, config.Command) error{},
	}

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Error: You must enter at least 2 arguments")
		os.Exit(1)
		return
	}

	cmd := config.Command{
		Name:      args[1],
		Arguments: args[2:],
	}
	commands.Register("login", handler.HandlerLogin)
	commands.Register("register", handler.HandlerRegister)
	commands.Register("reset", handler.HandlerDelete)
	commands.Register("users", handler.HandlerUsers)
	commands.Register("agg", middlewareLoggedIn(handler.HandlerAgg))
	commands.Register("addfeed", middlewareLoggedIn(handler.HandlerAddFeed))
	commands.Register("feeds", handler.HandlerFeeds)
	commands.Register("follow", middlewareLoggedIn(handler.HandlerFollow))
	commands.Register("following", middlewareLoggedIn(handler.HandlerFollowing))
	commands.Register("unfollow", middlewareLoggedIn(handler.HandlerUnfollow))
	commands.Register("browse", middlewareLoggedIn(handler.HanlderBrowse))

	errCmd := commands.Run(&state, cmd)
	if errCmd != nil {
		fmt.Println(errCmd)
		os.Exit(1)
	}
}

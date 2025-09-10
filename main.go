package main

import (
	"log"
	"os"
	"database/sql"

	"github.com/dulchik/blog_aggregator/internal/config"
	"github.com/dulchik/blog_aggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)


	s := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	c := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	c.register("login", handlerLogin)
	c.register("register", handlerRegister)
	c.register("reset", handlerReset)
	c.register("users", handlerListUsers)
	c.register("agg", handlerAgg)
	c.register("addfeed", handlerAddFeed)
	c.register("feeds", handlerListFeeds)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = c.run(s, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}


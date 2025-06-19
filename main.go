package main

import (
	"database/sql"

	_ "github.com/lib/pq"

	"fmt"
	"os"

	"github.com/a-wayne/gator/internal/config"
	"github.com/a-wayne/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	commandMap map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	f := c.commandMap[cmd.name]
	if f == nil {
		return fmt.Errorf("invalid command: '%s'", cmd.name)
	}

	err := f(s, cmd)
	return err
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandMap[name] = f
}

func main() {
	c, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s := state{}
	s.cfg = &c

	db, err := sql.Open("postgres", c.DBURL)
	if err != nil {
		fmt.Println("error connecting to database")
		os.Exit(1)
	}
	dbQueries := database.New(db)
	s.db = dbQueries

	cmds := commands{}
	cmds.commandMap = make(map[string]func(*state, command) error)

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("feeds", handlerFeeds)
	cmds.register("addfeed", middleWareLoggedIn(handlerAddfeed))
	cmds.register("follow", middleWareLoggedIn(handlerFollow))
	cmds.register("following", middleWareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middleWareLoggedIn(handlerUnfollow))
	cmds.register("browse", middleWareLoggedIn(handlerBrowse))

	if len(os.Args) < 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	var cmdArgs []string
	if len(os.Args) > 2 {
		cmdArgs = os.Args[2:]
	}

	cmd := command{
		name: cmdName,
		args: cmdArgs,
	}

	err = cmds.run(&s, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

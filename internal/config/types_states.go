package config

import (
	"errors"
	"fmt"

	"github.com/Manhnguyen981024/blog-aggregator/internal/database"
)

type State struct {
	DB     *database.Queries
	Config *Config
}

type Command struct {
	Name      string
	Arguments []string
}

type Commands struct {
	Cmds map[string]func(*State, Command) error
}

func (c Commands) Run(s *State, cmd Command) error {
	f, ok := c.Cmds[cmd.Name]
	if !ok {
		return errors.New("Command not found!")
	}
	err := f(s, cmd)
	if err != nil {
		return err
	}
	fmt.Println("Run successful!!")

	return nil
}

func (c Commands) Register(name string, f func(*State, Command) error) {
	c.Cmds[name] = f
}

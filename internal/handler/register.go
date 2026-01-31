package handler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Manhnguyen981024/blog-aggregator/internal/config"
	"github.com/Manhnguyen981024/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func HandlerRegister(s *config.State, cmd config.Command) error {
	if len(cmd.Arguments) == 0 {
		return errors.New("the register handler expects a single argument, the username.")
	}

	username := cmd.Arguments[0]
	nt := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	newUser := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: nt,
		UpdatedAt: nt,
		Name:      username,
	}
	ctx := context.Background()

	v, err := s.DB.CreateUser(ctx, newUser)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s.Config.SetUser(v.Name)
	fmt.Printf("Register user %v successfully!", v)

	return nil
}

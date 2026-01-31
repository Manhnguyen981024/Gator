package handler

import (
	"context"
	"fmt"
	"os"

	"github.com/Manhnguyen981024/blog-aggregator/internal/config"
)

func HandlerDelete(s *config.State, cmd config.Command) error {

	err := s.DB.DeleteAllUsers(context.Background())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return nil
}

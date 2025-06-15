package main

import (
	"fmt"
	"os"

	"github.com/blacksails/rocket-messages/pkg/server"
)

func main() {
	s := server.New()
	if err := s.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

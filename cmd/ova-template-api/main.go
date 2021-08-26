package main

import (
	"fmt"
	"github.com/ozonva/ova-game-api/internal/server"
	"github.com/rs/zerolog/log"
)

func main() {
	fmt.Println("Hello from ova-game-api. 👋")

	if err := server.Run(); err != nil {
		log.Fatal().Err(err)
	}
}

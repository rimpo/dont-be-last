package main

import (
	"github.com/rimpo/dont-be-last/client/app"
	"github.com/rimpo/dont-be-last/pkg/service/cli/auth"
	"github.com/rimpo/dont-be-last/pkg/service/cli/game"
)

func main() {
	authService := auth.NewService("http://127.0.0.1:5000/auth")
	gameService := game.NewService()

	application := app.New(authService, gameService)

	application.Run()
}

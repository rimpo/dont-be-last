package cli

import "github.com/rimpo/dont-be-last/pkg/api"

type GameService interface {
	GetCurrentState(gameId int) *api.GameState
	Move(gameId int, move []int) *api.GameState
}

type AuthService interface {
	Login(userName, password string) (string, int, error)
}

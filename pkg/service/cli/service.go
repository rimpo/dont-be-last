package cli

import "github.com/rimpo/dont-be-last/pkg/api"

type GameService interface {
	func GetCurrentState(gameId int) *api.GameState
	func Move(gameId int,move []int) *api.GameState
}

type AuthService interface {
	func Login(userName, password string) (string, string, error)
}

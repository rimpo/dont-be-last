package game

import (
	"github.com/rimpo/dont-be-last/pkg/api"
)

type service struct {
}

func NewService() *service {
	return &service{}
}

func (g *service) GameState(gameId int) *api.GameState {
	return nil
}

func (g *service) Move(gameId int, move []int) *api.GameState {
	return nil
}

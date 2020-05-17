package api

type LoginResponse struct {
	Id int64 `json: "id"`
}

type PlayerInfo struct {
	Id       int64  `json: "id"`
	Name     string `json: "name"`
	IsTurn   bool   `json: "is_turn"`
	IsWinner bool   `json: "is_winner"`
}

type Move struct {
	Id       int64   `json: "id"`
	PlayerId int64   `json: "player_id"`
	Move     []int64 `json: "move"`
}
type GameState struct {
	Id      int64 `json: "id"`
	Board   []int `json: "board"`
	Players struct {
		Own      PlayerInfo `json: "own"`
		Opponent PlayerInfo `json: "opponent"`
	} `json: "players"`
}

type Player struct {
	Id   int64  `json: "id"`
	Name string `json: "name"`
}

type GameWaiting struct {
	Id       int64  `json: "id"`
	Opponent Player `json: "opponent"`
}

type GameWaitingList struct {
	GameWaitings []GameWaiting `json:"games_waiting"`
}

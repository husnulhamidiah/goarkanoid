package constant

type GameState int

const (
	Initial GameState = iota
	Playing
	GameOver
)

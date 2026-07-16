package main

type GameState int

const (
	Initial GameState = iota
	Playing
	Spawning
	GameOver
)

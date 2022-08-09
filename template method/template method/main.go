package main

import "fmt"

type Game interface {
	Start()
	TakeTurn()
	HaveWinner() bool
	WinningPlayer() int
}

type chess struct {
	turn, maxTurns, currentPlayer int
}

func (c *chess) Start() {
	fmt.Println("Starting a new game of chess")
}

func (c *chess) TakeTurn() {
	c.turn++
	fmt.Printf("Turn %d taken by player %d\n", c.turn, c.currentPlayer)
	c.currentPlayer = 1 - c.currentPlayer
}

func (c *chess) HaveWinner() bool {
	return c.turn == c.maxTurns
}

func (c *chess) WinningPlayer() int {
	return c.currentPlayer
}

func NewChess() Game {
	return &chess{
		turn:          1,
		maxTurns:      10,
		currentPlayer: 0,
	}
}

func PlayGame(game Game) {
	game.Start()
	for !game.HaveWinner() {
		game.TakeTurn()
	}
	fmt.Printf("Player %d wins.\n", game.WinningPlayer())
}

func main() {
	PlayGame(NewChess())
}

package main

import (
	"log"

	"github.com/saturnengine/saturn-engine"
)

type Game struct{}

func (g *Game) Update() error {
	log.Println("Update called")
	return nil
}

func (g *Game) Draw() error {
	log.Println("Draw called")
	return nil
}

func main() {
	g := &Game{}

	if err := saturn.RunGame(g, saturn.WithTPS(60), saturn.WithFPS(120)); err != nil {
		log.Fatal(err)
	}
}

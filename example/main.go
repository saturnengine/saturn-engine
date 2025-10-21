package main

import (
	"log"

	"github.com/saturnengine/saturn-engine"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw() error {
	// log.Println("Draw called")
	return nil
}

func main() {
	g := &Game{}
	instance, err := saturn.NewInstance(g, saturn.WithFPS(60), saturn.WithFPS(120), saturn.WithWindowTitle("Example"))
	if err != nil {
		log.Fatal(err)
	}

	if err := instance.RunGame(); err != nil {
		log.Fatal(err)
	}
}

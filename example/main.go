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
	i, err := saturn.NewInstance(g, saturn.WithFPS(60), saturn.WithFPS(120))
	if err != nil {
		log.Fatal(err)
	}

	if err := i.RunGame(); err != nil {
		log.Fatal(err)
	}
}

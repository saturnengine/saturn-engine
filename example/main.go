package main

import (
	"fmt"
	"log"

	"github.com/saturnengine/saturn-engine"
)

type Game struct {
	count int
}

func (g *Game) Update() error {
	saturn.SetTitle(fmt.Sprintf("tick count: %d", g.count))
	g.count++
	return nil
}

func (g *Game) Draw() error {
	return nil
}

func main() {
	g := &Game{}
	err := saturn.NewInstance(g, saturn.WithFPS(60), saturn.WithFPS(120), saturn.WithWindowTitle("Example"))
	if err != nil {
		log.Fatal(err)
	}

	if err := saturn.RunGame(); err != nil {
		log.Fatal(err)
	}
}

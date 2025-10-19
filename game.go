package saturn

// Game represents a game.
type Game interface {
	// Update proceeds the game state.
	// Update is called every tick (1/60 [s] by default).
	Update() error
	// Draw draws the game screen.
	// Draw is called every frame (1/60 [s] by default), but the actual rate may vary
	// depending on the display's refresh rate or system performance.
	Draw() error
}

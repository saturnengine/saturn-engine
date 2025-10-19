package saturn

import "time"

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

// RunGame runs the given game.
func RunGame(game Game, opts ...gameOption) (err error) {
	// Apply options
	options := defaultGameOptions()
	applyGameOptions(options, opts)
	if err = options.validate(); err != nil {
		return
	}

	// Make channels
	errCh := make(chan error, 2)
	quit := make(chan struct{})

	// Update loop
	go func() {
		ticker := time.NewTicker(time.Second / time.Duration(options.tps))
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				if err := (game).Update(); err != nil {
					errCh <- err
					return
				}
			case <-quit:
				return
			}
		}
	}()

	// Draw loop
	go func() {
		ticker := time.NewTicker(time.Second / time.Duration(options.fps))
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				if err := (game).Draw(); err != nil {
					errCh <- err
					return
				}
			case <-quit:
				return
			}
		}
	}()

	// Wait for an error
	err = <-errCh
	close(quit)
	return err
}

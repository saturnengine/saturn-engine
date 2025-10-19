package saturn

import "time"

type Instance struct {
	game    Game
	options *gameOptions
}

// NewInstance creates a new game instance with the given game and options.
func NewInstance(g Game, opts ...gameOption) (i *Instance, err error) {
	options := defaultGameOptions()
	applyGameOptions(options, opts)
	if err = options.validate(); err != nil {
		return
	}
	i = &Instance{
		game:    g,
		options: options,
	}
	return
}

// RunGame runs the game instance.
func (i *Instance) RunGame() (err error) {
	// Make channels
	errCh := make(chan error, 2)
	quit := make(chan struct{})

	// Update loop
	go func() {
		ticker := time.NewTicker(time.Second / time.Duration(i.options.tps))
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				if err := i.game.Update(); err != nil {
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
		ticker := time.NewTicker(time.Second / time.Duration(i.options.fps))
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				if err := i.game.Draw(); err != nil {
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
	return
}

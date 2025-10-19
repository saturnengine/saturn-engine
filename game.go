package saturn

import "time"

// Game represents a game.
type Game interface {
	// Update proceeds the game state.
	// Update is called every tick (1/60 [s] by default).
	Update() error
	// Draw draws the game screen.
	// Draw is called every frame (typically 1/60[s] for 60Hz display).
	Draw() error
}

// GameOptions represents options for running a game.
type gameOptions struct {
	tps int // ticks per second
	fps int // frames per second
}

// defaultGameOptions returns the default game options.
func defaultGameOptions() *gameOptions {
	return &gameOptions{
		tps: 60,
		fps: 60,
	}
}

// gameOption is an interface for applying options to gameOptions.
type gameOption interface {
	apply(*gameOptions)
}

// gameOptionFunc is a function type that implements gameOption.
type gameOptionFunc func(*gameOptions)

// apply applies the gameOptionFunc to the given gameOptions.
func (f gameOptionFunc) apply(r *gameOptions) { f(r) }

// WithTPS sets the ticks per second for the game.
func WithTPS(tps int) gameOption {
	return gameOptionFunc(func(r *gameOptions) {
		r.tps = tps
	})
}

// WithFPS sets the frames per second for the game.
func WithFPS(fps int) gameOption {
	return gameOptionFunc(func(r *gameOptions) {
		r.fps = fps
	})
}

// RunGame runs the given game.
func RunGame(game Game, opts ...gameOption) (err error) {
	// Apply options
	options := defaultGameOptions()
	for _, opt := range opts {
		opt.apply(options)
	}
	if options.tps <= 0 {
		err = ErrInvalidTPS
		return
	}
	if options.fps <= 0 {
		err = ErrInvalidFPS
		return
	}

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

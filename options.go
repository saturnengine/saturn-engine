package saturn

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

// applyGameOptions applies a list of gameOption to the given gameOptions.
func applyGameOptions(before *gameOptions, opts []gameOption) {
	for _, o := range opts {
		o.apply(before)
	}
}

// validate validates the gameOptions.
func (g *gameOptions) validate() error {
	if g.tps <= 0 {
		return ErrInvalidTPS
	}
	if g.fps <= 0 {
		return ErrInvalidFPS
	}
	return nil
}

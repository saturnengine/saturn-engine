package saturn

// GameOptions represents options for running a game.
type gameOptions struct {
	tps         int    // ticks per second
	fps         int    // frames per second
	width       int    // screen width
	height      int    // screen height
	windowTitle string // window title
}

// defaultGameOptions returns the default game options.
func defaultGameOptions() *gameOptions {
	return &gameOptions{
		tps:         60,
		fps:         60,
		width:       640,
		height:      480,
		windowTitle: "Saturn Engine",
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

// WithScreenSize sets the screen width and height for the game.
func WithScreenSize(width, height int) gameOption {
	return gameOptionFunc(func(r *gameOptions) {
		r.width = width
		r.height = height
	})
}

// WithWindowTitle sets the window title for the game.
func WithWindowTitle(title string) gameOption {
	return gameOptionFunc(func(r *gameOptions) {
		r.windowTitle = title
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
	if g.width < 0 {
		return ErrInvalidWidth
	}
	if g.height < 0 {
		return ErrInvalidHeight
	}
	return nil
}

package saturn

import (
	"runtime"
	"time"

	"github.com/saturnengine/saturn-engine/internal/glfw"
)

// Instance represents a game instance with its configuration.
type Instance struct {
	game    Game
	options *gameOptions
	window  glfw.Window
	closed  bool
}

// NewInstance creates a new game instance with the given game and options.
//
// This function locks the current OS thread because GLFW requires
// all window operations to occur on the same thread (especially on macOS/Cocoa).
func NewInstance(g Game, opts ...gameOption) (i *Instance, err error) {
	// Lock only on macOS
	if runtime.GOOS == "darwin" {
		runtime.LockOSThread()
	}

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
	i.window, err = glfw.NewWindow(i.options.width, i.options.height, i.options.windowTitle)
	if err != nil {
		return
	}

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

	// Main loop
	for {
		glfw.PollEvents()

		if i.window.ShouldClose() {
			close(quit)
			i.Close()
			return nil
		}

		select {
		case err = <-errCh:
			close(quit)
			i.Close()
			return
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}

// Close gracefully shuts down the game instance and releases resources.
//
// It destroys the GLFW window, terminates GLFW, and unlocks the OS thread.
func (i *Instance) Close() {
	if i.closed {
		return
	}
	i.closed = true

	if i.window != nil {
		glfw.DestroyWindow(i.window)
		i.window = nil
	}

	glfw.Terminate()
	if runtime.GOOS == "darwin" {
		runtime.UnlockOSThread()
	}
}

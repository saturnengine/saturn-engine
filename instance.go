package saturn

import (
	"runtime"
	"time"

	"github.com/saturnengine/saturn-engine/internal/glfw"
)

// Instance represents a game instance with its configuration.
type Instance struct {
	game           Game
	options        *gameOptions
	window         glfw.Window
	operationQueue chan Operation
	closed         bool
}

// instance is a singleton instance of the game.
var instance *Instance

// NewInstance creates a new game instance with the given game and options.
//
// This function locks the current OS thread because GLFW requires
// all window operations to occur on the same thread (especially on macOS/Cocoa).
func NewInstance(g Game, opts ...gameOption) (err error) {
	// Lock only on macOS
	if runtime.GOOS == "darwin" {
		runtime.LockOSThread()
	}

	options := defaultGameOptions()
	applyGameOptions(options, opts)
	if err = options.validate(); err != nil {
		return
	}

	instance = &Instance{
		game:           g,
		options:        options,
		operationQueue: make(chan Operation, 128), // Buffered channel for operations
	}
	return
}

// RunGame runs the game instance.
func RunGame() (err error) {
	if instance == nil {
		err = ErrInstanceNotCreated
		return
	}
	instance.window, err = glfw.NewWindow(instance.options.width, instance.options.height, instance.options.windowTitle)
	if err != nil {
		return
	}

	errCh := make(chan error, 2)
	quit := make(chan struct{})

	// Update loop
	go func() {
		ticker := time.NewTicker(time.Second / time.Duration(instance.options.tps))
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				if err := instance.game.Update(); err != nil {
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
		ticker := time.NewTicker(time.Second / time.Duration(instance.options.fps))
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				if err := instance.game.Draw(); err != nil {
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

		// Process all queued operations
		for {
			select {
			case op := <-instance.operationQueue:
				op()
			default:
				// If there are no more operations, break the loop
				goto DoneQueue
			}
		}
	DoneQueue:
		if instance.window.ShouldClose() {
			close(quit)
			Close()
			return nil
		}

		select {
		case err = <-errCh:
			close(quit)
			Close()
			return
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}

// Close gracefully shuts down the game instance and releases resources.
//
// It destroys the GLFW window, terminates GLFW, and unlocks the OS thread.
func Close() {
	if instance.closed {
		return
	}
	instance.closed = true

	if instance.window != nil {
		glfw.DestroyWindow(instance.window)
		instance.window = nil
	}

	glfw.Terminate()
	if runtime.GOOS == "darwin" {
		runtime.UnlockOSThread()
	}
}

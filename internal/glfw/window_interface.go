package glfw

// Window represents a window with its properties.
type Window interface {
	// Title of the window.
	Title() string
	SetTitle(title string)
	// Size of the window.
	Width() int
	Height() int
	SetSize(width int, height int)
	SizeLimits() (minw, minh, maxw, maxh int)
	SetSizeLimits(minw, minh, maxw, maxh int)
	// Position of the window.
	Position() (x, y int)
	SetPosition(x, y int)
	// Floating state of the window.
	IsFloating() bool
	// Maximized state of the window.
	IsMaximized() bool
	Maximize()
	// Minimized state of the window.
	IsMinimized() bool
	Minimize()
	// ShouldClose indicates whether the window should close.
	ShouldClose() bool
}

// newWindow creates a new window with the specified options.
func NewWindow(width, height int, title string) (w Window, err error) {
	w, err = newWindow(width, height, title)
	return
}

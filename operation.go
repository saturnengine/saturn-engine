package saturn

type Operation func()

func Title() string {
	return instance.window.Title()
}

func SetTitle(title string) {
	instance.operationQueue <- func() {
		instance.window.SetTitle(title)
	}
}

func Size() (width, height int) {
	return instance.window.Width(), instance.window.Height()
}

func SetSize(width, height int) {
	instance.operationQueue <- func() {
		instance.window.SetSize(width, height)
	}
}

func SizeLimits() (minw, minh, maxw, maxh int) {
	return instance.window.SizeLimits()
}

func SetSizeLimits(minw, minh, maxw, maxh int) {
	instance.operationQueue <- func() {
		instance.window.SetSizeLimits(minw, minh, maxw, maxh)
	}
}

func Position() (x, y int) {
	return instance.window.Position()
}

func SetPosition(x, y int) {
	instance.operationQueue <- func() {
		instance.window.SetPosition(x, y)
	}
}

func IsFloating() bool {
	return instance.window.IsFloating()
}

func IsMaximized() bool {
	return instance.window.IsMaximized()
}

func Maximize() {
	instance.operationQueue <- func() {
		instance.window.Maximize()
	}
}

func IsMinimized() bool {
	return instance.window.IsMinimized()
}

func Minimize() {
	instance.operationQueue <- func() {
		instance.window.Minimize()
	}
}

package opengl

// Context represents an OpenGL rendering context.
type Context struct {
	versionMajor int
	versionMinor int
}

func NewContext() *Context {
	return &Context{versionMajor: 3, versionMinor: 3}
}

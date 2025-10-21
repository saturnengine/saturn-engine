package glfw

import "errors"

var (
	ErrFailedToCreateWindow = errors.New("failed to create GLFW window")
	ErrFailedToInitGLFW     = errors.New("failed to initialize GLFW")
)

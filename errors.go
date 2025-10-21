package saturn

import "errors"

var (
	ErrInvalidTPS    = errors.New("invalid TPS value")
	ErrInvalidFPS    = errors.New("invalid FPS value")
	ErrInvalidWidth  = errors.New("invalid Width value")
	ErrInvalidHeight = errors.New("invalid Height value")
)

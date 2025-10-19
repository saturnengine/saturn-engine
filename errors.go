package saturn

import "errors"

var (
	ErrInvalidTPS = errors.New("invalid TPS value")
	ErrInvalidFPS = errors.New("invalid FPS value")
)

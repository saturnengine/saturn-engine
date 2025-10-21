//go:build darwin || freebsd || linux || netbsd || openbsd

package glfw

/*
#cgo darwin pkg-config: glfw3
#cgo darwin LDFLAGS: -framework Cocoa -framework OpenGL -framework IOKit -framework CoreVideo
#cgo linux freebsd netbsd openbsd pkg-config: glfw3
#include "glfw3_unix.h"
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"
)

// glfwWindow wraps a native GLFWwindow pointer.
type glfwWindow struct {
	handle *C.GLFWwindow
	title  string
}

// Terminate terminates GLFW when the program exits.
func Terminate() {
	C.glfwTerminate()
}

// newWindow creates a new GLFW window, makes its context current, and shows it.
func newWindow(width, height int, title string) (w *glfwWindow, err error) {
	// Initialize GLFW once
	if C.glfwInit() == C.GLFW_FALSE {
		err = ErrFailedToInitGLFW
		return
	}

	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))

	// Create a window
	window := C.glfwCreateWindow(C.int(width), C.int(height), ctitle, nil, nil)
	if window == nil {
		err = ErrFailedToCreateWindow
		return
	}

	// Make the OpenGL context current (must be done before drawing)
	C.glfwMakeContextCurrent(window)

	// Show the window (hidden by default until first event loop)
	C.glfwShowWindow(window)

	w = &glfwWindow{handle: window, title: title}
	return
}

func (w *glfwWindow) Title() string {
	return w.title
}

func (w *glfwWindow) SetTitle(title string) {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))
	C.glfwSetWindowTitle(w.handle, ctitle)
	w.title = title
}

func (w *glfwWindow) Width() int {
	var width, height C.int
	C.glfwGetWindowSize(w.handle, &width, &height)
	return int(width)
}

func (w *glfwWindow) Height() int {
	var width, height C.int
	C.glfwGetWindowSize(w.handle, &width, &height)
	return int(height)
}

func (w *glfwWindow) SetSize(width int, height int) {
	C.glfwSetWindowSize(w.handle, C.int(width), C.int(height))
}

func (w *glfwWindow) SizeLimits() (minw, minh, maxw, maxh int) {
	return -1, -1, -1, -1
}

func (w *glfwWindow) SetSizeLimits(minw, minh, maxw, maxh int) {
	C.glfwSetWindowSizeLimits(
		w.handle,
		C.int(minw), C.int(minh),
		C.int(maxw), C.int(maxh),
	)
}

func (w *glfwWindow) Position() (x, y int) {
	var xpos, ypos C.int
	C.glfwGetWindowPos(w.handle, &xpos, &ypos)
	return int(xpos), int(ypos)
}

func (w *glfwWindow) SetPosition(x, y int) {
	C.glfwSetWindowPos(w.handle, C.int(x), C.int(y))
}

func (w *glfwWindow) IsFloating() bool {
	return C.glfwGetWindowAttrib(w.handle, C.GLFW_FLOATING) == C.GLFW_TRUE
}

func (w *glfwWindow) IsMaximized() bool {
	return C.glfwGetWindowAttrib(w.handle, C.GLFW_MAXIMIZED) == C.GLFW_TRUE
}

func (w *glfwWindow) Maximize() {
	C.glfwMaximizeWindow(w.handle)
}

func (w *glfwWindow) IsMinimized() bool {
	return C.glfwGetWindowAttrib(w.handle, C.GLFW_ICONIFIED) == C.GLFW_TRUE
}

func (w *glfwWindow) Minimize() {
	C.glfwIconifyWindow(w.handle)
}

func (w *glfwWindow) ShouldClose() bool {
	return C.glfwWindowShouldClose(w.handle) == C.GLFW_TRUE
}

// PollEvents processes OS window events (must be called regularly).
func PollEvents() {
	C.glfwPollEvents()
}

// DestroyWindow destroys a GLFW window and releases its native resources.
func DestroyWindow(w any) {
	gw, ok := w.(*glfwWindow)
	if !ok || gw == nil || gw.handle == nil {
		return
	}
	C.glfwDestroyWindow(gw.handle)
	gw.handle = nil
}

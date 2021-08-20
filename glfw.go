package glfw

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"unsafe"
)

// changes are accumulated in this file

func Init() int {
	err := glfw.Init()
	if err != nil {
		return 0
	}
	return 1
}

func CreateWindow(width int, height int, title string, monitor *Monitor, share *Window) *Window {
	w, _ := glfw.CreateWindow(width, height, title, (*glfw.Monitor)(monitor), (*glfw.Window)(share))
	return (*Window)(w)
}

type ErrorCallback func(int string)

var currentErrCb ErrorCallback

func SetErrorCallback(cb ErrorCallback) (prev ErrorCallback) {
	prev = currentErrCb
	currentErrCb = cb
	return prev
}

func (w *Window) GetFramebufferSize(x, y *int) {
	*x, *y = glfw.Window.GetFramebufferSize(w)
}

func (m *Monitor) GetPos(x, y *int) {
	*x, *y = glfw.Monitor.GetPos((*glfw.Monitor)(m))
}

func (m *Monitor) GetPhysicalSize(x, y *int) {
	*x, *y = glfw.Monitor.GetPhysicalSize((*glfw.Monitor)(m))
}

func (j Joystick) GetAxes(c *int) []float32 {
	a := glfw.Joystick.GetAxes(j)
	*c = len(a)
	return a
}

func (j Joystick) GetButtons(c *int) []Action {
	a := glfw.Joystick.GetButtons(j)
	*c = len(a)
	return []Action(unsafe.Slice((*Action)(&a[0]), len(a))) // avoids creating another slice
}

func (j Joystick) GetHats(c *int) []JoystickHatState {
	a := glfw.Joystick.GetHats(j)
	*c = len(a)
	return []JoystickHatState(unsafe.Slice((*JoystickHatState)(&a[0]), len(a))) // avoids creating another slice
}

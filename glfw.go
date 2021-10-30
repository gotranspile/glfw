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
	*x, *y = (*glfw.Window)(w).GetFramebufferSize()
}

func (w *Window) GetPos(x, y *int) {
	*x, *y = (*glfw.Window)(w).GetPos()
}

func (w *Window) GetSize(x, y *int) {
	*x, *y = (*glfw.Window)(w).GetSize()
}

func (m *Monitor) GetPos(x, y *int) {
	*x, *y = (*glfw.Monitor)(m).GetPos()
}

func (m *Monitor) GetPhysicalSize(x, y *int) {
	*x, *y = (*glfw.Monitor)(m).GetPhysicalSize()
}

func GetMonitors(count *int) **glfw.Monitor {
	m := glfw.GetMonitors()
	*count = len(m)
	return &m[0]
}

func (j Joystick) GetAxes(c *int) []float32 {
	a := glfw.Joystick(j).GetAxes()
	*c = len(a)
	return a
}

func (j Joystick) GetButtons(c *int) []Action {
	a := glfw.Joystick(j).GetButtons()
	*c = len(a)
	return (*[0xFFFFFFFF]Action)(unsafe.Pointer(&a[0]))[:len(a):len(a)]// avoids creating another slice
}

func (j Joystick) GetHats(c *int) []JoystickHatState {
	a := glfw.Joystick(j).GetHats()
	*c = len(a)
	return (*[0xFFFFFFFF]JoystickHatState)(unsafe.Pointer(&a[0]))[:len(a):len(a)] // avoids creating another slice
}

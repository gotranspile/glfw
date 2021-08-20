package main

import (
	"fmt"
	"strings"
)

func main() {
	var b strings.Builder
	b.WriteString("type (\n")
	for _, line := range strings.Split(s, "\n") {
		if len(line) <= 0 {
			continue
		}
		if !strings.HasPrefix(line, "type") {
			continue
		}
		f := strings.Fields(line)
		b.WriteString(fmt.Sprintf("\t%s glfw.%s\n", f[1], f[1]))
	}
	b.WriteString(")\n")
	fmt.Println(b.String())
}

const s = `func DefaultWindowHints()
func DetachCurrentContext()
func ExtensionSupported(extension string) bool
func GetClipboardString() string
func GetKeyName(key Key, scancode int) string
func GetKeyScancode(key Key) int
func GetProcAddress(procname string) unsafe.Pointer
func GetTime() float64
func GetTimerFrequency() uint64
func GetTimerValue() uint64
func GetVersion() (major, minor, revision int)
func GetVersionString() string
func GetVulkanGetInstanceProcAddress() unsafe.Pointer
func GetX11Display() *C.Display
func GetX11SelectionString() string
func Init() error
func InitHint(hint Hint, value int)
func PollEvents()
func PostEmptyEvent()
func RawMouseMotionSupported() bool
func SetClipboardString(str string)
func SetTime(time float64)
func SetX11SelectionString(str string)
func SwapInterval(interval int)
func Terminate()
func UpdateGamepadMappings(mapping string) bool
func VulkanSupported() bool
func WaitEvents()
func WaitEventsTimeout(timeout float64)
func WindowHint(target Hint, hint int)
func WindowHintString(hint Hint, value string)
type Action
type CharCallback
type CharModsCallback
type CloseCallback
type ContentScaleCallback
type Cursor
func CreateCursor(img image.Image, xhot, yhot int) *Cursor
func CreateStandardCursor(shape StandardCursor) *Cursor
func (c *Cursor) Destroy()
type CursorEnterCallback
type CursorPosCallback
type DropCallback
type Error
func (e *Error) Error() string
type ErrorCode
func (e ErrorCode) String() string
type FocusCallback
type FramebufferSizeCallback
type GamepadAxis
type GamepadButton
type GamepadState
type GammaRamp
type Hint
type IconifyCallback
type InputMode
type Joystick
func (joy Joystick) GetAxes() []float32
func (joy Joystick) GetButtons() []Action
func (joy Joystick) GetGUID() string
func (joy Joystick) GetGamepadName() string
func (joy Joystick) GetGamepadState() *GamepadState
func (joy Joystick) GetHats() []JoystickHatState
func (joy Joystick) GetName() string
func (joy Joystick) GetUserPointer() unsafe.Pointer
func (joy Joystick) IsGamepad() bool
func (joy Joystick) Present() bool
func (joy Joystick) SetUserPointer(pointer unsafe.Pointer)
type JoystickCallback
func SetJoystickCallback(cbfun JoystickCallback) (previous JoystickCallback)
type JoystickHatState
type Key
type KeyCallback
type MaximizeCallback
type ModifierKey
type Monitor
func GetMonitors() []*Monitor
func GetPrimaryMonitor() *Monitor
func (m *Monitor) GetContentScale() (float32, float32)
func (m *Monitor) GetGammaRamp() *GammaRamp
func (m *Monitor) GetName() string
func (m *Monitor) GetPhysicalSize() (width, height int)
func (m *Monitor) GetPos() (x, y int)
func (m *Monitor) GetUserPointer() unsafe.Pointer
func (m *Monitor) GetVideoMode() *VidMode
func (m *Monitor) GetVideoModes() []*VidMode
func (m *Monitor) GetWorkarea() (x, y, width, height int)
func (m *Monitor) GetX11Adapter() C.RRCrtc
func (m *Monitor) GetX11Monitor() C.RROutput
func (m *Monitor) SetGamma(gamma float32)
func (m *Monitor) SetGammaRamp(ramp *GammaRamp)
func (m *Monitor) SetUserPointer(pointer unsafe.Pointer)
type MonitorCallback
func SetMonitorCallback(cbfun MonitorCallback) MonitorCallback
type MouseButton
type MouseButtonCallback
type PeripheralEvent
type PosCallback
type RefreshCallback
type ScrollCallback
type SizeCallback
type StandardCursor
type VidMode
type Window
func CreateWindow(width, height int, title string, monitor *Monitor, share *Window) (*Window, error)
func GetCurrentContext() *Window
func GoWindow(window unsafe.Pointer) *Window
func (window *Window) CreateWindowSurface(instance interface{}, allocCallbacks unsafe.Pointer) (surface uintptr, err error)
func (w *Window) Destroy()
func (w *Window) Focus()
func (w *Window) GetAttrib(attrib Hint) int
func (w *Window) GetClipboardString() string
func (w *Window) GetContentScale() (float32, float32)
func (w *Window) GetCursorPos() (x, y float64)
func (w *Window) GetFrameSize() (left, top, right, bottom int)
func (w *Window) GetFramebufferSize() (width, height int)
func (w *Window) GetGLXContext() C.GLXContext
func (w *Window) GetGLXWindow() C.GLXWindow
func (w *Window) GetInputMode(mode InputMode) int
func (w *Window) GetKey(key Key) Action
func (w *Window) GetMonitor() *Monitor
func (w *Window) GetMouseButton(button MouseButton) Action
func (w *Window) GetOpacity() float32
func (w *Window) GetPos() (x, y int)
func (window *Window) GetRequiredInstanceExtensions() []string
func (w *Window) GetSize() (width, height int)
func (w *Window) GetUserPointer() unsafe.Pointer
func (w *Window) GetX11Window() C.Window
func (w *Window) Handle() unsafe.Pointer
func (w *Window) Hide()
func (w *Window) Iconify()
func (w *Window) MakeContextCurrent()
func (w *Window) Maximize()
func (w *Window) RequestAttention()
func (w *Window) Restore()
func (w *Window) SetAspectRatio(numer, denom int)
func (w *Window) SetAttrib(attrib Hint, value int)
func (w *Window) SetCharCallback(cbfun CharCallback) (previous CharCallback)
func (w *Window) SetCharModsCallback(cbfun CharModsCallback) (previous CharModsCallback)DEPRECATED
func (w *Window) SetClipboardString(str string)
func (w *Window) SetCloseCallback(cbfun CloseCallback) (previous CloseCallback)
func (w *Window) SetContentScaleCallback(cbfun ContentScaleCallback) ContentScaleCallback
func (w *Window) SetCursor(c *Cursor)
func (w *Window) SetCursorEnterCallback(cbfun CursorEnterCallback) (previous CursorEnterCallback)
func (w *Window) SetCursorPos(xpos, ypos float64)
func (w *Window) SetCursorPosCallback(cbfun CursorPosCallback) (previous CursorPosCallback)
func (w *Window) SetDropCallback(cbfun DropCallback) (previous DropCallback)
func (w *Window) SetFocusCallback(cbfun FocusCallback) (previous FocusCallback)
func (w *Window) SetFramebufferSizeCallback(cbfun FramebufferSizeCallback) (previous FramebufferSizeCallback)
func (w *Window) SetIcon(images []image.Image)
func (w *Window) SetIconifyCallback(cbfun IconifyCallback) (previous IconifyCallback)
func (w *Window) SetInputMode(mode InputMode, value int)
func (w *Window) SetKeyCallback(cbfun KeyCallback) (previous KeyCallback)
func (w *Window) SetMaximizeCallback(cbfun MaximizeCallback) MaximizeCallback
func (w *Window) SetMonitor(monitor *Monitor, xpos, ypos, width, height, refreshRate int)
func (w *Window) SetMouseButtonCallback(cbfun MouseButtonCallback) (previous MouseButtonCallback)
func (w *Window) SetOpacity(opacity float32)
func (w *Window) SetPos(xpos, ypos int)
func (w *Window) SetPosCallback(cbfun PosCallback) (previous PosCallback)
func (w *Window) SetRefreshCallback(cbfun RefreshCallback) (previous RefreshCallback)
func (w *Window) SetScrollCallback(cbfun ScrollCallback) (previous ScrollCallback)
func (w *Window) SetShouldClose(value bool)
func (w *Window) SetSize(width, height int)
func (w *Window) SetSizeCallback(cbfun SizeCallback) (previous SizeCallback)
func (w *Window) SetSizeLimits(minw, minh, maxw, maxh int)
func (w *Window) SetTitle(title string)
func (w *Window) SetUserPointer(pointer unsafe.Pointer)
func (w *Window) ShouldClose() bool
func (w *Window) Show()
func (w *Window) SwapBuffers()`

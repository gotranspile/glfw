package main

import (
	"fmt"
	"strings"
)

func main() {
	var b strings.Builder
	b.WriteString("const (\n")
	for _, line := range strings.Split(s, "\n") {
		if len(line) <= 0 {
			continue
		}
		if strings.HasPrefix(line, "const") {
			continue
		}
		if strings.HasPrefix(line, ")") {
			continue
		}
		f := strings.Fields(line)
		b.WriteString(fmt.Sprintf("\t%s = glfw.%s\n", f[0], f[0]))
	}
	b.WriteString(")\n")
	fmt.Println(b.String())
}

const s = `const (
	VersionMajor    = C.GLFW_VERSION_MAJOR    // This is incremented when the API is changed in non-compatible ways.
	VersionMinor    = C.GLFW_VERSION_MINOR    // This is incremented when features are added to the API but it remains backward-compatible.
	VersionRevision = C.GLFW_VERSION_REVISION // This is incremented when a bug fix release is made that does not contain any API changes.
)

const (
	CursorNormal   int = C.GLFW_CURSOR_NORMAL
	CursorHidden   int = C.GLFW_CURSOR_HIDDEN
	CursorDisabled int = C.GLFW_CURSOR_DISABLED
)

const (
	OpenGLAPI   int = C.GLFW_OPENGL_API
	OpenGLESAPI int = C.GLFW_OPENGL_ES_API
	NoAPI       int = C.GLFW_NO_API
)

const (
	NativeContextAPI int = C.GLFW_NATIVE_CONTEXT_API
	EGLContextAPI    int = C.GLFW_EGL_CONTEXT_API
	OSMesaContextAPI int = C.GLFW_OSMESA_CONTEXT_API
)

const (
	NoRobustness        int = C.GLFW_NO_ROBUSTNESS
	NoResetNotification int = C.GLFW_NO_RESET_NOTIFICATION
	LoseContextOnReset  int = C.GLFW_LOSE_CONTEXT_ON_RESET
)

const (
	AnyReleaseBehavior   int = C.GLFW_ANY_RELEASE_BEHAVIOR
	ReleaseBehaviorFlush int = C.GLFW_RELEASE_BEHAVIOR_FLUSH
	ReleaseBehaviorNone  int = C.GLFW_RELEASE_BEHAVIOR_NONE
)

const (
	OpenGLAnyProfile    int = C.GLFW_OPENGL_ANY_PROFILE
	OpenGLCoreProfile   int = C.GLFW_OPENGL_CORE_PROFILE
	OpenGLCompatProfile int = C.GLFW_OPENGL_COMPAT_PROFILE
)

const (
	True     int = 1 // GL_TRUE
	False    int = 0 // GL_FALSE
	DontCare int = C.GLFW_DONT_CARE
)

const (
	HatCentered  JoystickHatState = C.GLFW_HAT_CENTERED
	HatUp        JoystickHatState = C.GLFW_HAT_UP
	HatRight     JoystickHatState = C.GLFW_HAT_RIGHT
	HatDown      JoystickHatState = C.GLFW_HAT_DOWN
	HatLeft      JoystickHatState = C.GLFW_HAT_LEFT
	HatRightUp   JoystickHatState = C.GLFW_HAT_RIGHT_UP
	HatRightDown JoystickHatState = C.GLFW_HAT_RIGHT_DOWN
	HatLeftUp    JoystickHatState = C.GLFW_HAT_LEFT_UP
	HatLeftDown  JoystickHatState = C.GLFW_HAT_LEFT_DOWN
)

const (
	KeyUnknown      Key = C.GLFW_KEY_UNKNOWN
	KeySpace        Key = C.GLFW_KEY_SPACE
	KeyApostrophe   Key = C.GLFW_KEY_APOSTROPHE
	KeyComma        Key = C.GLFW_KEY_COMMA
	KeyMinus        Key = C.GLFW_KEY_MINUS
	KeyPeriod       Key = C.GLFW_KEY_PERIOD
	KeySlash        Key = C.GLFW_KEY_SLASH
	Key0            Key = C.GLFW_KEY_0
	Key1            Key = C.GLFW_KEY_1
	Key2            Key = C.GLFW_KEY_2
	Key3            Key = C.GLFW_KEY_3
	Key4            Key = C.GLFW_KEY_4
	Key5            Key = C.GLFW_KEY_5
	Key6            Key = C.GLFW_KEY_6
	Key7            Key = C.GLFW_KEY_7
	Key8            Key = C.GLFW_KEY_8
	Key9            Key = C.GLFW_KEY_9
	KeySemicolon    Key = C.GLFW_KEY_SEMICOLON
	KeyEqual        Key = C.GLFW_KEY_EQUAL
	KeyA            Key = C.GLFW_KEY_A
	KeyB            Key = C.GLFW_KEY_B
	KeyC            Key = C.GLFW_KEY_C
	KeyD            Key = C.GLFW_KEY_D
	KeyE            Key = C.GLFW_KEY_E
	KeyF            Key = C.GLFW_KEY_F
	KeyG            Key = C.GLFW_KEY_G
	KeyH            Key = C.GLFW_KEY_H
	KeyI            Key = C.GLFW_KEY_I
	KeyJ            Key = C.GLFW_KEY_J
	KeyK            Key = C.GLFW_KEY_K
	KeyL            Key = C.GLFW_KEY_L
	KeyM            Key = C.GLFW_KEY_M
	KeyN            Key = C.GLFW_KEY_N
	KeyO            Key = C.GLFW_KEY_O
	KeyP            Key = C.GLFW_KEY_P
	KeyQ            Key = C.GLFW_KEY_Q
	KeyR            Key = C.GLFW_KEY_R
	KeyS            Key = C.GLFW_KEY_S
	KeyT            Key = C.GLFW_KEY_T
	KeyU            Key = C.GLFW_KEY_U
	KeyV            Key = C.GLFW_KEY_V
	KeyW            Key = C.GLFW_KEY_W
	KeyX            Key = C.GLFW_KEY_X
	KeyY            Key = C.GLFW_KEY_Y
	KeyZ            Key = C.GLFW_KEY_Z
	KeyLeftBracket  Key = C.GLFW_KEY_LEFT_BRACKET
	KeyBackslash    Key = C.GLFW_KEY_BACKSLASH
	KeyRightBracket Key = C.GLFW_KEY_RIGHT_BRACKET
	KeyGraveAccent  Key = C.GLFW_KEY_GRAVE_ACCENT
	KeyWorld1       Key = C.GLFW_KEY_WORLD_1
	KeyWorld2       Key = C.GLFW_KEY_WORLD_2
	KeyEscape       Key = C.GLFW_KEY_ESCAPE
	KeyEnter        Key = C.GLFW_KEY_ENTER
	KeyTab          Key = C.GLFW_KEY_TAB
	KeyBackspace    Key = C.GLFW_KEY_BACKSPACE
	KeyInsert       Key = C.GLFW_KEY_INSERT
	KeyDelete       Key = C.GLFW_KEY_DELETE
	KeyRight        Key = C.GLFW_KEY_RIGHT
	KeyLeft         Key = C.GLFW_KEY_LEFT
	KeyDown         Key = C.GLFW_KEY_DOWN
	KeyUp           Key = C.GLFW_KEY_UP
	KeyPageUp       Key = C.GLFW_KEY_PAGE_UP
	KeyPageDown     Key = C.GLFW_KEY_PAGE_DOWN
	KeyHome         Key = C.GLFW_KEY_HOME
	KeyEnd          Key = C.GLFW_KEY_END
	KeyCapsLock     Key = C.GLFW_KEY_CAPS_LOCK
	KeyScrollLock   Key = C.GLFW_KEY_SCROLL_LOCK
	KeyNumLock      Key = C.GLFW_KEY_NUM_LOCK
	KeyPrintScreen  Key = C.GLFW_KEY_PRINT_SCREEN
	KeyPause        Key = C.GLFW_KEY_PAUSE
	KeyF1           Key = C.GLFW_KEY_F1
	KeyF2           Key = C.GLFW_KEY_F2
	KeyF3           Key = C.GLFW_KEY_F3
	KeyF4           Key = C.GLFW_KEY_F4
	KeyF5           Key = C.GLFW_KEY_F5
	KeyF6           Key = C.GLFW_KEY_F6
	KeyF7           Key = C.GLFW_KEY_F7
	KeyF8           Key = C.GLFW_KEY_F8
	KeyF9           Key = C.GLFW_KEY_F9
	KeyF10          Key = C.GLFW_KEY_F10
	KeyF11          Key = C.GLFW_KEY_F11
	KeyF12          Key = C.GLFW_KEY_F12
	KeyF13          Key = C.GLFW_KEY_F13
	KeyF14          Key = C.GLFW_KEY_F14
	KeyF15          Key = C.GLFW_KEY_F15
	KeyF16          Key = C.GLFW_KEY_F16
	KeyF17          Key = C.GLFW_KEY_F17
	KeyF18          Key = C.GLFW_KEY_F18
	KeyF19          Key = C.GLFW_KEY_F19
	KeyF20          Key = C.GLFW_KEY_F20
	KeyF21          Key = C.GLFW_KEY_F21
	KeyF22          Key = C.GLFW_KEY_F22
	KeyF23          Key = C.GLFW_KEY_F23
	KeyF24          Key = C.GLFW_KEY_F24
	KeyF25          Key = C.GLFW_KEY_F25
	KeyKP0          Key = C.GLFW_KEY_KP_0
	KeyKP1          Key = C.GLFW_KEY_KP_1
	KeyKP2          Key = C.GLFW_KEY_KP_2
	KeyKP3          Key = C.GLFW_KEY_KP_3
	KeyKP4          Key = C.GLFW_KEY_KP_4
	KeyKP5          Key = C.GLFW_KEY_KP_5
	KeyKP6          Key = C.GLFW_KEY_KP_6
	KeyKP7          Key = C.GLFW_KEY_KP_7
	KeyKP8          Key = C.GLFW_KEY_KP_8
	KeyKP9          Key = C.GLFW_KEY_KP_9
	KeyKPDecimal    Key = C.GLFW_KEY_KP_DECIMAL
	KeyKPDivide     Key = C.GLFW_KEY_KP_DIVIDE
	KeyKPMultiply   Key = C.GLFW_KEY_KP_MULTIPLY
	KeyKPSubtract   Key = C.GLFW_KEY_KP_SUBTRACT
	KeyKPAdd        Key = C.GLFW_KEY_KP_ADD
	KeyKPEnter      Key = C.GLFW_KEY_KP_ENTER
	KeyKPEqual      Key = C.GLFW_KEY_KP_EQUAL
	KeyLeftShift    Key = C.GLFW_KEY_LEFT_SHIFT
	KeyLeftControl  Key = C.GLFW_KEY_LEFT_CONTROL
	KeyLeftAlt      Key = C.GLFW_KEY_LEFT_ALT
	KeyLeftSuper    Key = C.GLFW_KEY_LEFT_SUPER
	KeyRightShift   Key = C.GLFW_KEY_RIGHT_SHIFT
	KeyRightControl Key = C.GLFW_KEY_RIGHT_CONTROL
	KeyRightAlt     Key = C.GLFW_KEY_RIGHT_ALT
	KeyRightSuper   Key = C.GLFW_KEY_RIGHT_SUPER
	KeyMenu         Key = C.GLFW_KEY_MENU
	KeyLast         Key = C.GLFW_KEY_LAST
)

const (
	ModShift    ModifierKey = C.GLFW_MOD_SHIFT
	ModControl  ModifierKey = C.GLFW_MOD_CONTROL
	ModAlt      ModifierKey = C.GLFW_MOD_ALT
	ModSuper    ModifierKey = C.GLFW_MOD_SUPER
	ModCapsLock ModifierKey = C.GLFW_MOD_CAPS_LOCK
	ModNumLock  ModifierKey = C.GLFW_MOD_NUM_LOCK
)

const (
	MouseButton1      MouseButton = C.GLFW_MOUSE_BUTTON_1
	MouseButton2      MouseButton = C.GLFW_MOUSE_BUTTON_2
	MouseButton3      MouseButton = C.GLFW_MOUSE_BUTTON_3
	MouseButton4      MouseButton = C.GLFW_MOUSE_BUTTON_4
	MouseButton5      MouseButton = C.GLFW_MOUSE_BUTTON_5
	MouseButton6      MouseButton = C.GLFW_MOUSE_BUTTON_6
	MouseButton7      MouseButton = C.GLFW_MOUSE_BUTTON_7
	MouseButton8      MouseButton = C.GLFW_MOUSE_BUTTON_8
	MouseButtonLast   MouseButton = C.GLFW_MOUSE_BUTTON_LAST
	MouseButtonLeft   MouseButton = C.GLFW_MOUSE_BUTTON_LEFT
	MouseButtonRight  MouseButton = C.GLFW_MOUSE_BUTTON_RIGHT
	MouseButtonMiddle MouseButton = C.GLFW_MOUSE_BUTTON_MIDDLE
)

const (
	Connected    PeripheralEvent = C.GLFW_CONNECTED
	Disconnected PeripheralEvent = C.GLFW_DISCONNECTED
)

const (
	ArrowCursor     StandardCursor = C.GLFW_ARROW_CURSOR
	IBeamCursor     StandardCursor = C.GLFW_IBEAM_CURSOR
	CrosshairCursor StandardCursor = C.GLFW_CROSSHAIR_CURSOR
	HandCursor      StandardCursor = C.GLFW_HAND_CURSOR
	HResizeCursor   StandardCursor = C.GLFW_HRESIZE_CURSOR
	VResizeCursor   StandardCursor = C.GLFW_VRESIZE_CURSOR
)

const (
	Joystick1    Joystick = C.GLFW_JOYSTICK_1
	Joystick2    Joystick = C.GLFW_JOYSTICK_2
	Joystick3    Joystick = C.GLFW_JOYSTICK_3
	Joystick4    Joystick = C.GLFW_JOYSTICK_4
	Joystick5    Joystick = C.GLFW_JOYSTICK_5
	Joystick6    Joystick = C.GLFW_JOYSTICK_6
	Joystick7    Joystick = C.GLFW_JOYSTICK_7
	Joystick8    Joystick = C.GLFW_JOYSTICK_8
	Joystick9    Joystick = C.GLFW_JOYSTICK_9
	Joystick10   Joystick = C.GLFW_JOYSTICK_10
	Joystick11   Joystick = C.GLFW_JOYSTICK_11
	Joystick12   Joystick = C.GLFW_JOYSTICK_12
	Joystick13   Joystick = C.GLFW_JOYSTICK_13
	Joystick14   Joystick = C.GLFW_JOYSTICK_14
	Joystick15   Joystick = C.GLFW_JOYSTICK_15
	Joystick16   Joystick = C.GLFW_JOYSTICK_16
	JoystickLast Joystick = C.GLFW_JOYSTICK_LAST
)

const (
	CursorMode             InputMode = C.GLFW_CURSOR               // See Cursor mode values
	StickyKeysMode         InputMode = C.GLFW_STICKY_KEYS          // Value can be either 1 or 0
	StickyMouseButtonsMode InputMode = C.GLFW_STICKY_MOUSE_BUTTONS // Value can be either 1 or 0
	LockKeyMods            InputMode = C.GLFW_LOCK_KEY_MODS        // Value can be either 1 or 0
	RawMouseMotion         InputMode = C.GLFW_RAW_MOUSE_MOTION     // Value can be either 1 or 0
)

const (
	CocoaFrameNAME  Hint = C.GLFW_COCOA_FRAME_NAME  // Specifies the UTF-8 encoded name to use for autosaving the window frame, or if empty disables frame autosaving for the window.
	X11ClassName    Hint = C.GLFW_X11_CLASS_NAME    // Specifies the desired ASCII encoded class parts of the ICCCM WM_CLASS window property.nd instance parts of the ICCCM WM_CLASS window property.
	X11InstanceName Hint = C.GLFW_X11_INSTANCE_NAME // Specifies the desired ASCII encoded instance parts of the ICCCM WM_CLASS window property.nd instance parts of the ICCCM WM_CLASS window property.
)

const (
	ContextRevision        Hint = C.GLFW_CONTEXT_REVISION
	RedBits                Hint = C.GLFW_RED_BITS                 // Specifies the desired bit depth of the default framebuffer.
	GreenBits              Hint = C.GLFW_GREEN_BITS               // Specifies the desired bit depth of the default framebuffer.
	BlueBits               Hint = C.GLFW_BLUE_BITS                // Specifies the desired bit depth of the default framebuffer.
	AlphaBits              Hint = C.GLFW_ALPHA_BITS               // Specifies the desired bit depth of the default framebuffer.
	DepthBits              Hint = C.GLFW_DEPTH_BITS               // Specifies the desired bit depth of the default framebuffer.
	StencilBits            Hint = C.GLFW_STENCIL_BITS             // Specifies the desired bit depth of the default framebuffer.
	AccumRedBits           Hint = C.GLFW_ACCUM_RED_BITS           // Specifies the desired bit depth of the accumulation buffer.
	AccumGreenBits         Hint = C.GLFW_ACCUM_GREEN_BITS         // Specifies the desired bit depth of the accumulation buffer.
	AccumBlueBits          Hint = C.GLFW_ACCUM_BLUE_BITS          // Specifies the desired bit depth of the accumulation buffer.
	AccumAlphaBits         Hint = C.GLFW_ACCUM_ALPHA_BITS         // Specifies the desired bit depth of the accumulation buffer.
	AuxBuffers             Hint = C.GLFW_AUX_BUFFERS              // Specifies the desired number of auxiliary buffers.
	Stereo                 Hint = C.GLFW_STEREO                   // Specifies whether to use stereoscopic rendering. Hard constraint.
	Samples                Hint = C.GLFW_SAMPLES                  // Specifies the desired number of samples to use for multisampling. Zero disables multisampling.
	SRGBCapable            Hint = C.GLFW_SRGB_CAPABLE             // Specifies whether the framebuffer should be sRGB capable.
	RefreshRate            Hint = C.GLFW_REFRESH_RATE             // Specifies the desired refresh rate for full screen windows. If set to zero, the highest available refresh rate will be used. This hint is ignored for windowed mode windows.
	DoubleBuffer           Hint = C.GLFW_DOUBLEBUFFER             // Specifies whether the framebuffer should be double buffered. You nearly always want to use double buffering. This is a hard constraint.
	CocoaGraphicsSwitching Hint = C.GLFW_COCOA_GRAPHICS_SWITCHING // Specifies whether to in Automatic Graphics Switching, i.e. to allow the system to choose the integrated GPU for the OpenGL context and move it between GPUs if necessary or whether to force it to always run on the discrete GPU.
	CocoaRetinaFramebuffer Hint = C.GLFW_COCOA_RETINA_FRAMEBUFFER // Specifies whether to use full resolution framebuffers on Retina displays.
)

const (
	ClientAPI               Hint = C.GLFW_CLIENT_API               // Specifies which client API to create the context for. Hard constraint.
	ContextVersionMajor     Hint = C.GLFW_CONTEXT_VERSION_MAJOR    // Specifies the client API version that the created context must be compatible with.
	ContextVersionMinor     Hint = C.GLFW_CONTEXT_VERSION_MINOR    // Specifies the client API version that the created context must be compatible with.
	ContextRobustness       Hint = C.GLFW_CONTEXT_ROBUSTNESS       // Specifies the robustness strategy to be used by the context.
	ContextReleaseBehavior  Hint = C.GLFW_CONTEXT_RELEASE_BEHAVIOR // Specifies the release behavior to be used by the context.
	OpenGLForwardCompatible Hint = C.GLFW_OPENGL_FORWARD_COMPAT    // Specifies whether the OpenGL context should be forward-compatible. Hard constraint.
	OpenGLDebugContext      Hint = C.GLFW_OPENGL_DEBUG_CONTEXT     // Specifies whether to create a debug OpenGL context, which may have additional error and performance issue reporting functionality. If OpenGL ES is requested, this hint is ignored.
	OpenGLProfile           Hint = C.GLFW_OPENGL_PROFILE           // Specifies which OpenGL profile to create the context for. Hard constraint.
	ContextCreationAPI      Hint = C.GLFW_CONTEXT_CREATION_API     // Specifies which context creation API to use to create the context.
)

const (
	Focused                Hint = C.GLFW_FOCUSED                 // Specifies whether the window will be given input focus when created. This hint is ignored for full screen and initially hidden windows.
	Iconified              Hint = C.GLFW_ICONIFIED               // Specifies whether the window will be minimized.
	Maximized              Hint = C.GLFW_MAXIMIZED               // Specifies whether the window is maximized.
	Visible                Hint = C.GLFW_VISIBLE                 // Specifies whether the window will be initially visible.
	Hovered                Hint = C.GLFW_HOVERED                 // Specifies whether the cursor is currently directly over the content area of the window, with no other windows between. See Cursor enter/leave events for details.
	Resizable              Hint = C.GLFW_RESIZABLE               // Specifies whether the window will be resizable by the user.
	Decorated              Hint = C.GLFW_DECORATED               // Specifies whether the window will have window decorations such as a border, a close widget, etc.
	Floating               Hint = C.GLFW_FLOATING                // Specifies whether the window will be always-on-top.
	AutoIconify            Hint = C.GLFW_AUTO_ICONIFY            // Specifies whether fullscreen windows automatically iconify (and restore the previous video mode) on focus loss.
	CenterCursor           Hint = C.GLFW_CENTER_CURSOR           // Specifies whether the cursor should be centered over newly created full screen windows. This hint is ignored for windowed mode windows.
	TransparentFramebuffer Hint = C.GLFW_TRANSPARENT_FRAMEBUFFER // Specifies whether the framebuffer should be transparent.
	FocusOnShow            Hint = C.GLFW_FOCUS_ON_SHOW           // Specifies whether the window will be given input focus when glfwShowWindow is called.
	ScaleToMonitor         Hint = C.GLFW_SCALE_TO_MONITOR        // Specified whether the window content area should be resized based on the monitor content scale of any monitor it is placed on. This includes the initial placement when the window is created.
)

const (
	JoystickHatButtons  Hint = C.GLFW_JOYSTICK_HAT_BUTTONS  // Specifies whether to also expose joystick hats as buttons, for compatibility with earlier versions of GLFW that did not have glfwGetJoystickHats.
	CocoaChdirResources Hint = C.GLFW_COCOA_CHDIR_RESOURCES // Specifies whether to set the current directory to the application to the Contents/Resources subdirectory of the application's bundle, if present.
	CocoaMenubar        Hint = C.GLFW_COCOA_MENUBAR         // Specifies whether to create a basic menu bar, either from a nib or manually, when the first window is created, which is when AppKit is initialized.
)

const (
	ButtonA           GamepadButton = C.GLFW_GAMEPAD_BUTTON_A
	ButtonB           GamepadButton = C.GLFW_GAMEPAD_BUTTON_B
	ButtonX           GamepadButton = C.GLFW_GAMEPAD_BUTTON_X
	ButtonY           GamepadButton = C.GLFW_GAMEPAD_BUTTON_Y
	ButtonLeftBumper  GamepadButton = C.GLFW_GAMEPAD_BUTTON_LEFT_BUMPER
	ButtonRightBumper GamepadButton = C.GLFW_GAMEPAD_BUTTON_RIGHT_BUMPER
	ButtonBack        GamepadButton = C.GLFW_GAMEPAD_BUTTON_BACK
	ButtonStart       GamepadButton = C.GLFW_GAMEPAD_BUTTON_START
	ButtonGuide       GamepadButton = C.GLFW_GAMEPAD_BUTTON_GUIDE
	ButtonLeftThumb   GamepadButton = C.GLFW_GAMEPAD_BUTTON_LEFT_THUMB
	ButtonRightThumb  GamepadButton = C.GLFW_GAMEPAD_BUTTON_RIGHT_THUMB
	ButtonDpadUp      GamepadButton = C.GLFW_GAMEPAD_BUTTON_DPAD_UP
	ButtonDpadRight   GamepadButton = C.GLFW_GAMEPAD_BUTTON_DPAD_RIGHT
	ButtonDpadDown    GamepadButton = C.GLFW_GAMEPAD_BUTTON_DPAD_DOWN
	ButtonDpadLeft    GamepadButton = C.GLFW_GAMEPAD_BUTTON_DPAD_LEFT
	ButtonLast        GamepadButton = C.GLFW_GAMEPAD_BUTTON_LAST
	ButtonCross       GamepadButton = C.GLFW_GAMEPAD_BUTTON_CROSS
	ButtonCircle      GamepadButton = C.GLFW_GAMEPAD_BUTTON_CIRCLE
	ButtonSquare      GamepadButton = C.GLFW_GAMEPAD_BUTTON_SQUARE
	ButtonTriangle    GamepadButton = C.GLFW_GAMEPAD_BUTTON_TRIANGLE
)

const (
	AxisLeftX        GamepadAxis = C.GLFW_GAMEPAD_AXIS_LEFT_X
	AxisLeftY        GamepadAxis = C.GLFW_GAMEPAD_AXIS_LEFT_Y
	AxisRightX       GamepadAxis = C.GLFW_GAMEPAD_AXIS_RIGHT_X
	AxisRightY       GamepadAxis = C.GLFW_GAMEPAD_AXIS_RIGHT_Y
	AxisLeftTrigger  GamepadAxis = C.GLFW_GAMEPAD_AXIS_LEFT_TRIGGER
	AxisRightTrigger GamepadAxis = C.GLFW_GAMEPAD_AXIS_RIGHT_TRIGGER
	AxisLast         GamepadAxis = C.GLFW_GAMEPAD_AXIS_LAST
)

const (
	Release Action = C.GLFW_RELEASE // The key or button was released.
	Press   Action = C.GLFW_PRESS   // The key or button was pressed.
	Repeat  Action = C.GLFW_REPEAT  // The key was held down until it repeated.
)
`
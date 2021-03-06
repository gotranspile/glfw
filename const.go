package glfw

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	VersionMajor = glfw.VersionMajor
	VersionMinor = glfw.VersionMinor
	VersionRevision = glfw.VersionRevision
	CursorNormal = glfw.CursorNormal
	CursorHidden = glfw.CursorHidden
	CursorDisabled = glfw.CursorDisabled
	OpenGLAPI = glfw.OpenGLAPI
	OpenGLESAPI = glfw.OpenGLESAPI
	NoAPI = glfw.NoAPI
	NativeContextAPI = glfw.NativeContextAPI
	EGLContextAPI = glfw.EGLContextAPI
	OSMesaContextAPI = glfw.OSMesaContextAPI
	NoRobustness = glfw.NoRobustness
	NoResetNotification = glfw.NoResetNotification
	LoseContextOnReset = glfw.LoseContextOnReset
	AnyReleaseBehavior = glfw.AnyReleaseBehavior
	ReleaseBehaviorFlush = glfw.ReleaseBehaviorFlush
	ReleaseBehaviorNone = glfw.ReleaseBehaviorNone
	OpenGLAnyProfile = glfw.OpenGLAnyProfile
	OpenGLCoreProfile = glfw.OpenGLCoreProfile
	OpenGLCompatProfile = glfw.OpenGLCompatProfile
	True = glfw.True
	False = glfw.False
	DontCare = glfw.DontCare
	HatCentered = glfw.HatCentered
	HatUp = glfw.HatUp
	HatRight = glfw.HatRight
	HatDown = glfw.HatDown
	HatLeft = glfw.HatLeft
	HatRightUp = glfw.HatRightUp
	HatRightDown = glfw.HatRightDown
	HatLeftUp = glfw.HatLeftUp
	HatLeftDown = glfw.HatLeftDown
	KeyUnknown = Key(glfw.KeyUnknown)
	KeySpace = Key(glfw.KeySpace)
	KeyApostrophe = Key(glfw.KeyApostrophe)
	KeyComma = Key(glfw.KeyComma)
	KeyMinus = Key(glfw.KeyMinus)
	KeyPeriod = Key(glfw.KeyPeriod)
	KeySlash = Key(glfw.KeySlash)
	Key0 = Key(glfw.Key0)
	Key1 = Key(glfw.Key1)
	Key2 = Key(glfw.Key2)
	Key3 = Key(glfw.Key3)
	Key4 = Key(glfw.Key4)
	Key5 = Key(glfw.Key5)
	Key6 = Key(glfw.Key6)
	Key7 = Key(glfw.Key7)
	Key8 = Key(glfw.Key8)
	Key9 = Key(glfw.Key9)
	KeySemicolon = Key(glfw.KeySemicolon)
	KeyEqual = Key(glfw.KeyEqual)
	KeyA = Key(glfw.KeyA)
	KeyB = Key(glfw.KeyB)
	KeyC = Key(glfw.KeyC)
	KeyD = Key(glfw.KeyD)
	KeyE = Key(glfw.KeyE)
	KeyF = Key(glfw.KeyF)
	KeyG = Key(glfw.KeyG)
	KeyH = Key(glfw.KeyH)
	KeyI = Key(glfw.KeyI)
	KeyJ = Key(glfw.KeyJ)
	KeyK = Key(glfw.KeyK)
	KeyL = Key(glfw.KeyL)
	KeyM = Key(glfw.KeyM)
	KeyN = Key(glfw.KeyN)
	KeyO = Key(glfw.KeyO)
	KeyP = Key(glfw.KeyP)
	KeyQ = Key(glfw.KeyQ)
	KeyR = Key(glfw.KeyR)
	KeyS = Key(glfw.KeyS)
	KeyT = Key(glfw.KeyT)
	KeyU = Key(glfw.KeyU)
	KeyV = Key(glfw.KeyV)
	KeyW = Key(glfw.KeyW)
	KeyX = Key(glfw.KeyX)
	KeyY = Key(glfw.KeyY)
	KeyZ = Key(glfw.KeyZ)
	KeyLeftBracket = Key(glfw.KeyLeftBracket)
	KeyBackslash = Key(glfw.KeyBackslash)
	KeyRightBracket = Key(glfw.KeyRightBracket)
	KeyGraveAccent = Key(glfw.KeyGraveAccent)
	KeyWorld1 = Key(glfw.KeyWorld1)
	KeyWorld2 = Key(glfw.KeyWorld2)
	KeyEscape = Key(glfw.KeyEscape)
	KeyEnter = Key(glfw.KeyEnter)
	KeyTab = Key(glfw.KeyTab)
	KeyBackspace = Key(glfw.KeyBackspace)
	KeyInsert = Key(glfw.KeyInsert)
	KeyDelete = Key(glfw.KeyDelete)
	KeyRight = Key(glfw.KeyRight)
	KeyLeft = Key(glfw.KeyLeft)
	KeyDown = Key(glfw.KeyDown)
	KeyUp = Key(glfw.KeyUp)
	KeyPageUp = Key(glfw.KeyPageUp)
	KeyPageDown = Key(glfw.KeyPageDown)
	KeyHome = Key(glfw.KeyHome)
	KeyEnd = Key(glfw.KeyEnd)
	KeyCapsLock = Key(glfw.KeyCapsLock)
	KeyScrollLock = Key(glfw.KeyScrollLock)
	KeyNumLock = Key(glfw.KeyNumLock)
	KeyPrintScreen = Key(glfw.KeyPrintScreen)
	KeyPause = Key(glfw.KeyPause)
	KeyF1 = Key(glfw.KeyF1)
	KeyF2 = Key(glfw.KeyF2)
	KeyF3 = Key(glfw.KeyF3)
	KeyF4 = Key(glfw.KeyF4)
	KeyF5 = Key(glfw.KeyF5)
	KeyF6 = Key(glfw.KeyF6)
	KeyF7 = Key(glfw.KeyF7)
	KeyF8 = Key(glfw.KeyF8)
	KeyF9 = Key(glfw.KeyF9)
	KeyF10 = Key(glfw.KeyF10)
	KeyF11 = Key(glfw.KeyF11)
	KeyF12 = Key(glfw.KeyF12)
	KeyF13 = Key(glfw.KeyF13)
	KeyF14 = Key(glfw.KeyF14)
	KeyF15 = Key(glfw.KeyF15)
	KeyF16 = Key(glfw.KeyF16)
	KeyF17 = Key(glfw.KeyF17)
	KeyF18 = Key(glfw.KeyF18)
	KeyF19 = Key(glfw.KeyF19)
	KeyF20 = Key(glfw.KeyF20)
	KeyF21 = Key(glfw.KeyF21)
	KeyF22 = Key(glfw.KeyF22)
	KeyF23 = Key(glfw.KeyF23)
	KeyF24 = Key(glfw.KeyF24)
	KeyF25 = Key(glfw.KeyF25)
	KeyKP0 = Key(glfw.KeyKP0)
	KeyKP1 = Key(glfw.KeyKP1)
	KeyKP2 = Key(glfw.KeyKP2)
	KeyKP3 = Key(glfw.KeyKP3)
	KeyKP4 = Key(glfw.KeyKP4)
	KeyKP5 = Key(glfw.KeyKP5)
	KeyKP6 = Key(glfw.KeyKP6)
	KeyKP7 = Key(glfw.KeyKP7)
	KeyKP8 = Key(glfw.KeyKP8)
	KeyKP9 = Key(glfw.KeyKP9)
	KeyKPDecimal = Key(glfw.KeyKPDecimal)
	KeyKPDivide = Key(glfw.KeyKPDivide)
	KeyKPMultiply = Key(glfw.KeyKPMultiply)
	KeyKPSubtract = Key(glfw.KeyKPSubtract)
	KeyKPAdd = Key(glfw.KeyKPAdd)
	KeyKPEnter = Key(glfw.KeyKPEnter)
	KeyKPEqual = Key(glfw.KeyKPEqual)
	KeyLeftShift = Key(glfw.KeyLeftShift)
	KeyLeftControl = Key(glfw.KeyLeftControl)
	KeyLeftAlt = Key(glfw.KeyLeftAlt)
	KeyLeftSuper = Key(glfw.KeyLeftSuper)
	KeyRightShift = Key(glfw.KeyRightShift)
	KeyRightControl = Key(glfw.KeyRightControl)
	KeyRightAlt = Key(glfw.KeyRightAlt)
	KeyRightSuper = Key(glfw.KeyRightSuper)
	KeyMenu = Key(glfw.KeyMenu)
	KeyLast = Key(glfw.KeyLast)
	ModShift = glfw.ModShift
	ModControl = glfw.ModControl
	ModAlt = glfw.ModAlt
	ModSuper = glfw.ModSuper
	ModCapsLock = glfw.ModCapsLock
	ModNumLock = glfw.ModNumLock
	MouseButton1 = glfw.MouseButton1
	MouseButton2 = glfw.MouseButton2
	MouseButton3 = glfw.MouseButton3
	MouseButton4 = glfw.MouseButton4
	MouseButton5 = glfw.MouseButton5
	MouseButton6 = glfw.MouseButton6
	MouseButton7 = glfw.MouseButton7
	MouseButton8 = glfw.MouseButton8
	MouseButtonLast = glfw.MouseButtonLast
	MouseButtonLeft = glfw.MouseButtonLeft
	MouseButtonRight = glfw.MouseButtonRight
	MouseButtonMiddle = glfw.MouseButtonMiddle
	Connected = glfw.Connected
	Disconnected = glfw.Disconnected
	ArrowCursor = glfw.ArrowCursor
	IBeamCursor = glfw.IBeamCursor
	CrosshairCursor = glfw.CrosshairCursor
	HandCursor = glfw.HandCursor
	HResizeCursor = glfw.HResizeCursor
	VResizeCursor = glfw.VResizeCursor
	Joystick1 = glfw.Joystick1
	Joystick2 = glfw.Joystick2
	Joystick3 = glfw.Joystick3
	Joystick4 = glfw.Joystick4
	Joystick5 = glfw.Joystick5
	Joystick6 = glfw.Joystick6
	Joystick7 = glfw.Joystick7
	Joystick8 = glfw.Joystick8
	Joystick9 = glfw.Joystick9
	Joystick10 = glfw.Joystick10
	Joystick11 = glfw.Joystick11
	Joystick12 = glfw.Joystick12
	Joystick13 = glfw.Joystick13
	Joystick14 = glfw.Joystick14
	Joystick15 = glfw.Joystick15
	Joystick16 = glfw.Joystick16
	JoystickLast = glfw.JoystickLast
	CursorMode = glfw.CursorMode
	StickyKeysMode = glfw.StickyKeysMode
	StickyMouseButtonsMode = glfw.StickyMouseButtonsMode
	LockKeyMods = glfw.LockKeyMods
	RawMouseMotion = glfw.RawMouseMotion
	CocoaFrameNAME = glfw.CocoaFrameNAME
	X11ClassName = glfw.X11ClassName
	X11InstanceName = glfw.X11InstanceName
	ContextRevision = glfw.ContextRevision
	RedBits = glfw.RedBits
	GreenBits = glfw.GreenBits
	BlueBits = glfw.BlueBits
	AlphaBits = glfw.AlphaBits
	DepthBits = glfw.DepthBits
	StencilBits = glfw.StencilBits
	AccumRedBits = glfw.AccumRedBits
	AccumGreenBits = glfw.AccumGreenBits
	AccumBlueBits = glfw.AccumBlueBits
	AccumAlphaBits = glfw.AccumAlphaBits
	AuxBuffers = glfw.AuxBuffers
	Stereo = glfw.Stereo
	Samples = glfw.Samples
	SRGBCapable = glfw.SRGBCapable
	RefreshRate = glfw.RefreshRate
	DoubleBuffer = glfw.DoubleBuffer
	CocoaGraphicsSwitching = glfw.CocoaGraphicsSwitching
	CocoaRetinaFramebuffer = glfw.CocoaRetinaFramebuffer
	ClientAPI = glfw.ClientAPI
	ContextVersionMajor = glfw.ContextVersionMajor
	ContextVersionMinor = glfw.ContextVersionMinor
	ContextRobustness = glfw.ContextRobustness
	ContextReleaseBehavior = glfw.ContextReleaseBehavior
	OpenGLForwardCompatible = glfw.OpenGLForwardCompatible
	OpenGLDebugContext = glfw.OpenGLDebugContext
	OpenGLProfile = glfw.OpenGLProfile
	ContextCreationAPI = glfw.ContextCreationAPI
	Focused = glfw.Focused
	Iconified = glfw.Iconified
	Maximized = glfw.Maximized
	Visible = glfw.Visible
	Hovered = glfw.Hovered
	Resizable = glfw.Resizable
	Decorated = glfw.Decorated
	Floating = glfw.Floating
	AutoIconify = glfw.AutoIconify
	CenterCursor = glfw.CenterCursor
	TransparentFramebuffer = glfw.TransparentFramebuffer
	FocusOnShow = glfw.FocusOnShow
	ScaleToMonitor = glfw.ScaleToMonitor
	JoystickHatButtons = glfw.JoystickHatButtons
	CocoaChdirResources = glfw.CocoaChdirResources
	CocoaMenubar = glfw.CocoaMenubar
	ButtonA = glfw.ButtonA
	ButtonB = glfw.ButtonB
	ButtonX = glfw.ButtonX
	ButtonY = glfw.ButtonY
	ButtonLeftBumper = glfw.ButtonLeftBumper
	ButtonRightBumper = glfw.ButtonRightBumper
	ButtonBack = glfw.ButtonBack
	ButtonStart = glfw.ButtonStart
	ButtonGuide = glfw.ButtonGuide
	ButtonLeftThumb = glfw.ButtonLeftThumb
	ButtonRightThumb = glfw.ButtonRightThumb
	ButtonDpadUp = glfw.ButtonDpadUp
	ButtonDpadRight = glfw.ButtonDpadRight
	ButtonDpadDown = glfw.ButtonDpadDown
	ButtonDpadLeft = glfw.ButtonDpadLeft
	ButtonLast = glfw.ButtonLast
	ButtonCross = glfw.ButtonCross
	ButtonCircle = glfw.ButtonCircle
	ButtonSquare = glfw.ButtonSquare
	ButtonTriangle = glfw.ButtonTriangle
	AxisLeftX = glfw.AxisLeftX
	AxisLeftY = glfw.AxisLeftY
	AxisRightX = glfw.AxisRightX
	AxisRightY = glfw.AxisRightY
	AxisLeftTrigger = glfw.AxisLeftTrigger
	AxisRightTrigger = glfw.AxisRightTrigger
	AxisLast = glfw.AxisLast
	Release = glfw.Release
	Press = glfw.Press
	Repeat = glfw.Repeat
)

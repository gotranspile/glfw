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
)`
// +build linux

package windstorm

/*
#define XK_LATIN1
#define XK_3270
#define XK_MISCELLANY
#include <X11/keysymdef.h>
#include <X11/Xlib.h>
*/
import "C"

// Action is a constant type for key and button states.
type Action int

// Constants representing button and key states.
const (
	_ Action = iota
	Press
	Release
)

// MouseButton is a constant type for physical mouse buttons.
type MouseButton int

// Constants representing mouse buttons.
const (
	MouseButton1 = MouseButton(C.Button1)
	MouseButton2 = MouseButton(C.Button2)
	MouseButton3 = MouseButton(C.Button3)
	MouseButton4 = MouseButton(C.Button4)
	MouseButton5 = MouseButton(C.Button5)
)

// Key is a constant type for physical keyboard keys.
type Key int

// Constants representing keyboard buttons.
const (
	KeySpace        = Key(C.XK_space)
	KeyApostrophe   = Key(C.XK_apostrophe)
	KeyComma        = Key(C.XK_comma)
	KeyMinus        = Key(C.XK_minus)
	KeyPeriod       = Key(C.XK_period)
	KeySlash        = Key(C.XK_slash)
	Key0            = Key(C.XK_0)
	Key1            = Key(C.XK_1)
	Key2            = Key(C.XK_2)
	Key3            = Key(C.XK_3)
	Key4            = Key(C.XK_4)
	Key5            = Key(C.XK_5)
	Key6            = Key(C.XK_6)
	Key7            = Key(C.XK_7)
	Key8            = Key(C.XK_8)
	Key9            = Key(C.XK_9)
	KeySemicolon    = Key(C.XK_semicolon)
	KeyEqual        = Key(C.XK_equal)
	KeyA            = Key(C.XK_A)
	KeyB            = Key(C.XK_B)
	KeyC            = Key(C.XK_C)
	KeyD            = Key(C.XK_D)
	KeyE            = Key(C.XK_E)
	KeyF            = Key(C.XK_F)
	KeyG            = Key(C.XK_G)
	KeyH            = Key(C.XK_H)
	KeyI            = Key(C.XK_I)
	KeyJ            = Key(C.XK_J)
	KeyK            = Key(C.XK_K)
	KeyL            = Key(C.XK_L)
	KeyM            = Key(C.XK_M)
	KeyN            = Key(C.XK_N)
	KeyO            = Key(C.XK_O)
	KeyP            = Key(C.XK_P)
	KeyQ            = Key(C.XK_Q)
	KeyR            = Key(C.XK_R)
	KeyS            = Key(C.XK_S)
	KeyT            = Key(C.XK_T)
	KeyU            = Key(C.XK_U)
	KeyV            = Key(C.XK_V)
	KeyW            = Key(C.XK_W)
	KeyX            = Key(C.XK_X)
	KeyY            = Key(C.XK_Y)
	KeyZ            = Key(C.XK_Z)
	KeyLeftBracket  = Key(C.XK_bracketleft)
	KeyBackslash    = Key(C.XK_backslash)
	KeyRightBracket = Key(C.XK_bracketright)
	KeyGraveAccent  = Key(C.XK_grave)
	KeyEscape       = Key(C.XK_Escape)
	KeyEnter        = Key(C.XK_Return)
	KeyTab          = Key(C.XK_Tab)
	KeyBackspace    = Key(C.XK_BackSpace)
	KeyInsert       = Key(C.XK_Insert)
	KeyDelete       = Key(C.XK_Delete)
	KeyRight        = Key(C.XK_Right)
	KeyLeft         = Key(C.XK_Left)
	KeyDown         = Key(C.XK_Down)
	KeyUp           = Key(C.XK_Up)
	KeyPageUp       = Key(C.XK_Page_Up)
	KeyPageDown     = Key(C.XK_Page_Down)
	KeyHome         = Key(C.XK_Home)
	KeyEnd          = Key(C.XK_End)
	KeyCapsLock     = Key(C.XK_Caps_Lock)
	KeyScrollLock   = Key(C.XK_Scroll_Lock)
	KeyNumLock      = Key(C.XK_Num_Lock)
	KeyPrKeyScreen  = Key(C.XK_3270_PrintScreen)
	KeyPause        = Key(C.XK_Pause)
	KeyF1           = Key(C.XK_F1)
	KeyF2           = Key(C.XK_F2)
	KeyF3           = Key(C.XK_F3)
	KeyF4           = Key(C.XK_F4)
	KeyF5           = Key(C.XK_F5)
	KeyF6           = Key(C.XK_F6)
	KeyF7           = Key(C.XK_F7)
	KeyF8           = Key(C.XK_F8)
	KeyF9           = Key(C.XK_F9)
	KeyF10          = Key(C.XK_F10)
	KeyF11          = Key(C.XK_F11)
	KeyF12          = Key(C.XK_F12)
	KeyF13          = Key(C.XK_F13)
	KeyF14          = Key(C.XK_F14)
	KeyF15          = Key(C.XK_F15)
	KeyF16          = Key(C.XK_F16)
	KeyF17          = Key(C.XK_F17)
	KeyF18          = Key(C.XK_F18)
	KeyF19          = Key(C.XK_F19)
	KeyF20          = Key(C.XK_F20)
	KeyF21          = Key(C.XK_F21)
	KeyF22          = Key(C.XK_F22)
	KeyF23          = Key(C.XK_F23)
	KeyF24          = Key(C.XK_F24)
	KeyKp0          = Key(C.XK_KP_0)
	KeyKp1          = Key(C.XK_KP_1)
	KeyKp2          = Key(C.XK_KP_2)
	KeyKp3          = Key(C.XK_KP_3)
	KeyKp4          = Key(C.XK_KP_4)
	KeyKp5          = Key(C.XK_KP_5)
	KeyKp6          = Key(C.XK_KP_6)
	KeyKp7          = Key(C.XK_KP_7)
	KeyKp8          = Key(C.XK_KP_8)
	KeyKp9          = Key(C.XK_KP_9)
	KeyKpDecimal    = Key(C.XK_KP_Decimal)
	KeyKpDivide     = Key(C.XK_KP_Divide)
	KeyKpMultiply   = Key(C.XK_KP_Multiply)
	KeyKpSubtract   = Key(C.XK_KP_Subtract)
	KeyKpAdd        = Key(C.XK_KP_Add)
	KeyKpEnter      = Key(C.XK_KP_Enter)
	KeyKpEqual      = Key(C.XK_KP_Equal)
	KeyLeftShift    = Key(C.XK_Shift_L)
	KeyRightShift   = Key(C.XK_Shift_R)
	KeyLeftControl  = Key(C.XK_Control_L)
	KeyLeftAlt      = Key(C.XK_Alt_L)
	KeyLeftSuper    = Key(C.XK_Super_L)
	KeyRightControl = Key(C.XK_Control_R)
	KeyRightAlt     = Key(C.XK_Alt_R)
	KeyRightSuper   = Key(C.XK_Super_R)
	KeyMenu         = Key(C.XK_Menu)
)

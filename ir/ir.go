// package IR provides an intermediate representation so that
// keys can be transported cross platform and translated at the receiving
// platform. For example keys are registered on a iOS device and sent to a
// windows box.
package ir

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

const (
	XK_Control      = 0x3B
	XK_RightShift   = 0x3C
	XK_RightControl = 0x3E
	XK_Command      = 0x37
	XK_Shift        = 0x38
	XK_ALT          = 0x383

	XK_A              = 0x00
	XK_S              = 0x01
	XK_D              = 0x02
	XK_F              = 0x03
	XK_H              = 0x04
	XK_G              = 0x05
	XK_Z              = 0x06
	XK_X              = 0x07
	XK_C              = 0x08
	XK_V              = 0x09
	XK_B              = 0x0B
	XK_Q              = 0x0C
	XK_W              = 0x0D
	XK_E              = 0x0E
	XK_R              = 0x0F
	XK_Y              = 0x10
	XK_T              = 0x11
	XK_1              = 0x12
	XK_2              = 0x13
	XK_3              = 0x14
	XK_4              = 0x15
	XK_6              = 0x16
	XK_5              = 0x17
	XK_EQUAL          = 0x18
	XK_9              = 0x19
	XK_7              = 0x1A
	XK_MINUS          = 0x1B
	XK_8              = 0x1C
	XK_0              = 0x1D
	XK_RightBracket   = 0x1E
	XK_O              = 0x1F
	XK_U              = 0x20
	XK_LeftBracket    = 0x21
	XK_I              = 0x22
	XK_P              = 0x23
	XK_L              = 0x25
	XK_J              = 0x26
	XK_Quote          = 0x27
	XK_K              = 0x28
	XK_SEMICOLON      = 0x29
	XK_BACKSLASH      = 0x2A
	XK_COMMA          = 0x2B
	XK_SLASH          = 0x2C
	XK_N              = 0x2D
	XK_M              = 0x2E
	XK_Period         = 0x2F
	XK_GRAVE          = 0x32
	XK_KeypadDecimal  = 0x41
	XK_KeypadMultiply = 0x43
	XK_KeypadPlus     = 0x45
	XK_KeypadClear    = 0x47
	XK_KeypadDivide   = 0x4B
	XK_KeypadEnter    = 0x4C
	XK_KeypadMinus    = 0x4E
	XK_KeypadEquals   = 0x51
	XK_Keypad0        = 0x52
	XK_Keypad1        = 0x53
	XK_Keypad2        = 0x54
	XK_Keypad3        = 0x55
	XK_Keypad4        = 0x56
	XK_Keypad5        = 0x57
	XK_Keypad6        = 0x58
	XK_Keypad7        = 0x59
	XK_Keypad8        = 0x5B
	XK_Keypad9        = 0x5C

	XK_ENTER         = 0x24
	XK_TAB           = 0x30
	XK_SPACE         = 0x31
	XK_DELETE        = 0x33
	XK_ESC           = 0x35
	XK_CAPSLOCK      = 0x39
	XK_Option        = 0x3A
	XK_RightOption   = 0x3D
	XK_Function      = 0x3F
	XK_F17           = 0x40
	XK_VOLUMEUP      = 0x48
	XK_VOLUMEDOWN    = 0x49
	XK_MUTE          = 0x4A
	XK_F18           = 0x4F
	XK_F19           = 0x50
	XK_F20           = 0x5A
	XK_F5            = 0x60
	XK_F6            = 0x61
	XK_F7            = 0x62
	XK_F3            = 0x63
	XK_F8            = 0x64
	XK_F9            = 0x65
	XK_F11           = 0x67
	XK_F13           = 0x69
	XK_F16           = 0x6A
	XK_F14           = 0x6B
	XK_F10           = 0x6D
	XK_F12           = 0x6F
	XK_F15           = 0x71
	XK_HELP          = 0x72
	XK_HOME          = 0x73
	XK_PAGEUP        = 0x74
	XK_ForwardDelete = 0x75
	XK_F4            = 0x76
	XK_END           = 0x77
	XK_F2            = 0x78
	XK_PAGEDOWN      = 0x79
	XK_F1            = 0x7A
	XK_LEFT          = 0x7B
	XK_RIGHT         = 0x7C
	XK_DOWN          = 0x7D
	XK_UP            = 0x7E
)

// Todo: we should use map[int]rune
var stringMap = map[int]string{
	XK_A:              "a",
	XK_S:              "s",
	XK_D:              "d",
	XK_F:              "f",
	XK_H:              "h",
	XK_G:              "g",
	XK_Z:              "z",
	XK_X:              "x",
	XK_C:              "c",
	XK_V:              "v",
	XK_B:              "b",
	XK_Q:              "q",
	XK_W:              "w",
	XK_E:              "e",
	XK_R:              "r",
	XK_Y:              "y",
	XK_T:              "t",
	XK_1:              "1",
	XK_2:              "2",
	XK_3:              "3",
	XK_4:              "4",
	XK_6:              "6",
	XK_5:              "5",
	XK_EQUAL:          "=",
	XK_9:              "9",
	XK_7:              "7",
	XK_MINUS:          "-",
	XK_8:              "8",
	XK_0:              "0",
	XK_RightBracket:   "]",
	XK_O:              "o",
	XK_U:              "u",
	XK_LeftBracket:    "[",
	XK_I:              "i",
	XK_P:              "p",
	XK_L:              "l",
	XK_J:              "j",
	XK_Quote:          "'",
	XK_K:              "k",
	XK_SEMICOLON:      ";",
	XK_BACKSLASH:      `\`,
	XK_COMMA:          ",",
	XK_SLASH:          "/",
	XK_N:              "n",
	XK_M:              "m",
	XK_Period:         ".",
	XK_GRAVE:          "`",
	XK_KeypadDecimal:  ".",
	XK_KeypadMultiply: "*",
	XK_KeypadPlus:     "+",
	XK_KeypadClear:    "",
	XK_KeypadDivide:   "/",
	XK_KeypadEnter:    "",
	XK_KeypadMinus:    "-",
	XK_KeypadEquals:   "=",
	XK_Keypad0:        "0",
	XK_Keypad1:        "1",
	XK_Keypad2:        "2",
	XK_Keypad3:        "3",
	XK_Keypad4:        "4",
	XK_Keypad5:        "5",
	XK_Keypad6:        "6",
	XK_Keypad7:        "7",
	XK_Keypad8:        "8",
	XK_Keypad9:        "9",

	XK_TAB: "	",
	XK_SPACE: " ",
}

// Special characters are ignored?
func ToString(keys []int) (string, error) {
	buf := &bytes.Buffer{}
	var is = struct {
		shift bool
	}{}
	for _, key := range keys {
		ks, ok := stringMap[key]
		if !ok {
			return "", fmt.Errorf("Key %v not found", key)
		}
		switch {
		default:
			_, err := buf.WriteString(ks)
			if err != nil {
				return "", err
			}
		case is.shift:
			_, err := buf.WriteString(strings.ToUpper(ks))
			if err != nil {
				return "", err
			}
			is.shift = false

		case key == XK_Shift || key == XK_RightShift:
			is.shift = true
		case key == XK_ALT || key == XK_Control ||
			key == XK_RightControl || key == XK_Command:
			// todo
		}
	}
	return buf.String(), nil
}

// Attempts to convert strings to the key code
// sequence required to construct them in an input field
func ToKeys(s string) ([]int, error) {
	var keys []int
	for _, v := range s {
		var key int
		var found bool
		for k, ks := range stringMap {
			if strings.EqualFold(ks, string(v)) {
				key = k
				found = true
				break
			}
		}
		switch {
		case found == false:
			return nil, fmt.Errorf("key for rune %v (hex: %+q) not found, s  %s", v, v, s)
		case unicode.IsUpper(v):
			keys = append(keys, XK_Shift)
		}
		keys = append(keys, key)
	}
	return keys, nil
}

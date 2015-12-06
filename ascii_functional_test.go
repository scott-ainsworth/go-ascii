package ascii

import (
	"bytes"
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

//******************************************************************************
// Test the Is* functions.
//******************************************************************************

const cntrl = "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0A\x0B\x0C\x0D\x0E\x0F" +
	"\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1A\x1B\x1C\x1D\x1E\x1F" +
	"\x7F"
const space = " "
const white = space + "\r\n\t\v\f"
const digit = "0123456789"
const hexdg = digit + "ABCDEF" + "abcdef"
const upper = "ABCEDFGHIJKLMNOPQRSTUVWXYZ"
const lower = "abcedfghijklmnopqrstuvwxyz"
const alpha = upper + lower
const alnum = upper + lower + digit
const punct = "!\"#%&'()*,-./:;?@[\\]_{}"
const symbl = "$+<=>^`|~"
const graph = alnum + punct + symbl
const print = alnum + punct + symbl + space
const ascii = alnum + punct + symbl + space + cntrl

func invert(chars string) string {
	var buffer bytes.Buffer
	for c := 0x00; c <= 0xFF; c++ {
		if strings.IndexByte(chars, byte(c)) < 0 {
			buffer.WriteByte(byte(c))
		}
	}
	return buffer.String()
}

func testIsFunction(t *testing.T, f func(byte) bool, fName, chars, tDesc, fDesc string) {
	Convey(tDesc+", then "+fName+" should return true.", t, func() {
		for _, c := range chars {
			So(f(byte(c)), ShouldBeTrue)
		}
	})
	Convey(fDesc+", then "+fName+" should return false.", t, func() {
		for _, c := range invert(chars) {
			So(f(byte(c)), ShouldBeFalse)
		}
	})
}

func TestIsASCII(t *testing.T) {
	testIsFunction(t, IsASCII, "IsASCII(c)", ascii,
		"Given a byte c that is a valid ASCII character (0x00..0x7F)",
		"Given a byte c that is not a valid ASCII character")
}

func TestIsLetter(t *testing.T) {
	testIsFunction(t, IsLetter, "IsLetter(c)", alpha,
		"Given a byte c that is a letter (A-Z, a-z)",
		"Given a byte c that is not a letter")
}

func TestIsUpper(t *testing.T) {
	testIsFunction(t, IsUpper, "IsUpper(c)", upper,
		"Given a byte c that is an uppercase letter (A-Z)",
		"Given a byte c that is not an uppercase letter")
}

func TestIsLower(t *testing.T) {
	testIsFunction(t, IsLower, "IsLower(c)", lower,
		"Given a byte c that is a lowercase letter (a-z)",
		"Given a byte c that is not a lowercase letter")
}

func TestIsDigit(t *testing.T) {
	testIsFunction(t, IsDigit, "IsDigit(c)", digit,
		"Given a byte c that is a decimal digit (0-9)",
		"Given a byte c that is not a decimal digit")
}

func TestIsHexDigit(t *testing.T) {
	testIsFunction(t, IsHexDigit, "IsHexDigit(c)", hexdg,
		"Given a byte c that is a hexidecmail decimal digit (0-9, A-F, a-f)",
		"Given a byte c that is not a hexidecmail decimal digit")
}

func TestIsAlnum(t *testing.T) {
	testIsFunction(t, IsAlnum, "IsAlnum(c)", alnum,
		"Given a byte c that is alphanumeric (A-Z, a-z, 0-9)",
		"Given a byte c that is not alphanumeric")
}

func TestIsSpace(t *testing.T) {
	testIsFunction(t, IsSpace, "IsSpace(c)", white,
		"Given a byte c that is a space character (space, \\r, \\n, \\t, \\v, \\f)",
		"Given a byte c that is not a space character")
}

func TestIsPunct(t *testing.T) {
	testIsFunction(t, IsPunct, "IsPunct(c)", punct,
		"Given a byte c that is punctuation (printable except space, letters and digits)",
		"Given a byte c that is not punctuation")
}

func TestIsSymbol(t *testing.T) {
	testIsFunction(t, IsSymbol, "ISymbol(c)", punct,
		"Given a byte c that is a symbol ($+<=>^`|~)",
		"Given a byte c that is not a symbol")
}

func TestIsControl(t *testing.T) {
	testIsFunction(t, IsControl, "IsControl(c)", cntrl,
		"Given a byte c that is a control character (0x00..0x1F, 0x7F)",
		"Given a byte c that is not a control character")
}

func TestIsPrint(t *testing.T) {
	testIsFunction(t, IsPrint, "IsPrint(c)", print,
		"Given a byte c that is a printable (' '..'~')",
		"Given a byte c that is not a printable")
}

func TestIsGraph(t *testing.T) {
	testIsFunction(t, IsGraph, "IsGraph(c)", graph,
		"Given a byte c that has a graphical representation ('!'..'~')",
		"Given a byte c that has not a graphical representation")
}

//******************************************************************************
// Test the conversion functions.
//******************************************************************************

func testToFunction(t *testing.T, f func(byte) byte, fName, fromChars, toChars, tDesc, tResult, fDesc, fResult string) {
	Convey(tDesc+", "+fName+" should return "+tResult+".", t, func() {
		for i, c := range fromChars {
			So(f(byte(c)), ShouldEqual, byte(toChars[i]))
		}
	})
	Convey(fDesc+", "+fName+" should return "+fResult+".", t, func() {
		for _, c := range invert(fromChars) {
			So(f(byte(c)), ShouldEqual, byte(c))
		}
	})
}

func TestToUpper(t *testing.T) {
	testToFunction(t, ToUpper, "ToUpper(c)", lower, upper,
		"Given a byte c is a lower case letter ('a'..'z')", "the upper case equivalent",
		"Given a byte c is not a lower case letter", "c")
}

func TestToLower(t *testing.T) {
	testToFunction(t, ToLower, "ToLower(c)", upper, lower,
		"Given a byte c is an upper case letter ('A'..'Z')", "the lower case equivalent",
		"Given a byte c is not an upper case letter", "c")
}

// eof

package ternary

import (
	"testing"
	"time"

	"github.com/theTardigrade/tests"
)

func TestString(t *testing.T) {
	x, y := "x", "y"

	if msg, fail := tests.Message(x, String(true, x, y)); fail {
		t.Error(msg)
	}

	if msg, fail := tests.Message(y, String(false, x, y)); fail {
		t.Error(msg)
	}
}

func TestInt(t *testing.T) {
	x, y := 1, 2

	if msg, fail := tests.Message(x, Int(true, x, y)); fail {
		t.Error(msg)
	}

	if msg, fail := tests.Message(y, Int(false, x, y)); fail {
		t.Error(msg)
	}
}

func TestUint(t *testing.T) {
	x, y := uint(1), uint(2)

	if msg, fail := tests.Message(x, Uint(true, x, y)); fail {
		t.Error(msg)
	}

	if msg, fail := tests.Message(y, Uint(false, x, y)); fail {
		t.Error(msg)
	}
}

func TestDuration(t *testing.T) {
	x, y := time.Second, time.Minute

	if msg, fail := tests.Message(x, Duration(true, x, y)); fail {
		t.Error(msg)
	}

	if msg, fail := tests.Message(y, Duration(false, x, y)); fail {
		t.Error(msg)
	}
}

func TestByte(t *testing.T) {
	x, y := byte('x'), byte('y')

	if msg, fail := tests.Message(x, Byte(true, x, y)); fail {
		t.Error(msg)
	}

	if msg, fail := tests.Message(y, Byte(false, x, y)); fail {
		t.Error(msg)
	}
}

func TestRune(t *testing.T) {
	x, y := rune('世'), rune('界')

	if msg, fail := tests.Message(x, Rune(true, x, y)); fail {
		t.Error(msg)
	}

	if msg, fail := tests.Message(y, Rune(false, x, y)); fail {
		t.Error(msg)
	}
}

func TestFloat32(t *testing.T) {
	x, y := float32(1.1), float32(1.2)

	if msg, fail := tests.Message(x, Float32(true, x, y)); fail {
		t.Error(msg)
	}

	if msg, fail := tests.Message(y, Float32(false, x, y)); fail {
		t.Error(msg)
	}
}

func TestFloat64(t *testing.T) {
	x, y := float64(1.5), float64(1.6)

	if msg, fail := tests.Message(x, Float64(true, x, y)); fail {
		t.Error(msg)
	}

	if msg, fail := tests.Message(y, Float64(false, x, y)); fail {
		t.Error(msg)
	}
}

func TestInterface(t *testing.T) {
	x, y := interface{}(600), interface{}("y")

	if msg, fail := tests.Message(x, Interface(true, x, y)); fail {
		t.Error(msg)
	}

	if msg, fail := tests.Message(y, Interface(false, x, y)); fail {
		t.Error(msg)
	}
}

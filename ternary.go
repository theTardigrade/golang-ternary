package ternary

import "time"

func String(condition bool, x, y string) string {
	if condition {
		return x
	}

	return y
}

func Int(condition bool, x, y int) int {
	if condition {
		return x
	}

	return y
}

func Uint(condition bool, x, y uint) uint {
	if condition {
		return x
	}

	return y
}

func Time(condition bool, x, y time.Time) time.Time {
	if condition {
		return x
	}

	return y
}

func Duration(condition bool, x, y time.Duration) time.Duration {
	if condition {
		return x
	}

	return y
}

func Byte(condition bool, x, y byte) byte {
	if condition {
		return x
	}

	return y
}

func Rune(condition bool, x, y rune) rune {
	if condition {
		return x
	}

	return y
}

func Float32(condition bool, x, y float32) float32 {
	if condition {
		return x
	}

	return y
}

func Float64(condition bool, x, y float64) float64 {
	if condition {
		return x
	}

	return y
}

func Interface(condition bool, x, y interface{}) interface{} {
	if condition {
		return x
	}

	return y
}

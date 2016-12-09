package main

import "testing"

var bunnytests = []struct {
	in  string
	len int
}{
	{"foo", 3},
	{"f o o", 3},
	{"ADVENT", 6},
	{"A(1x5)BC", 7},
	{"(3x3)XYZ", 9},
	{"A(2x2)BCD(2x2)EFG", 11},
	{"(6x1)(1x3)A", 6},
	{"X(8x2)(3x3)ABCY", 18},
}

var bunnytwotests = []struct {
	in  string
	len int
}{
	{"foo", 3},
	{"(3x3)XYZ", 9},
	{"X(8x2)(3x3)ABCY", 20},
	{"(27x12)(20x12)(13x14)(7x10)(1x12)A", 241920},
	{"(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN", 445},
}

func TestBunnyUncompress(t *testing.T) {
	for _, tt := range bunnytests {
		output, err := UncompressBunny(tt.in)
		if err != nil {
			t.Errorf("%s should not cause an error", tt.in)
		}
		if len(output) != tt.len {
			t.Errorf("'%s' => '%s', len was not %d", tt.in, output, tt.len)
		}
	}
}

func TestBunnyUncompressTwo(t *testing.T) {
	for _, tt := range bunnytwotests {
		length := UncompressBunnyTwo(tt.in)
		if length != tt.len {
			t.Errorf("'%s' length was %d, not %d", tt.in, length, tt.len)
		}
	}
}

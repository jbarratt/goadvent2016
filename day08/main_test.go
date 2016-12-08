package main

import "testing"

func TestRect(t *testing.T) {
	s := NewScreen(3, 5)
	lp := s.litPixels()
	if lp != 0 {
		t.Error("Expected 0, got ", lp)
	}
	s.rect(1, 1)
	lp = s.litPixels()
	if lp != 1 {
		t.Error("Expected 1, got ", lp)
	}
	s.rect(3, 5)
	lp = s.litPixels()
	if lp != 15 {
		t.Error("Expected 15, got ", lp)
	}
	// invalid values exit the program which is annoying to test for
	// I should refactor that
	// s.rect(6, 10)
}

func TestRectCmd(t *testing.T) {
	s := NewScreen(3, 5)
	s.command("rect 1x1")
	lp := s.litPixels()
	if lp != 1 {
		t.Error("Expected 1, got ", lp)
	}
	s.command("rect 3x5")
	lp = s.litPixels()
	if lp != 15 {
		t.Error("Expected 15, got ", lp)
	}
}

func TestRotateColumn(t *testing.T) {
	s := NewScreen(2, 3)
	s.rect(1, 1)
	s.rotateColumn(0, 1)
	if !s.display[0][1] || s.litPixels() != 1 {
		t.Error("only column 0 row 1 pixel should be lit ")
	}
	s.rotateColumn(0, 3)
	if !s.display[0][1] || s.litPixels() != 1 {
		t.Error("only column 0 row 1 pixel should be lit ")
	}
	s.rotateColumn(0, 2)
	if !s.display[0][0] || s.litPixels() != 1 {
		t.Error("only column 0 row 0 pixel should be lit ")
	}
}

func TestRotateRow(t *testing.T) {
	s := NewScreen(3, 2)
	s.rect(1, 1)
	s.rotateRow(0, 1)
	if !s.display[1][0] || s.litPixels() != 1 {
		t.Error("only row 0 column 1 pixel should be lit ")
	}
	s.rotateRow(0, 3)
	if !s.display[1][0] || s.litPixels() != 1 {
		t.Error("only row 0 column 1 pixel should be lit ")
	}
	s.rotateRow(0, 2)
	if !s.display[0][0] || s.litPixels() != 1 {
		s.print()
		t.Error("only column 0 row 0 pixel should be lit ")
	}
}

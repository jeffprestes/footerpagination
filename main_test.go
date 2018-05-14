package main

import (
	"testing"
)

func TestPrintFooter(t *testing.T) {
	var stringTest, stringReturned string

	stringTest = "1 2 3 4 5 6 ... 9 10"
	stringReturned = printFooter(4, 10, 2, 2)
	if stringTest != stringReturned {
		t.Errorf("Return was incorrect. Got [%s], wanted: [%s]", stringTest, stringReturned)
		return
	}
	stringTest = "1 ... 4 5"
	stringReturned = printFooter(4, 5, 1, 0)
	if stringTest != stringReturned {
		t.Errorf("Return was incorrect. Got [%s], wanted: [%s]", stringTest, stringReturned)
		return
	}
	stringTest = "1 ... 4 ... 20"
	stringReturned = printFooter(4, 20, 1, 0)
	if stringTest != stringReturned {
		t.Errorf("Return was incorrect. Got [%s], wanted: [%s]", stringTest, stringReturned)
		return
	}
	stringTest = "1 ... 4 ... 20"
	stringReturned = printFooter(4, 20, 1, 2)
	if stringTest == stringReturned {
		t.Error("Return was incorrect. The results must be different")
		return
	}
}

func TestIsWithinAroundRange(t *testing.T) {
	if isWithinAroundRange(1, 1, 0) {
		t.Error("Return was incorrect. The result must be false")
		return
	}
	if !isWithinAroundRange(3, 2, 2) {
		t.Error("Return was incorrect. The result must be true")
		return
	}
	if !isWithinAroundRange(6, 5, 2) {
		t.Error("Return was incorrect. The result must be true")
		return
	}
}

func TestIsWithinBoundaries(t *testing.T) {
	if !isWithinBoundaries(1, 1, 5, 1) {
		t.Error("Return was incorrect. The result must be true")
		return
	}
	if !isWithinBoundaries(1, 5, 5, 1) {
		t.Error("Return was incorrect. The result must be true")
		return
	}
	if !isWithinBoundaries(9, 9, 10, 2) {
		t.Error("Return was incorrect. The result must be true")
		return
	}
}

func TestShouldPrintGap(t *testing.T) {
	if shouldPrintGap(3, true, false, 5) {
		t.Error("Return was incorrect. The result must be false")
		return
	}
	if !shouldPrintGap(4, false, false, 6) {
		t.Error("Return was incorrect. The result must be true")
		return
	}
	if shouldPrintGap(7, true, true, 5) {
		t.Error("Return was incorrect. The result must be false")
		return
	}
	if !shouldPrintGap(7, true, false, 5) {
		t.Error("Return was incorrect. The result must be true")
		return
	}
}

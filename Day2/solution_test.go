package main

import "testing"

func TestSolution(t *testing.T) {
	result := solution("./test.txt")
	if result != 8 {
		t.Error("Excpected 8 got")
	}
}

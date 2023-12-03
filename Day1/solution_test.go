package main

import (
	"testing"
)

func TestDay1(t *testing.T) {
	result := solution("./test.txt")

	if result != 281 {
		t.Errorf("got %d, expected %d", result, 142)
	}
}

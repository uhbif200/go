package main

import "testing"

func TestSum(t *testing.T) {
	s := sum(1, 1)
	if s != 2 {
		t.Error("Sum of 1 and 1 not equal to 2")
	}
}

package main

import (
	"testing"
)

func TestDoTheThing1(t *testing.T) {
	out := DoTheThing1("testinput.txt")

	if out!=4361 {
		t.Errorf("output not equal 4361 instead go: %v", out)
	}
}

func TestDoTheThing2(t *testing.T) {
	out := DoTheThing2("testinput2.txt")

	if out!=467835{
		t.Errorf("output not equal 467835 instead go: %v", out)
	}
}

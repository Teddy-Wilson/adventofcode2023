package main

import "testing"

func TestDoTheThing1 (t *testing.T){

	output := DoTheThing1("testinput.txt")
	if output != 13 {
		t.Errorf("test failed, expected 13 got %v", output)
	}
}

func TestDoTheThing2 (t *testing.T){
	output := DoTheThing2("testinput.txt")
	if output != 30 {
		t.Errorf("test failed, expected 30 got %v", output)
	}
}

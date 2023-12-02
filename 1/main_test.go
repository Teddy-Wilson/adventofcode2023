package main

import (
	"fmt"
	"testing"
)

func TestSolveTheThing1(t *testing.T){
	output := SolveTheThing1("1.testinput.txt")
	if output != 142 {
		fmt.Errorf("output not equal to 142 got: %v", output)
	}
}

func TestSolveTheThing2(t *testing.T){
	output := SolveTheThing2("12.testinput.txt")
	if output != 281 {
		fmt.Errorf("output not equal to 142 got: %v", output)
	}
}

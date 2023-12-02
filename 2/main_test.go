package main

import "testing"

func TestSolveTheThing1(t *testing.T){
	solution := SolveTheThing1("testinput.txt")
	if solution != 8 {
		t.Fatalf("exptected 8 got: \"%v\"", solution)
	}
}

func TestSolveTheThing2(t *testing.T){
	solution := SolveTheThing2("testinput2.txt")
	if solution != 2286 {
		t.Fatalf("exptected 2286 got: \"%v\"", solution)
	}
}

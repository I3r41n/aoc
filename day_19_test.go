package main

import (
	"testing"
)

func Test_solve_1(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{"puzzles/day_19/example_1"}, 19114},
		{"input", args{"puzzles/day_19/input.txt"}, 362930},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_1(tt.args.url); got != tt.want {
				t.Errorf("solve_1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve_2(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{"puzzles/day_19/example_1"}, 167409079868000},
		{"input", args{"puzzles/day_19/input.txt"}, 116365820987729},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_2(tt.args.url); got != tt.want {
				t.Errorf("solve_2() = %v, want %v", got, tt.want)
			}
		})
	}
}

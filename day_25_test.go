package main

import "testing"

func Test_solve_25_1(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{"puzzles/day_25/example_1"}, 54},
		{"input", args{"puzzles/day_25/input.txt"}, 592171},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_25_1(tt.args.url); got != tt.want {
				t.Errorf("solve_25_1() = %v, want %v", got, tt.want)
			}
		})
	}
}

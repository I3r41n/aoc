package main

import (
	"testing"
)

func Test_solve_21_1(t *testing.T) {
	type args struct {
		url   string
		steps int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {"example 1", args{"puzzles/day_21/example_1", 6}, 16},
		{"example 1", args{"puzzles/day_21/input.txt", 64}, 3651},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_21_1(tt.args.url, tt.args.steps); got != tt.want {
				t.Errorf("solve_21_1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve_21_2(t *testing.T) {
	type args struct {
		url   string
		steps int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {"example 1", args{"puzzles/day_21/example_1", 5000}, 1594},
		{"example 1", args{"puzzles/day_21/input.txt", 26501365}, 607334325965751},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_21_2(tt.args.url, tt.args.steps); got != tt.want {
				t.Errorf("solve_21_2() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"testing"
)

func Test_solve_7_part1(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{url: "puzzles/day_7/example_1"}, 6440},
		{"input 1", args{url: "puzzles/day_7/input.txt"}, 250370104},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_7_part1(tt.args.url); got != tt.want {
				t.Errorf("solve_7_part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve_7_part2(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {"example 1", args{url: "puzzles/day_7/example_1"}, 5905},
		{"input 1", args{url: "puzzles/day_7/input.txt"}, 251735672},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_7_part2(tt.args.url); got != tt.want {
				t.Errorf("solve_7_part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"testing"
)

func Test_solve_8_part1(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{url: "puzzles/day_8/example_1"}, 2},
		{"example 2", args{url: "puzzles/day_8/example_2"}, 6},
		{"input 1", args{url: "puzzles/day_8/input.txt"}, 12643},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_8_part1(tt.args.url); got != tt.want {
				t.Errorf("solve_8_part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve_8_part2(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{url: "puzzles/day_8/example_1"}, 2},
		{"example 2", args{url: "puzzles/day_8/example_2"}, 6},
		{"example 3", args{url: "puzzles/day_8/example_3"}, 6},
		{"input 1", args{url: "puzzles/day_8/input.txt"}, 13133452426987},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_8_part2(tt.args.url); got != tt.want {
				t.Errorf("solve_8_part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

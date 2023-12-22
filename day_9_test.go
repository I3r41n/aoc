package main

import (
	"testing"
)

func Test_solve_9_part1(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{url: "puzzles/day_9/example_1"}, 114},
		{"input", args{url: "puzzles/day_9/input.txt"}, 1725987467},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_9_part1(tt.args.url); got != tt.want {
				t.Errorf("solve_9_part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve9_part2(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{url: "puzzles/day_9/example_1"}, 2},
		{"input", args{url: "puzzles/day_9/input.txt"}, 971},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve9_part2(tt.args.url); got != tt.want {
				t.Errorf("solve9_part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

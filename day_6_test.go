package main

import (
	"testing"
)

func Test_solve_6_part1(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{url: "puzzles/day_6/example_1"}, 288},
		{"input 1", args{url: "puzzles/day_6/input.txt"}, 140220},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_6_part1(tt.args.url); got != tt.want {
				t.Errorf("solve_6_part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve_6_part2(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{url: "puzzles/day_6/example_1"}, 71503},
		{"input 1", args{url: "puzzles/day_6/input.txt"}, 39570185},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_6_part2(tt.args.url); got != tt.want {
				t.Errorf("solve_6_part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"testing"
)

func Test_solve_5_part1(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{url: "puzzles/day_5/example_1"}, 35},
		// {"input", args{url: "puzzles/day_5/input.txt"}, 340994526},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_5_part1(tt.args.url); got != tt.want {
				t.Errorf("solve_5_part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve_5_part2(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{url: "puzzles/day_5/example_1"}, 46},
		{"input", args{url: "puzzles/day_5/input.txt"}, 52210644},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_5_part2(tt.args.url); got != tt.want {
				t.Errorf("solve_5_part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve_5_part2_reverse(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{url: "puzzles/day_5/example_1"}, 46},
		{"input", args{url: "puzzles/day_5/input.txt"}, 52210644},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_5_part2_reverse(tt.args.url); got != tt.want {
				t.Errorf("solve_5_part2_reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

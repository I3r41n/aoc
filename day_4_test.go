package main

import (
	"testing"
)

func Test_solve_4_part1(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{url: "puzzles/day_4/example_1"}, 13},
		{"input", args{url: "puzzles/day_4/input.txt"}, 22897},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_4_part1(tt.args.url); got != tt.want {
				t.Errorf("solve_4_part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve_4_part2(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{url: "puzzles/day_4/example_1"}, 30},
		{"input", args{url: "puzzles/day_4/input.txt"}, 5095824},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_4_part2(tt.args.url); got != tt.want {
				t.Errorf("solve_4_part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

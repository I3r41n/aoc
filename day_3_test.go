package main

import (
	"testing"
)

func Test_solve_day3(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{url: "puzzles/day_3/example_1"}, 4361},
		{"input", args{url: "puzzles/day_3/input.txt"}, 544664},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_day3(tt.args.url); got != tt.want {
				t.Errorf("solve_day3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve2_day3(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{url: "puzzles/day_3/example_1"}, 467835},
		{"input", args{url: "puzzles/day_3/input.txt"}, 544664},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve2_day3(tt.args.url); got != tt.want {
				t.Errorf("solve2_day3() = %v, want %v", got, tt.want)
			}
		})
	}
}

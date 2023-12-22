package main

import (
	"testing"
)

func Test_solve_day_17_1(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{url: "puzzles/day_17/example_1"}, 102},
		{"input", args{url: "puzzles/day_17/input.txt"}, 698},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_day_17_1(tt.args.url); got != tt.want {
				t.Errorf("solve_day_17_1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve_day_17_2(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{url: "puzzles/day_17/example_1"}, 94},
		{"example 2", args{url: "puzzles/day_17/example_2"}, 71},
		{"input_test", args{url: "puzzles/day_17/input_test.txt"}, 1210},
		{"input_test", args{url: "puzzles/day_17/input_test_2.txt"}, 1382},
		{"input", args{url: "puzzles/day_17/input.txt"}, 825},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_day_17_2(tt.args.url); got != tt.want {
				t.Errorf("solve_day_17_2() = %v, want %v", got, tt.want)
			}
		})
	}
}

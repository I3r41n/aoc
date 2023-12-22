package main

import "testing"

func Test_solve_13_1(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{url: "puzzles/day_13/example_0"}, 208},
		{"example 1", args{url: "puzzles/day_13/example_1"}, 405},
		{"input1", args{url: "puzzles/day_13/input.txt"}, 36448},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_13_1(tt.args.url); got != tt.want {
				t.Errorf("solve_13_1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve_13_2(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 0", args{url: "puzzles/day_13/example_2"}, 600},
		{"example 1", args{url: "puzzles/day_13/example_1"}, 400},
		{"input1", args{url: "puzzles/day_13/input.txt"}, 35799},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_13_2(tt.args.url); got != tt.want {
				t.Errorf("solve_13_2() = %v, want %v", got, tt.want)
			}
		})
	}
}

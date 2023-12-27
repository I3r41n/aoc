package main

import "testing"

func Test_solve_23_1(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{"example 1", args{url: "puzzles/day_23/example_1"}, 94},
		{"input", args{url: "puzzles/day_23/input.txt"}, 2438},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_23_1(tt.args.url); got != tt.want {
				t.Errorf("solve_23_1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve_23_2(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{url: "puzzles/day_23/example_1"}, 154},
		{"input", args{url: "puzzles/day_23/input.txt"}, 6658},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_23_2(tt.args.url); got != tt.want {
				t.Errorf("solve_23_2() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_solve_day_16_1(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 0", args{url: "puzzles/day_16/example_0"}, 8},
		{"example cycle", args{url: "puzzles/day_16/example_cycle"}, 7},
		{"example 1", args{url: "puzzles/day_16/example_1"}, 46},
		{"input", args{url: "puzzles/day_16/input.txt"}, 8112},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_day_16_1(tt.args.url); got != tt.want {
				t.Errorf("solve_day_16_1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve_day_16_2(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{url: "puzzles/day_16/example_1"}, 51},
		{"input", args{url: "puzzles/day_16/input.txt"}, 8314},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_day_16_2(tt.args.url); got != tt.want {
				t.Errorf("solve_day_16_2() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_solve_day_15_1(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"HASH", args{url: "puzzles/day_15/example_0"}, 52},
		{"example 1", args{url: "puzzles/day_15/example_1"}, 1320},
		{"input", args{url: "puzzles/day_15/input.txt"}, 505459},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_day_15_1(tt.args.url); got != tt.want {
				t.Errorf("solve_day_15_1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve_day_15_2(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{url: "puzzles/day_15/example_1"}, 145},
		{"input", args{url: "puzzles/day_15/input.txt"}, 228508},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_day_15_2(tt.args.url); got != tt.want {
				t.Errorf("solve_day_15_2() = %v, want %v", got, tt.want)
			}
		})
	}
}

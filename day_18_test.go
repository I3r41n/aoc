package main

import "testing"

func Test_solve_day_18_1(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{url: "puzzles/day_18/example_1"}, 62},
		{"input", args{url: "puzzles/day_18/input.txt"}, 42317},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_day_18_1(tt.args.url); got != tt.want {
				t.Errorf("solve_day_18_1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve_day_18_2(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{url: "puzzles/day_18/example_1"}, 952408144115},
		{"input", args{url: "puzzles/day_18/input.txt"}, 83605563360288},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_day_18_2(tt.args.url); got != tt.want {
				t.Errorf("solve_day_18_2() = %v, want %v", got, tt.want)
			}
		})
	}
}

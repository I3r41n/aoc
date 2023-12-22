package main

import "testing"

func Test_solve_22_1(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{url: "puzzles/day_22/example_1"}, 5},
		{"input", args{"puzzles/day_22/input.txt"}, 501},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_22_1(tt.args.url); got != tt.want {
				t.Errorf("solve_22_1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve_22_2(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{url: "puzzles/day_22/example_1"}, 7},
		{"example 1", args{"puzzles/day_22/input.txt"}, 80948},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_22_2(tt.args.url); got != tt.want {
				t.Errorf("solve_22_2() = %v, want %v", got, tt.want)
			}
		})
	}
}

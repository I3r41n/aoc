package main

import "testing"

func Test_solve_12_1(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{url: "puzzles/day_12/example_1"}, 21},
		{"input", args{url: "puzzles/day_12/input.txt"}, 7017},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_12_1(tt.args.url); got != tt.want {
				t.Errorf("solve_12_1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve_12_2(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{url: "puzzles/day_12/example_1"}, 525152},
		{"input", args{url: "puzzles/day_12/input.txt"}, 527570479489},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_12_2(tt.args.url); got != tt.want {
				t.Errorf("solve_12_1() = %v, want %v", got, tt.want)
			}
		})
	}
}

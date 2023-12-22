package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name      string
		args      args
		wantTotal int
	}{
		{"puzzle1", args{url: "puzzles/day_2/example_1"}, 8},
		{"part1", args{url: "puzzles/day_2/input.txt"}, 2600},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTotal := solve(tt.args.url); gotTotal != tt.wantTotal {
				t.Errorf("solve() = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}

func Test_solve2(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name      string
		args      args
		wantTotal int
	}{
		{"puzzle1", args{url: "puzzles/day_2/example_1"}, 2286},
		{"part1", args{url: "puzzles/day_2/input.txt"}, 86036},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTotal := solve2(tt.args.url); gotTotal != tt.wantTotal {
				t.Errorf("solve2() = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}

package main

import (
	"testing"
)

func Test_solve10_1(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{url: "puzzles/day_10/example_1"}, 4},
		{"example 2", args{url: "puzzles/day_10/example_2"}, 8},
		{"example 2.1", args{url: "puzzles/day_10/example_2_extra"}, 8},
		{"input", args{url: "puzzles/day_10/input.txt"}, 7066},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve10_1(tt.args.url); got != tt.want {
				t.Errorf("solve10_1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve10_2(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{url: "puzzles/day_10/example_1"}, 1},
		{"example 2", args{url: "puzzles/day_10/example_2"}, 1},
		{"example 2.1", args{url: "puzzles/day_10/example_2_extra"}, 1},
		{"example 3", args{url: "puzzles/day_10/example_3"}, 4},
		{"example 3", args{url: "puzzles/day_10/example_4"}, 8},
		{"example 3", args{url: "puzzles/day_10/example_5"}, 10},
		{"input", args{url: "puzzles/day_10/input.txt"}, 401},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve10_2(tt.args.url); got != tt.want {
				t.Errorf("solve10_2() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"testing"
)

func Test_solve_day_14_1(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{url: "puzzles/day_14/example_1"}, 136},
		{"example 1", args{url: "puzzles/day_14/input.txt"}, 109833},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_day_14_1(tt.args.url); got != tt.want {
				t.Errorf("solve_day_14_1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve_day_14_2(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{url: "puzzles/day_14/example_1"}, 64},
		{"example 1", args{url: "puzzles/day_14/input.txt"}, 109833},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_day_14_2(tt.args.url); got != tt.want {
				t.Errorf("solve_day_14_2() = %v, want %v", got, tt.want)
			}
		})
	}
}

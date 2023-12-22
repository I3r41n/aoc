package main

import "testing"

func Test_solve_day_20_1(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{url: "puzzles/day_20/example_1"}, 32000000},
		{"example 2", args{url: "puzzles/day_20/example_2"}, 11687500},
		{"example 2", args{url: "puzzles/day_20/input.txt"}, 818649769},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_day_20_1(tt.args.url); got != tt.want {
				t.Errorf("solve_day_20_1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve_day_20_2(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 2", args{url: "puzzles/day_20/input.txt"}, 246313604784977},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_day_20_2(tt.args.url); got != tt.want {
				t.Errorf("solve_day_20_2() = %v, want %v", got, tt.want)
			}
		})
	}
}

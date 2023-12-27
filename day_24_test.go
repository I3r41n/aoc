package main

import "testing"

func Test_solve_24_1(t *testing.T) {
	type args struct {
		url string
		x   int
		y   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{"puzzles/day_24/example_1", 7, 27}, 2},
		{"input", args{"puzzles/day_24/input.txt", 200000000000000, 400000000000000}, 25433},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_24_1(tt.args.url, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("solve_24_1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve_24_2(t *testing.T) {
	type args struct {
		url string
		x   int
		y   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{"puzzles/day_24/example_1", 7, 27}, 47},
		{"input", args{"puzzles/day_24/input.txt", 200000000000000, 400000000000000}, 885093461440405},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_24_2(tt.args.url, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("solve_24_2() = %v, want %v", got, tt.want)
			}
		})
	}
}

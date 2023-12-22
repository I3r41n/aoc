package main

import "testing"

func Test_solve11_1(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{url: "puzzles/day_11/example_1"}, 374},
		{"input", args{url: "puzzles/day_11/input.txt"}, 9274989},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve11_1(tt.args.url); got != tt.want {
				t.Errorf("solve11_1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve11_2(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{url: "puzzles/day_11/example_1"}, 82000210},
		{"input", args{url: "puzzles/day_11/input.txt"}, 357134560737},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve11_2(tt.args.url); got != tt.want {
				t.Errorf("solve11_2() = %v, want %v", got, tt.want)
			}
		})
	}
}

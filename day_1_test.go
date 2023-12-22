package main

import "testing"

func Test_getTotalCalibration(t *testing.T) {
	type args struct{ puzzle string }
	tests := []struct {
		name      string
		args      args
		wantTotal int64
	}{
		{"puzzle1", args{puzzle: "puzzles/day_1/day_1_example_1"}, 142},
		{"input1", args{puzzle: "puzzles/day_1/input_day_1.txt"}, 53386},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTotal := getTotalCalibration(tt.args.puzzle); gotTotal != tt.wantTotal {
				t.Errorf("getTotalCalibration() = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}

func Test_getTotalCalibration_2(t *testing.T) {
	type args struct{ puzzle string }
	tests := []struct {
		name      string
		args      args
		wantTotal int64
	}{
		{"puzzle2", args{puzzle: "puzzles/day_1/day_1_example_2"}, 281},
		{"input1", args{puzzle: "puzzles/day_1/input_day_1.txt"}, 53312},
		{"erro 1", args{puzzle: "puzzles/day_1/day_1_error_1"}, 26},
		{"erro 2", args{puzzle: "puzzles/day_1/day_1_error_2"}, 83},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTotal := getTotalCalibration_2(tt.args.puzzle); gotTotal != tt.wantTotal {
				t.Errorf("getTotalCalibration_2() = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}

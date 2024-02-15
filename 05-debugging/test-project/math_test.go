package main

import (
	"fmt"
	"testing"
)

func Test_add(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Addition 1",
			args: args{
				num1: 1,
				num2: 1,
			},
			want: 2,
		},
		{
			name: "Addition 2",
			args: args{
				num1: 10,
				num2: 10,
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.name == "Addition 2" {
				fmt.Println("Enabling Debugging Mode...")
			}

			if got := add(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sub(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Sub 1",
			args: args{
				num1: 1,
				num2: 1,
			},
			want: 0,
		},
		{
			name: "Sub 2",
			args: args{
				num1: 1,
				num2: 1,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sub(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

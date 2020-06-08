package util

import "testing"

func TestSize(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test-1",
			args: args{a: -1},
			want: "negative",
		},
		{
			name: "test0",
			args: args{a: 0},
			want: "zero",
		},
		{
			name: "test10",
			args: args{a: 9},
			want: "small",
		},
		{
			name: "test100",
			args: args{a: 99},
			want: "big",
		},
		{
			name: "test-1000",
			args: args{a: 999},
			want: "huge",
		},
		{
			name: "test-10000",
			args: args{a: 9999},
			want: "enormous",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Size(tt.args.a); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	type args struct {
		n1 int
		n2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{n1: 1, n2: 1},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.n1, tt.args.n2); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type args struct {
		n1 int
		n2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{n1: 1, n2: 2},
			want: 2,
		},
		{
			name: "test2",
			args: args{n1: 2, n2: 1},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.n1, tt.args.n2); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_min(t *testing.T) {
	type args struct {
		n1 int
		n2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{n1: 1, n2: 2},
			want: 1,
		},
		{
			name: "test2",
			args: args{n1: 2, n2: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args.n1, tt.args.n2); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}

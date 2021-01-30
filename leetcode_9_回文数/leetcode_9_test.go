package leetcode_9_回文数

import (
	"reflect"
	"testing"
)

func Test_intToSlice(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{99991},
			want: []int{9, 9, 9, 9, 1},
		},
		{
			name: "2",
			args: args{0},
			want: []int{0},
		},
		{
			name: "3",
			args: args{-123},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToSlice(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{123321},
			want: true,
		},
		{
			name: "2",
			args: args{-1123123123123123},
			want: false,
		},
		{
			name: "3",
			args: args{0},
			want: true,
		},
		{
			name: "4",
			args: args{123421},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindrome2(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{123321},
			want: true,
		},
		{
			name: "2",
			args: args{-1123123123123123},
			want: false,
		},
		{
			name: "3",
			args: args{0},
			want: true,
		},
		{
			name: "4",
			args: args{123421},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome2(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome2() = %v, want %v", got, tt.want)
			}
		})
	}
}

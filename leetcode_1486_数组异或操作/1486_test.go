package leetcode_1486_数组异或操作

import (
	"testing"
)

func Test_xorOperation(t *testing.T) {
	type args struct {
		n     int
		start int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n:     4,
				start: 3,
			},
			want: 8,
		},
		{
			args: args{
				n:     1,
				start: 7,
			},
			want: 7,
		},
		{
			args: args{
				n:     10,
				start: 5,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xorOperation(tt.args.n, tt.args.start); got != tt.want {
				t.Errorf("xorOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}

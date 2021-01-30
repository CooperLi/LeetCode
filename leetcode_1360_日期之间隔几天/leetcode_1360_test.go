package leetcode_1360_日期之间隔几天

import (
	"testing"
)

func TestDaysBetweenDates(t *testing.T) {
	type args struct {
		date1 string
		date2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				date1: "1971-01-01",
				date2: "1994-08-01",
			},
			want: 8613,
		},
		{
			name: "2",
			args: args{
				date1: "1994-10-16",
				date2: "1994-08-01",
			},
			want: 76,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := daysBetweenDates(tt.args.date1, tt.args.date2); got != tt.want {
				t.Errorf("daysBetweenDates() = %v, want %v", got, tt.want)
			}
			if got := daysBetweenDates2(tt.args.date1, tt.args.date2); got != tt.want {
				t.Errorf("daysBetweenDates() = %v, want %v", got, tt.want)
			}
		})
	}
}

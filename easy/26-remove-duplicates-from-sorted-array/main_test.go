package main

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	type result struct {
		k    int
		nums []int
	}
	tests := []struct {
		name string
		args args
		want result
	}{
		{
			name: "Example 1",
			args: args{nums: []int{1, 1, 2}},
			want: result{
				k:    2,
				nums: []int{1, 2},
			},
		},
		{
			name: "Example 2",
			args: args{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}},
			want: result{
				k:    5,
				nums: []int{0, 1, 2, 3, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want.k || !reflect.DeepEqual(tt.args.nums[0:tt.want.k], tt.want.nums) {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

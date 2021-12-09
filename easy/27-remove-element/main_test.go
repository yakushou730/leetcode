package main

import (
	"reflect"
	"testing"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
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
			args: args{nums: []int{3, 2, 2, 3}, val: 3},
			want: result{
				k:    2,
				nums: []int{2, 2},
			},
		},
		{
			name: "Example 2",
			args: args{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2},
			want: result{
				k:    5,
				nums: []int{0, 1, 3, 0, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want.k || !reflect.DeepEqual(tt.args.nums[0:tt.want.k], tt.want.nums) {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

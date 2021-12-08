package _089_duplicate_zeros

import (
	"reflect"
	"testing"
)

func Test_duplicateZeros(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{arr: []int{1, 0, 2, 3, 0, 4, 5, 0}},
			want: []int{1, 0, 0, 2, 3, 0, 0, 4},
		},
		{
			name: "Example 2",
			args: args{arr: []int{1, 2, 3}},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			duplicateZeros(tt.args.arr)
			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("arr after duplicateZeros() is %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}

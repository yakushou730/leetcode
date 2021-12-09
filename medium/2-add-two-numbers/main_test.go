package main

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example 1",
			args: args{
				l1: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
				l2: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
		{
			name: "Example 2",
			args: args{
				l1: &ListNode{
					Val:  0,
					Next: nil,
				},
				l2: &ListNode{
					Val:  0,
					Next: nil,
				},
			},
			want: &ListNode{
				Val:  0,
				Next: nil,
			},
		},
		{
			name: "Example 3",
			args: args{
				l1: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 9,
								Next: &ListNode{
									Val: 9,
									Next: &ListNode{
										Val: 9,
										Next: &ListNode{
											Val:  9,
											Next: nil,
										},
									},
								},
							},
						},
					},
				},
				l2: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val:  9,
								Next: nil,
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 8,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 0,
								Next: &ListNode{
									Val: 0,
									Next: &ListNode{
										Val: 0,
										Next: &ListNode{
											Val:  1,
											Next: nil,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

//
// func TestListNode_decimalValue(t *testing.T) {
// 	tests := []struct {
// 		name   string
// 		fields *ListNode
// 		want   int
// 	}{
// 		{
// 			name: "Test case 2 -> 4 -> 3",
// 			fields: &ListNode{
// 				Val: 2,
// 				Next: &ListNode{
// 					Val: 4,
// 					Next: &ListNode{
// 						Val:  3,
// 						Next: nil,
// 					},
// 				},
// 			},
// 			want: 342,
// 		},
// 		{
// 			name: "Test case 5 -> 6 -> 4",
// 			fields: &ListNode{
// 				Val: 5,
// 				Next: &ListNode{
// 					Val: 6,
// 					Next: &ListNode{
// 						Val:  4,
// 						Next: nil,
// 					},
// 				},
// 			},
// 			want: 465,
// 		},
// 		{
// 			name: "Test case 0",
// 			fields: &ListNode{
// 				Val:  0,
// 				Next: nil,
// 			},
// 			want: 0,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := decimalValue(tt.fields); got != tt.want {
// 				t.Errorf("decimalValue() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func Test_newListNode(t *testing.T) {
// 	type args struct {
// 		decimal int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want ListNode
// 	}{
// 		{
// 			name: "Test case 807",
// 			args: args{decimal: 807},
// 			want: ListNode{
// 				Val: 7,
// 				Next: &ListNode{
// 					Val: 0,
// 					Next: &ListNode{
// 						Val:  8,
// 						Next: nil,
// 					},
// 				},
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := newListNode(tt.args.decimal); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("newListNode() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

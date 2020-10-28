package dp

import "testing"

func Test_houseRobber(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{"1", []int{1, 2, 3, 1}, 4},
		{"2", []int{2, 7, 9, 3, 1}, 12},
		{"3", []int{2, 1, 1, 2}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := houseRobber(tt.args); got != tt.want {
				t.Errorf("houseRobber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHouseRobber2(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{"1", []int{2, 3, 2}, 3},
		{"2", []int{1, 2, 3, 1}, 4},
		{"3", []int{1, 1, 1, 1}, 2},
		{"4", []int{2, 1, 1, 1}, 3},
		{"5", []int{2, 1, 1, 1, 3, 5, 2, 3, 1, 1, 1, 1}, 12},
		{"6", []int{100, 1, 1, 1, 3, 200, 2, 3, 1, 1, 1, 1}, 305},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HouseRobber2(tt.args); got != tt.want {
				t.Errorf("HouseRobber2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHouseRobber3(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{root: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 1,
				},
			},
		}}, 7},
		{"2", args{root: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 5,
				Right: &TreeNode{
					Val: 1,
				},
			},
		}}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HouseRobber3(tt.args.root); got != tt.want {
				t.Errorf("HouseRobber3() = %v, want %v", got, tt.want)
			}
		})
	}
}

package binary_search

import (
	"fmt"
	"testing"
)

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{
			nums:   nil,
			target: 1,
		}, 0},
		{"2", args{
			nums:   []int{5},
			target: 1,
		}, 0},
		{"3", args{
			nums:   []int{5},
			target: 10,
		}, 1},
		{"4", args{
			nums:   []int{1, 3, 5, 6},
			target: 5,
		}, 2},
		{"5", args{
			nums:   []int{1, 3, 5, 6},
			target: 2,
		}, 1},
		{"6", args{
			nums:   []int{1, 3, 5, 6},
			target: 7,
		}, 4},
		{"7", args{
			nums:   []int{1, 3, 5, 6},
			target: 0,
		}, 0},
		{"8", args{
			nums:   []int{1, 3, 5},
			target: 4,
		}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("SearchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchInsert1(t *testing.T) {
	nums := []int{5}
	fmt.Println(SearchInsert(nums, 10))
}

func Benchmark_searchInsertPosition(b *testing.B) {
	//b.StopTimer()
	nums := make([]int, 10000)
	for i := 0; i < 100; i++ {
		nums[i] = i * 2
	}
	//b.StartTimer()
	for i := 0; i < b.N; i++ {
		SearchInsert(nums, 8009)
	}
}

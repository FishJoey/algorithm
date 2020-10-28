package dp

import (
	"testing"
)

var array = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4, 5}

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		array  []int
		expect int
	}{
		{array, 10},
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5}, 6},
	}
	for _, v := range tests {
		if MaxSubArray(v.array) != v.expect {
			t.Fatal("max sub array wrong result, array:", v.array)
		}
	}
}

func BenchmarkMaxSubString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxSubArray(array)
	}
}

func BenchmarkMaxSubString2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxSubArray2(array)
	}
}

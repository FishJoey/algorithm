package dp

import "testing"

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		array  []int
		expect int
	}{
		{[]int{2, 3, -2, 4}, 6},
		{[]int{2}, 2},
		{[]int{0}, 0},
		{[]int{-2}, -2},
		{[]int{2, 3, -2, 4, -1}, 48},
		{[]int{2, 3, -2, 4, -1, -2}, 48},
	}

	for _, v := range tests {
		if MaxProductSubarray(v.array) != v.expect {
			t.Fatal("max product subarray wrong result with array:", v.array)
		}
	}
}

func BenchmarkMaxProduct(b *testing.B) {
	nums := []int{-4, -3, -2}
	for i := 0; i < b.N; i++ {
		MaxProductSubarray(nums)
	}
}

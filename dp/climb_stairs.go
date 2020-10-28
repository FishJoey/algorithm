package dp

import "fmt"

func main() {
	//n := 5
	//fmt.Println(climb(n))
	//cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	//              1, 1,   2, 2, 3, 3,   4, 4, 5,   6
	cost := []int{10, 15, 20, 18}
	//			  10, 15, 30, 33,
	//            10, 10, 15, 30
	//            15, 10, 20
	//            15, 10, 10
	fmt.Println(minCostClimb(cost))
}

func climb(n int) int {
	if n < 2 {
		return n
	}

	method_i_1 := 1
	method_i_2 := 1
	var method int
	for i := 2; i <= n; i++ {
		method = method_i_1 + method_i_2
		method_i_1, method_i_2 = method, method_i_1
	}
	return method
}

func minCostClimb(cost []int) int {
	n := len(cost)
	if n == 1 {
		return cost[n-1]
	}

	fn_1 := cost[0]
	fn_2 := cost[1]
	var fn int
	for i := 2; i < n; i++ {
		fn = cost[i] + min(fn_1, fn_2)
		fn_2, fn_1 = fn, fn_2
	}
	return min(fn_1, fn)
}

func min(x, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}

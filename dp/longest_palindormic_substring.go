package dp

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	n := len(s)
	start, end := 0, 1
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			start, end = i, i+2
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if i > j+1 {
				if s[i] == s[j] && dp[j+1][i-1] {
					dp[j][i] = true
					if end-start < i-j+1 {
						start, end = j, i+1
					}
				}
			}
		}
	}
	return s[start:end]
}

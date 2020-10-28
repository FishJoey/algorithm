package dp

import "testing"

func TestLongest(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{"aaaabaaa", "aaabaaa"},
		{"absbx", "bsb"},
		{"abba", "abba"},
		{"abadd", "aba"},
		{"cbbd", "bb"},
		{"a", "a"},
		{"", ""},
	}

	for _, v := range tests {
		if longestPalindrome(v.input) != v.expect {
			t.Fatal("wrong with ", v.input)
		}
	}
}

func BenchmarkLongest(b *testing.B) {
	s := "aaaabaaa"
	for i := 0; i < b.N; i++ {
		longestPalindrome(s)
	}
}

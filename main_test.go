package main

import "testing"

func TestCountPalindromes(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"abc", 3}, // a, b, c
		{"aaa", 6}, // a,a,a,aa,aa,aaa
		{"", 0},
		{"a", 1},
		{"aba", 4},  // a,b,a,aba
		{"abba", 6}, // a,b,b,a,bb,abba
	}
	for _, c := range cases {
		if got := CountPalindromes(c.in); got != c.want {
			t.Errorf("CountPalindromes(%q) = %d, want %d", c.in, got, c.want)
		}
	}
}

// Property: must agree with brute force over a few strings.
func TestAgainstBrute(t *testing.T) {
	brute := func(s string) int {
		isPal := func(x string) bool {
			for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
				if x[i] != x[j] {
					return false
				}
			}
			return true
		}
		n := 0
		for i := 0; i < len(s); i++ {
			for j := i + 1; j <= len(s); j++ {
				if isPal(s[i:j]) {
					n++
				}
			}
		}
		return n
	}
	for _, s := range []string{"banana", "racecar", "abcba", "zzzz"} {
		if got, want := CountPalindromes(s), brute(s); got != want {
			t.Errorf("%q: %d vs brute %d", s, got, want)
		}
	}
}

package main

import "fmt"

// CountPalindromes returns how many substrings of s are palindromes
// (each occurrence counts, "aaa" has six). Expand around all 2n-1 centers,
// counting every successful expansion step.
func CountPalindromes(s string) int {
	count := 0
	expand := func(l, r int) {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			count++
			l--
			r++
		}
	}
	for i := 0; i < len(s); i++ {
		expand(i, i)   // odd lengths
		expand(i, i+1) // even lengths
	}
	return count
}

func main() {
	fmt.Println(CountPalindromes("abc"))
	fmt.Println(CountPalindromes("aaa"))
}

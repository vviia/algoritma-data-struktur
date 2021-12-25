/*

   A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

   Find the largest palindrome made from the product of two 3-digit numbers.
*/
package main

import (
	"fmt"
	"strconv"
)

func IsDecimalPalindromeNumber(n int64) bool {
	if n < 0 {
		n = -n
	}

	s := strconv.FormatInt(n, 10)
	bound := (len(s) / 2) + 1
	for i := 0; i < bound; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	currentMax := 0
	maxI := 0
	maxJ := 0
	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			product := i * j
			if currentMax > product {
				break
			}
			if IsDecimalPalindromeNumber(int64(product)) {
				currentMax = product
				maxI = i
				maxJ = j
				continue
			}
		}
	}
	fmt.Println("answer: ", currentMax)
	fmt.Println("i: ", maxI)
	fmt.Println("j: ", maxJ)
}

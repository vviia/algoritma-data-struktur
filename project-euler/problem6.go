/*

   The sum of the squares of the first ten natural numbers is,

   1^2 + 2^2 + ... + 10^2 = 385

   The square of the sum of the first ten natural numbers is,

   (1 + 2 + ... + 10)^2 = 55^2 = 3025

   Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

   Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/
package main

import (
	"fmt"
)

func main() {
	var sumsquare = uint64(1)
	var squaresum = uint64(1)
	for i := uint64(2); i <= 100; i++ {
		squaresum += i
		sumsquare += i * i
	}
	squaresum = squaresum * squaresum
	fmt.Print(squaresum)
	fmt.Print(" - ")
	fmt.Print(sumsquare)
	fmt.Print(" = ")
	fmt.Print(squaresum - sumsquare)
}

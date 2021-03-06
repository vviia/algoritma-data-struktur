/*

   Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:

   1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

   By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
*/

package main

import (
	"fmt"
)

func main() {
	var a, b, sum uint64 = 1, 2, 2
	for a+b < 4000000 {
		next := a + b
		//fmt.Println(next)
		if next%2 == 0 {
			sum += next
		}
		a = b
		b = next
	}
	fmt.Println("the sum of the even-valued terms below 4,000,000 is", sum)
}

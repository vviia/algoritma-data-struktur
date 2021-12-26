/*

A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
a2 + b2 = c2

For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/
package main

import (
	"fmt"
)

func findTriplet() (a, b, c int) {
	for a = 1; a < 500; a++ {
		for b = a + 1; b < 1000; b++ {
			c = 1000 - a - b
			if b > c {
				break
			}
			if a*a+b*b == c*c {
				return
			}
		}
	}

	// cannot find. return all 0
	a = 0
	b = 0
	c = 0
	return
}

func main() {
	a, b, c := findTriplet()
	fmt.Println(a, b, c)
	fmt.Println(a * b * c)
}

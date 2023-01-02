package main

import "fmt"

func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		num = num & (num - 1) // clear the rightmost set bit
		count += 1
	}

	return count
}

func main() {
	number := 22
	fmt.Printf("%d in binary is %b\n", number, number)
	fmt.Printf("the number of set bits in %d is %d\n", number, hammingWeight(22))
}

/*
output:
22 in binary is 10110
the number of set bits in 22 is 3
*/

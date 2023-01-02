package main

import "fmt"

func hammingWeight(num uint32) int {
	var count uint32
	for num > 0 {
		count += num & 1 // check last bit
		num = num >> 1
	}

	return int(count)
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

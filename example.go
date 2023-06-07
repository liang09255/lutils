package main

import (
	"fmt"
	"lutils/dfo"
)

func main() {
	exampleDefault()
}

func exampleDefault() {
	// sometimes when we design a rpc server
	// we will use nil pointer to define that it is not set
	// But when we use it, we expect nil to represent a default value
	nums1, nums2 := 1, 1
	fmt.Println(addTwoNum(nil, nil))       // 0
	fmt.Println(addTwoNum(&nums1, nil))    // 1
	fmt.Println(addTwoNum(nil, &nums2))    // 1
	fmt.Println(addTwoNum(&nums1, &nums2)) // 2
}

// This is a service which can add two numbers
func addTwoNum(a *int, b *int) int {
	return dfo.IntPtr(a) + dfo.IntPtr(b)
}

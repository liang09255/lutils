package main

import (
	"fmt"
	"github.com/liang09255/lutils/conv"
	"github.com/liang09255/lutils/dfo"
)

func main() {
	exampleDefault()
	exampleConv()
}

func exampleDefault() {
	// sometimes when we design a rpc server
	// we will use nil pointer to define that it is not set
	// But when we use it, we expect nil to represent a default value
	nums1, nums2 := 1, 1
	addTwoNum := func(a *int, b *int) int {
		// This is a service which can add two numbers
		return dfo.IntPtr(a) + dfo.IntPtr(b)
	}
	fmt.Println(addTwoNum(nil, nil))       // 0 dfo turn nil to 0
	fmt.Println(addTwoNum(&nums1, nil))    // 1
	fmt.Println(addTwoNum(nil, &nums2))    // 1
	fmt.Println(addTwoNum(&nums1, &nums2)) // 2
}

func exampleConv() {
	// If we sure that can convert without error, we can use conv to ignore error.
	// Error message will be logged by default logger
	// If we need to use our own logger, we can use conv.SetLogger to set it
	str := "1.0"
	fmt.Println(conv.ToInt(str))     // 1 (int)
	fmt.Println(conv.ToUint(str))    // 1 (uint)
	fmt.Println(conv.ToFloat64(str)) // 1 (float64)
	str = "true"
	fmt.Println(conv.ToBool(str)) // true (bool)
}

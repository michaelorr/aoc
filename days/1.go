package days

import (
	"fmt"

	"orr.co/adventofcode/data"
)

func Day1() {
	var nums1 []int
	var nums2 []int
	for x := range data.AsInts(data.Day(1)) {
		for _, i := range nums1 {
			for _, j := range nums2 {
				if x+i+j == 2020 {
					fmt.Println(x * i * j)
					return
				}
			}
		}
		nums1 = append(nums1, x)
		nums2 = append(nums2, x)
	}
}

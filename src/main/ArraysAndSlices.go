package main

import "fmt"

func main() {

	/*
		Arrays vs Slices
		arrays - fixed length
		slices - variable length, passed by Ref by default
					- underlying = array
	*/

	// slices
	// length=5, capacity=10
	// if you leave the 10 out, capacity=length
	nums := make([]int, 5, 10)
	println("length=", len(nums), "capacity=", cap(nums))


	nums2 := []int{1,2,3}
	println("length=", len(nums2), "capacity=", cap(nums2))
	println(nums2[0]) // getting value by index

	nums2[0] = 0
	println(nums2[0]) // getting value by index

	// inclusive:exclusive
	// other options = [:2] [1:]
	sliceOfSlice := nums2[0:2]
	fmt.Println("sliceOfSlice=",sliceOfSlice)

	newSlice := append(sliceOfSlice, 4)
	fmt.Println("newSlice=",newSlice)

	// append 2 slices
	newSlice2 := append(newSlice, nums2...)
	fmt.Println("newSlice2=",newSlice2)
}

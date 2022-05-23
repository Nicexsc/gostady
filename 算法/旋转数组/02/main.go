package main

import "fmt"

func rotate(nums []int, k int) {

	length := len(nums)
	var num = make([]int, length)
	// fmt.Println(len(num))
	count := k % length
	// fmt.Println(count)
	for j := 0; j < count; j++ {
		// fmt.Println(nums)
		num = make([]int, length)
		for i := 0; i < length-1; i++ {
			num[i+1] = nums[i]
			// fmt.Println(num)
			// fmt.Println(num)
		}
		num[0] = nums[length-1]
		// copy(nums,num)
		nums = num
		// fmt.Println(num,nums)
	}
	fmt.Println(nums)

}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)

}
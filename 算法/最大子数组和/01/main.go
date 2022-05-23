package main

import "fmt"

func maxSubArray(nums []int) int {
    max := nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i] + nums[i-1] > nums[i] {
            nums[i] += nums[i-1]
        }
        if nums[i] > max {
            max = nums[i]
        }
		fmt.Println(nums)
    }
    return max
}


func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	// nums := []int{-1,-2}
	n :=  maxSubArray(nums)
	fmt.Println(n)
}
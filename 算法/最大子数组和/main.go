package main

import "fmt"

func maxSubArray(nums []int) int {
	if len(nums) == 1{
		return  nums[0]
	}
	var sum int = 0
	// var tmp int = 0 
	var list []int
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			sum = sum + nums[j]
			// fmt.Println(sum)	
			list = append(list, sum)
		}
		sum = 0
	}
	fmt.Println(list)
	for t := 0; t < len(list); t++ {
		if list[t] > list[0]{
			list[0] = list[t]
		}
		
	}
    return list[0]
}
func main() {
	// nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	nums := []int{-1,-2}
	n :=  maxSubArray(nums)
	fmt.Println(n)
}
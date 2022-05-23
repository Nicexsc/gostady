package main

import "fmt"

func searchInsert(nums []int, target int) int {
	n := []int{}
	var v []int
	
	if nums[0] > target{
		v = append(n,target)
		nums = append(v, nums... )
		return 0
	}
	if nums[len(nums)-1] < target{
		nums = append(nums, target)
		return len(nums)
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < target && target <= nums[i+1]{
			tmp := append([]int{}, nums[i+1:]...)
			nums = append(nums[0:i+1],target )
			nums = append(nums, tmp...)
			fmt.Println(nums)
			
		}
		if nums[i] == target {
			return i

		}
		
	}
	return 0

}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 2
	num := searchInsert(nums, target)
	fmt.Println(num)

}
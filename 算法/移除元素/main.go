package main

import "fmt"

func removeElement(nums []int, val int) int {

	l := len(nums)
	if l == 0 {
		return 0
	}

	for i := 0; i < l; i ++ {
		
		if nums[i] == val && len(nums) != 1 {
			
			nums = append(nums[:i],nums[i+1:]... )
			if i == l-1{
				return len(nums)
			}
			break
	
		}else if len(nums) == 1 && nums[i] == val{
			nums = nums[:i]
		}
		if i == l-1{
			return len(nums)
		}
		
		
	}
	l = removeElement(nums,val)
	return l
	

}

func main() {

	nums := []int{3, 2,2 ,3}
	val := 3

	n := removeElement(nums, val)
	fmt.Println(n)

}
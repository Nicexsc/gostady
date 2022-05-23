package main

import "fmt"

var length int
func removeDuplicates(nums []int) []int {
	
	if len(nums) <=1{
		return nums
	}
	//遍历相邻之间是否有存在相同的，然后触发退出递归体
	count := 1
	for j := 0; j < len(nums) -1; j++ {
		if nums[j] != nums[j+1]{
			count ++
		}
		
	}
	if count == len(nums){
		return nums
	}
	//相邻之间相同则剔除它
	for i := 0; i <len(nums)-1; i++ {

		if nums[i] == nums[i+1]{
			nums = append(nums[:i+1], nums[i+2:]...)
			

		}
		
	}
	
	//[0,1,1,2,3,4]
	nums = removeDuplicates(nums)
	return nums
	
		
	
	
	
}
func main() {
	// a := []int{1, 1, 2, 2, 3, 4, 5}
	a := []int{0,0,1, 1,1,1,1,1, 2, 2, 3, 3, 4}
	// a := []int{1,1,1}
	k := removeDuplicates(a)
	fmt.Println(k)

}
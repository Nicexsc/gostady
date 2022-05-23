package main

import "fmt"

func removeDuplicates(nums []int) int {

	length := len(nums)
	//[1,1,2,3,3,3,5,6]
	//[1,1,1]
	for i := 0; i < length-1;  {
		j := i + 1
		
		if nums[i] == nums[j] {

			nums = append(nums[:j], nums[j+1:]...)
			// fmt.Println(nums)
			i = 0
			

		}else{
			i ++
		}
		// fmt.Println(i)
		length = len(nums)

	}
	// fmt.Println(nums)
	return len(nums)
	

}

func main() {
	// a := []int{1, 1, 2, 2, 3, 4, 5}
	a := []int{0,0,1, 1,1, 2, 2, 3, 3, 4}
	// a := []int{1,1,1}
	k := removeDuplicates(a)
	fmt.Println(k)

}
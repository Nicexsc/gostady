package main

import "fmt"


func twoSum(nums []int, target int) []int {
    lo := len(nums) 
	fmt.Println(nums,lo)
	var k []int
	var i int
    for i = 0 ; i < lo;i++ {
        for j := i+1;j<lo;j++{
			sum := nums[i] + nums[j]
			fmt.Println(sum)
            if sum == target{
				k = []int{i,j}
				return k
			}else{
				continue
			}
        }
		
    }
	return k
	

}


func main(){

	// nums := []int{3,2,4}
	// target := 6
	// n := twoSum(nums,target)
	// fmt.Println(n)

}
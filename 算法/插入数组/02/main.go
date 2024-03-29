package main

import "fmt"



func searchInsert(nums []int, target int) int {
    n := len(nums)
    left, right := 0, n - 1
    ans := n
    for left <= right {
        mid := (right + left) >> 1
        if target <= nums[mid] {
            ans = mid
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return ans
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 0
	num := searchInsert(nums, target)
	fmt.Println(num)

}
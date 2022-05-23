package main

import "fmt"

func rotate(nums []int, k int) {
    n := len(nums) //n = 6
    k %= n // k=2 
	//strat =0 , count = 2
	//控制左边2个长度循环
    for start, count := 0, gcd(k, n); start < count; start++ {
        pre, cur := nums[start], start
		//控制相隔空间长度循环{ok = cur != start}控制最多次停止
        for ok := true; ok; ok = cur != start {
            next := (cur + k) % n //4 ， 8 / 12
            nums[next], pre, cur = pre, nums[next], next
        }
    }
}

func gcd(a, b int) int {
    for a != 0 {
        a, b = b%a, a
    }
    return b
}


func main(){


	curl :=2 
	start := 1
	ok := true
	ok = curl != start 
	fmt.Println(ok)
	
}






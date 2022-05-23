// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数

package main

import (
	// "fmt"
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {	
	c := strconv.Itoa(x)
	a := []rune(c)
	num := len(a)	 
		j:= num 
		k := 0
		for i:=0 ;i<num/2 ;i++{
			j --
			if a[i] == a[j]{
				k++
			}
		}	
		return k == num/2
}

func main(){
	x := -121
	st := isPalindrome(x)
	fmt.Println(st)

}
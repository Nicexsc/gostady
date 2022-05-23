package main

import (
	"fmt"
	"sort"
)

func maxHeight(cuboids [][]int) int {
	n := len(cuboids)
	// 正方体内部排序 -> 升序
	for i := 0; i < n; i++ {
		sort.Ints(cuboids[i])
	}
	fmt.Println(cuboids)
	// 正方体排序(最短边) -> 升序
	sort.Slice(cuboids, func(i, j int) bool {
		if cuboids[i][0] != cuboids[j][0] {
			return cuboids[i][0] < cuboids[j][0]
		}
		if cuboids[i][1] != cuboids[j][1] {
			return cuboids[i][1] < cuboids[j][1]
		}
		return cuboids[i][2] < cuboids[j][2]
	})

	// 初始化dp
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = cuboids[i][2]
	}
	fmt.Println(dp)

	ans := 0
	// for i := 0; i < n; i++ {
	// 	for j := i + 1; j < n; j++ {
	// 		if cuboids[i][0] <= cuboids[j][0] && cuboids[i][1] <= cuboids[j][1] && cuboids[i][2] <= cuboids[j][2] {
	// 			//dp[1] = 45 + 50 , dp[2] = 45 + 50 + 95
	// 			dp[j] = max(dp[j], dp[i]+cuboids[j][2])
	// 		}
	// 	}
	// 	ans = max(ans, dp[i])
	// }
	num := len(cuboids[0]) -1
	fmt.Println(num)
	

	for _, values := range cuboids{
		
		fmt.Println(values)
		ans = values[num] + ans
		// if i == (n - 1){
		// 	return ans
		// }
		 

		
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main(){

	// a := [][]int{
	// 	{45,23,12},
	// 	{50,45,20},
	// 	{95,37,53},
	// 	{101,54,44},
		
	// }
	a := [][]int{
		{38,25,45},
		{76,35,3},
	}
	n := maxHeight(a)
	fmt.Println(n)


}
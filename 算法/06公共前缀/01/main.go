package main

import "fmt"

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    var lcp func(int, int) string
	//0 2
    lcp = func(start, end int) string {
        if start == end {
            return strs[start]
        }
		
        mid := (start + end) / 2 //0  5
		fmt.Println(start,end,mid)
		//0 2 = 0 1 2 2  0 0 1 1 2 2      3 5 = 3 4 5 5 = 3 3 4 4 5 5
        lcpLeft, lcpRight := lcp(start, mid), lcp(mid + 1, end)
        minLength := min(len(lcpLeft), len(lcpRight))
        for i := 0; i < minLength; i++ {
            if lcpLeft[i] != lcpRight[i] {
                return lcpLeft[:i]
            }
        }
        return lcpLeft[:minLength]
    }
    return lcp(0, len(strs)-1)
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}




func main(){

		c := []byte("(){}[]")
		b := c[0] + c[1]
		//81 248 184

		fmt.Println(c,b)
}
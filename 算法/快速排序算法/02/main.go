package main

import "fmt"

func QuickSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }
    splitdata := arr[0]          //第一个数据
	fmt.Println(splitdata)
    low := make([]int, 0, 0)     //比我小的数据
    hight := make([]int, 0, 0)   //比我大的数据
    mid := make([]int, 0, 0)     //与我一样大的数据
    mid = append(mid, splitdata) //加入一个
	// fmt.Println(mid)
    for i := 1; i < len(arr); i++ {
        if arr[i] < splitdata {
            low = append(low, arr[i])
        } else if arr[i] > splitdata {
            hight = append(hight, arr[i])
			//1.[9,10,30]
			//2.[10,30]
			//3.[30]
		
        } else {
            mid = append(mid, arr[i])
        }
    }
	//[1,9,10,30]
	fmt.Println(hight)
    low, hight = QuickSort(low), QuickSort(hight)
	
    myarr := append(append(low, mid...), hight...)
    return myarr
}

//快读排序算法
func main() {
    arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
    fmt.Println(QuickSort(arr))
	fmt.Println(QuickSort(nil))
}
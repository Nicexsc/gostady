package main

import "fmt"

func main() {
    var arr [6][]int
    myarr := []int{1, 2, 3, 1, 1, 2, 2, 2, 2, 2, 3,4,6}
    for i := 0; i < len(myarr); i++ {
        arr[myarr[i]-1] = append(arr[myarr[i]-1], myarr[i])
    }
    fmt.Println(arr)
    var tmp []int
    for j := 0; j < len(arr); j++ {
        // tmp = append(tmp, arr[j][0])
        if len(arr[j]) > 0{
            tmp = append(tmp, arr[j][0])
        }
    }
    fmt.Println(tmp)
}
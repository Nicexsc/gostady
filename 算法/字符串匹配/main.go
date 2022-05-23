package main

import "fmt"

func main() {
    a := "ababaeabacaaaaaddfdfdfdfdf"
    b := "acaad"
    pos := Kmp(b, a)
    fmt.Println(pos)
}
func Kmp(needle string, str string) int {
    next := getNext(needle)
    fmt.Println(next)
    j := 0  // i,j {0,0} {1,1}{2,0}{3,1}{4,0}{5,1}{6,0}{7,1}{8,0}{9,1}{10,2}
	//  a := "ababaeabacaaaaaddfdfdfdfdf"
	// b := "aca"
    for i := 0; i < len(str); i++ {
        for j > 0 && str[i] != needle[j] {
            j = next[j-1] + 1 //0
        }
        if str[i] == needle[j] {
            j++
        }
        if j == len(needle) {
            return i - j + 1
        }
    }
    return -1
}

func getNext(needle string) []int {
    var next = make([]int, len(needle))
    // fmt.Println(next)
    next[0] = -1
    k := -1//acaa 
    for i := 1; i < len(needle); i++ {
        for k != -1 && needle[k+1] != needle[i] {
            k = next[k]
        }
        if needle[k+1] == needle[i] {
            k++
        }
        next[i] = k
		
    }
    return next
}
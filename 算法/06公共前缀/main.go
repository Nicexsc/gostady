package main

import "fmt"

// func longestCommonPrefix(strs []string) string {
// 	stlength := len(strs)

// 	for i := 0; i < stlength -1 ; i++ {
// 		for j := i+1; j <stlength; j++ {
// 			if len(strs[i]) >  len(strs[j]){
// 				strs[i], strs[j] = strs[j],strs[i]
//  			}
// 		}

// 	}
// 	// fmt.Println(strs[0])
// 	length01 := len(strs[0])

// 	for v := 1; v < stlength; v++ {
// 		var st = ""
// 		// fmt.Println(strs[v])
// 		for i := 0; i < length01; i++ {
// 			if string(strs[0][i]) == string(strs[v][i]){
// 				st = st + string(strs[0][i])
// 				// fmt.Println(st)
// 				// if length01 < len(st){
// 				// 	length01 = len(st)

// 				// }

// 			}else{
// 				break
// 			}

// 		}

// 		if length01 >= len(st){
// 			length01 = len(st)

// 		}

// 	}
// 	if length01 !=  0{
// 		return strs[0][0:length01]

// 	}else{
// 		return "\"\""
// 	}

// }

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		//遍历数组
		for j := 1; j < len(strs); j++ {
			//判断改字符长度，利用长度比较，最小长度
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func main() {

	strs := []string{
		"flower", "flow", "flowight",
	}

	st := longestCommonPrefix(strs)
	fmt.Println(st)

}

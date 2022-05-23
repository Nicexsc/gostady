package main

import "fmt"

var runman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	//初始化计量值
	var sum = 0
	//字符串切片
	char := []byte(s)
	//长度获取
	length := len(char)
	//循环遍历，i 需小于倒数两位索引，因为j = i + 1,防止索引超出范围
	for i := 0; i <= length - 1;i ++{
		j := i + 1
		//判断左边是否大于相邻右边值，是则sum 加上 [i] 索引对应的值
		if i == length-1 || runman[string(char[i])] >= runman[string(char[j])] {
			//防止最后一位遗漏，判断以上满足后j是否等于最后一个索引位置时， 需sum赋值加上[j]值
			// if j != length -1 {
				sum = sum + runman[string(char[i])] 
				// fmt.Println(sum)
			// }else{
			// 	sum = sum + runman[string(char[i])]  + runman[string(char[j])]
			// }
			
		} else {
			//左边小于右边时
			sum = sum + runman[string(char[j])] - runman[string(char[i])]
			//此时i需再自加一次
			i = i + 1
		}

	}
	return sum

}

func main() {
	s := "III"
	// s := "IIII"
	nu := romanToInt(s)
	fmt.Println(nu)

}
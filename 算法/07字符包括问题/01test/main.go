package main

import "fmt"

func isValid(s string) bool {
	length := len(s)
	// fmt.Println(s[0])
	
	if length%2 == 1 {
		return false
	}
	stack := []byte{}
	for i := 0; i < length; i++ {
		if s[i] != 40 && s[i] != 123 && s[i] != 91 {
			//判断首个方向。判断相邻是否对应
			//({(()())[]})
			if len(stack) == 0 || s[i] - stack[len(stack)-1] != 1 && s[i] - stack[len(stack)-1] != 2  {
				return false
			}
			//消去匹配对应的
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
		// fmt.Println(stack)

	}
	// fmt.Printf("%#v",stack)
	return len(stack) == 0

}

func main() {

	s := "({(()())[]})"
	t := isValid(s)
	fmt.Println(t)

}
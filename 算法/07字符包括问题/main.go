package main
import "fmt"
func isValid(s string) bool {
    n := len(s)
    if n % 2 == 1 {
        return false
    }
    pairs := map[byte]byte{
        ')': '(',
        ']': '[',
        '}': '{',
    }
    stack := []byte{}
    for i := 0; i < n; i++ {
		//判断括号方向'（'为等于于0 ， ')'为阿斯克编码40 大于零
        if pairs[s[i]] > 0 {
			//判断首个方向。判断相邻是否对应
			//({(()())[]})
            if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
                return false
            }
			//消去匹配对应的
            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, s[i])
        }
    }
    return len(stack) == 0
}


func main() {

	test := map[byte]byte{
        ')': '(',
        ']': '[',
        '}': '{',
    }
	// c := []byte("(){}[]")
	// b := c[0] + c[1]
	//81 248 184
	fmt.Println(test['('])

}


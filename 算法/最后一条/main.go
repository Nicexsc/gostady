package main

import "fmt"

func lengthOfLastWord(s string) int {
	n := len(s)
	for i := n -1; i >= 0; i-- {	
		fmt.Println(s[i])	
	}
	

	return n
}

func main() {

	s := "d   f"
	fmt.Println(string(s[3]))
	// lengthOfLastWord(s)

}
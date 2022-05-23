package main

import "fmt"

/**
* Definition for singly-linked list.
type ListNode struct {
   Val int
  Next *ListNode
}
*/

 type ListNode struct {
    Val int
   Next *ListNode
 }

 //[1 ,3 ,7] [1, 2 , 5, 6] 
 func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    prehead := &ListNode{}
    result := prehead
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            prehead.Next = l1
            l1 = l1.Next
        }else{
            prehead.Next = l2
            l2 = l2.Next
        }
        prehead = prehead.Next
    }
    if l1 != nil {
        prehead.Next = l1
    }
    if l2 != nil {
        prehead.Next = l2
    }
    return result.Next
}





func main(){
	a := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{},
			},
		},
	}
	b := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{},
			},
		},
	}
	
	k := mergeTwoLists(&a, &b)
	fmt.Println(k)
	// for curl !=  nil{
	// 	tmp := curl.Next
	// 	curl.Next = pre
	// 	pre =curl
	// 	curl = tmp 
	// }

}
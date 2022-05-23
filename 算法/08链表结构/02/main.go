package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 type ListNode struct {
    Val int
   Next *ListNode
 }

 type List struct {
    Val int
//    Next *ListNode
 }


 func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var qer *ListNode
	qer = &ListNode{}
	
	var a = qer
	for l1 != nil && l2 != nil{
		if l1.Val >= l2.Val{
			qer.Next = l2
			l2 = l2.Next
			
			



		}else {
			qer.Next = l1
			l1 = l1.Next
		}
		qer = qer.Next
	}
	if l1 != nil{
		qer.Next = l1
	}
	if l2 != nil {
		qer.Next = l2
	}
	return  a.Next

 }


func main(){
	// a := &ListNode{
	// 	Val: 1,
	// 	Next: &ListNode{
	// 		Val: 2,
	// 		Next: &ListNode{
	// 			Val: 4,
	// 			Next: &ListNode{},
	// 		},
	// 	},
	// }
	// b := &ListNode{
	// 	Val: 1,
	// 	Next: &ListNode{
	// 		Val: 3,
	// 		Next: &ListNode{
	// 			Val: 4,
	// 			Next: &ListNode{},
	// 		},
	// 	},
	// }
	
	u := &List{
		Val: 1,
	}
	p := &List{
		Val: 2,
	}
	fmt.Println(u,p,u)
	// k := mergeTwoLists(a, b)
	// fmt.Println(k,a,b)

}
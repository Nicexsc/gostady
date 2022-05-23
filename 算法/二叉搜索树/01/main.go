package main

type Node struct{
	value int
	left *Node
	right *Node
	parent *Node
}

type Tree struct{
	root *Node
	length int
}



func main(){
	

}


func creatTree(){
	arrList := []int {14,2,5,7,23,35,12,17,31}
	myTree := Tree{}
	for i := 0 ; i <len(arrList); i++{
		myTree = insertNode(myTree, arrList[i])
	}

}


func insertNode(tree Tree, insertValue int)Tree{
	var currentNode *Node
	var tmp *Node
	i := 0
	if tree.length == 0{
		currentNode = new(Node)
		currentNode.value = insertValue
		tree.root = currentNode
		return tree
	}else{
		currentNode = tree.root

	}
	for {
		if currentNode.value < insertValue{
			if currentNode.right == nil{
				tmp = new(Node)
				tmp.value = insertValue
				currentNode.right = tmp
				tmp.parent = currentNode
				break
			}else {
				currentNode = currentNode.right
				continue
			}
		}else{
			if cu
		}

	}


}
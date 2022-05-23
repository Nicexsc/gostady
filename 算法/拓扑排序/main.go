package main

import (
    "fmt"
)

//有向图
type graph struct {
    vertex int           //顶点
    list   map[int][]int //连接表边
}

//添加边
func (g *graph) addVertex(t int, s int) {
	//将结构体中list [t]本为空，推入s数据
    g.list[t] = push(g.list[t], s)
}

func main() {
    g := NewGraph(8)
    g.addVertex(2, 1)
    g.addVertex(3, 1)
    g.addVertex(7, 1)
    g.addVertex(4, 2)
    g.addVertex(5, 2)
    g.addVertex(8, 7)
	fmt.Println(g.list)
	//g.list = map[0:[] 1:[] 2:[1] 3:[1] 4:[2] 5:[2] 6:[] 7:[1] 8:[7]]
    g.DfsSort()
}

//创建图
func NewGraph(v int) *graph {
    g := new(graph)
    g.vertex = v  //8
    g.list = map[int][]int{}
    i := 0
    for i < v {
        g.list[i] = make([]int, 0)
        i++
    }
    return g  //{ 0: [],1:[],...7:[],}
}

//取出切片第一个
func pop(list []int) (int, []int) {
    if len(list) > 0 {
        a := list[0]
        b := list[1:]
        return a, b
    } else {
        return -1, list
    }
}

//推入切片

//g.list[2] = [] s =1
func push(list []int, value int) []int {
    result := append(list, value)
    return result
}

//添加边
func (g *graph) KhanSort() {
    var inDegree = make(map[int]int)
    var queue []int
    for i := 1; i <= g.vertex; i++ {
        for _, m := range g.list[i] {
            inDegree[m]++
        }
    }
    for i := 1; i <= g.vertex; i++ {

        if inDegree[i] == 0 {
            queue = push(queue, i)
        }
    }
    for len(queue) > 0 {
        var now int
        now, queue = pop(queue)
        fmt.Println("->", now)
        for _, k := range g.list[now] {
            inDegree[k]--
            if inDegree[k] == 0 {
                queue = push(queue, k)
            }
        }
    }
}

func (g *graph) DfsSort() {
    inverseList := make(map[int][]int)
	test := make(map[int][]int)
	fmt.Println(g.list)
    //初始化逆向邻接表 g.vertex = 8 类似基数排序
    for i := 1; i <= g.vertex; i++ {
		//循环取出有数据键值对
        for _, k := range g.list[i] {
			fmt.Println(k)
            inverseList[k] = append(inverseList[k], i)
			test[k] = append(test[k], g.list[i]...)
        }
    }
	//取出相同值的索引归类为数组
	fmt.Println(inverseList)
	fmt.Println(test)
    visited := make([]bool, g.vertex+1) //make([]bool,9)
    visited[0] = true
	fmt.Println(visited[1])
    for i := 1; i <= g.vertex; i++ {
        if visited[i] == false {
            visited[i] = true
            dfs(i, inverseList, visited)
        }
    }

}
// i =1 , invereseList = map[1:[2 3 7] 2:[4 5] 7:[8]] visited = true
func dfs(vertex int, inverseList map[int][]int, visited []bool) {
	
    for _, w := range inverseList[vertex] {
        if visited[w] == true {
            continue
        } else {
			//[2,3,7] [4,5] [8]
            visited[w] = true
			
            dfs(w, inverseList, visited)
        }
		
    }
    fmt.Println("->", vertex)
}
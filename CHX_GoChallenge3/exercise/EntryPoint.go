package main

import (
	"fmt"
)

func main() {
	edges := [][]int32{
		{1, 8}, {3, 1}, {5, 2}, {10, 1},
	}
	root := findRootUndirection(edges, 1)

	printRoot(root)
}

type node struct {
	x      int32
	childs map[int32]*node
}

// Find Root, but Order is matter
func findRootByOrder(edges [][]int32, s int32) *node {

	var parents map[int32]*node = map[int32]*node{}

	for i := len(edges) - 1; i >= 0; i-- {
		parent := edges[i][0]
		child := edges[i][1]

		if v, ok := parents[parent]; ok {
			parents[parent] = v
		} else {
			parents[parent] = new(node)
			parents[parent].x = parent
			parents[parent].childs = map[int32]*node{}
		}

		parents[parent].childs[child] = new(node)
		parents[parent].childs[child].x = child
		if v, ok := parents[child]; ok {
			parents[parent].childs[child].childs = v.childs
		}
	}

	return parents[s]
}

// Find Root, but Order doesn't matter
func findRoot(edges [][]int32, s int32) *node {
	var parents map[int32]*node = map[int32]*node{}
	for _, edge := range edges {
		parent := edge[0]
		child := edge[1]

		if v, ok := parents[parent]; ok {
			parents[parent] = v
		} else {
			parents[parent] = new(node)
			parents[parent].x = parent
			parents[parent].childs = map[int32]*node{}
		}

		parents[parent].childs[child] = new(node)
		parents[parent].childs[child].x = child
	}

	for _, edge := range edges {
		parent := edge[0]
		child := edge[1]

		if v, ok := parents[child]; ok {
			parents[parent].childs[child].childs = v.childs
		}
	}

	return parents[s]
}

// Find Root, Undirection Graph
func findRootUndirection(edges [][]int32, s int32) *node {
	var parents map[int32]*node = map[int32]*node{}
	for _, edge := range edges {
		parent := edge[0]
		child := edge[1]

		if v, ok := parents[parent]; ok {
			parents[parent] = v
		} else {
			parents[parent] = new(node)
			parents[parent].x = parent
			parents[parent].childs = map[int32]*node{}
		}

		parents[parent].childs[child] = new(node)
		parents[parent].childs[child].x = child
	}

	for _, edge := range edges {
		parent := edge[0]
		child := edge[1]

		if v, ok := parents[child]; ok {
			parents[parent].childs[child].childs = v.childs
		}
	}

	for _, v := range parents {
		fmt.Println(*v)
	}

	return parents[s]
}

func printRoot(root *node) {
	print(root, 0)
}

func print(root *node, level int) {
	if root == nil {
	} else {
		str := ""
		for i := 0; i < level; i++ {
			if i == 0 {
				str += "|"
			}
			str += "-"
		}
		fmt.Println(str, root.x)
		for _, v := range root.childs {
			print(v, level+1)
		}
		// if v := root.parent; v != nil {
		// 	fmt.Println(str, root.parent.x)
		// }
	}
}

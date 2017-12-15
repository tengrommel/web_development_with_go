package main

import "fmt"

type Node struct {
	Value int
	Left *Node
	Right *Node
}

func printNode(n *Node)  {
	fmt.Print("Value: ", n.Value)
	if n.Left != nil{
		fmt.Println("Left:", n.Left.Value)
	}
	if n.Right!= nil{
		fmt.Println("Right:", n.Right.Value)
	}
	fmt.Println()
}

func read() []Node {
	var N int
	fmt.Scanf("%d", &N)
	fmt.Println(N)
	var nodes []Node = make([]Node, N)
	for i:=0;i<N;i++{
		var val, indexLeft, indexRight int
		fmt.Scanf("%d %d %d", &val, &indexLeft, &indexRight)
		nodes[i].Value = val
		if indexLeft >=0 {
			nodes[i].Left = &nodes[indexLeft]
		}
		if indexRight >=0 {
			nodes[i].Right = &nodes[indexRight]
		}
		for _, node := range nodes{
			printNode(&node)
		}
	}
	return nodes
}

func main() {
	nodes := read()
	for _, node := range nodes{
		printNode(&node)
	}
}

package main

import "fmt"

// 循环链表是尾结点的指针域指向第一个节点的地址(**注意不是头结点**)
// ** 所以第一个节点会被两个指针域指向(头结点, 尾结点) **
// 第一个位置插入新节点时, 注意要将尾部的节点指向新节点; 删除第一个节点也需要调整尾部节点
type CircularLinkNode struct {
	Data interface{}
	Next *CircularLinkNode
}

// create
func (node *CircularLinkNode) Create(data ...interface{}) {
	if node == nil || len(data) == 0 {
		return
	}

	head := node // 记录头结点
	for _, v := range data {
		tempNode := new(CircularLinkNode)
		tempNode.Data = v

		node.Next = tempNode
		node = tempNode
	}
	// 尾结点指向第一个节点
	node.Next = head.Next
}

// print
func (node *CircularLinkNode) Print() {
	// 判断头结点是否为空
	if node == nil {
		fmt.Println("")
	}

	// 记录第一个节点
	start := node.Next
	for start != nil {
		node = node.Next
		if node.Data != nil {
			fmt.Print(node.Data, " ")
		}
		// 判断node是否到末尾了
		if node.Next == start {
			break
		}
	}
	fmt.Println()
}

// length
func (node *CircularLinkNode) Length() int {
	if node == nil {
		return -1
	}

	start := node.Next
	var count int
	for start != nil {
		node = node.Next
		count++
		if node.Next == start {
			break
		}
	}
	return count
}

// insert ** 注意处理在第一个节点插入问题(需更改尾结点指向) **
func (node *CircularLinkNode)InsertByIndex(index int, data interface{}) {
	if node == nil || index <= 0 || data == nil {
		return
	}
	// 限制index的范围
	if index > node.Length() + 1 {
		return
	}
	// 没有第一个节点
	if node.Next == nil {
		return
	}

	// 此处记录初始第一个节点
	start := node.Next
	// 将节点移到前一个位置, 创建新节点, 调整指针域指向
	// 将当前节点移到index-1位置, i从1开始
	for i := 1; i < index; i++ {
		node = node.Next
	}
	tempNode := new(CircularLinkNode)
	tempNode.Data = data
	tempNode.Next = node.Next
	node.Next = tempNode

	// 特殊处理第一个位置插入节点的尾结点指向问题
	if index == 1 {
		// 将当前节点移动到尾结点
		node = start
		// 注意: node初始化为start, 是为了让链表从原来的起始位置开始寻找,
		// 而不是从新创建的节点开始找(原来的首节点在新建节点之后, 循环没到尾部就停止了)
		for node.Next != start { // 循环完后, node就是尾结点
			node = node.Next
		}
		// 尾结点指向新创建的节点
		node.Next = tempNode
	}
}

// delete ** 同样要注意删除第一个节点时, 尾结点指向问题 **
func (node *CircularLinkNode) DeleteByIndex(index int) {
	if node == nil || index <= 0 {
		return
	}
	if index > node.Length() {
		return
	}
	if node.Next == nil {
		return
	}

	// 记录初始第一个节点, 便于后面找尾结点和调整尾结点指向
	start := node.Next
	head := node // 方便尾结点指向, 方便找到新的第一个节点
	// 移动到index前一个节点
	for i := 1; i < index; i++ {
		node = node.Next
	}
	node.Next = node.Next.Next

	// 特殊处理第一个节点删除, 尾结点指向问题
	if index == 1 {
		// 这里node不用初始化, 因为原来的指向已经被断了, 不存在提前退出循环
		for node.Next != start {
			node = node.Next
		}
		// 将尾结点指向新的第一个节点
		node.Next = head.Next
		// 释放原起始节点
		start.Data = nil
		start.Next = nil
	}
}

// destroy

func OutCall_c() {
	cl := new(CircularLinkNode)
	cl.Create(1, 2, 3)
	cl.Print()
	count := cl.Length()
	fmt.Println("长度:", count)
	cl.InsertByIndex(4, 4)
	cl.Print()
	cl.DeleteByIndex(5)
	cl.Print()
}
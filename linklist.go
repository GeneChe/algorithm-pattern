package main

import "fmt"

// 节点
// 通过结构体嵌套本结构体指针来实现链表
type LinkNode struct {
	Data interface{}
	// Next LinkNode  //结构体中不能嵌套本结构体, invalid recursive type LinkNode
	Next *LinkNode
}

// create
func (node *LinkNode) Create(data ...interface{}) {
	if node == nil || len(data) == 0 {
		return
	}

	// 创建节点存储数据
	for _, v := range data {
		tempNode := new(LinkNode)
		tempNode.Data = v

		// 将新节点地址赋值给当前节点的next
		node.Next = tempNode
		// 将当前节点移到新创建的节点
		node = tempNode
	}
}

// print
func (node *LinkNode) Print() {
	if node == nil {
		return
	}

	// 1.循环方式遍历链表
	for node != nil {
		if node.Data != nil {
			fmt.Print(node.Data, " ")
		}

		// 移动当前节点
		node = node.Next
	}
	// 2. 递归方式
	//if node.Data != nil {
	//	fmt.Print(node.Data, " ")
	//}
	//node.Next.Print()

	fmt.Println()
}

// length(不计算头结点)
func (node *LinkNode) Length() int {
	if node == nil {
		return -1
	}
	var count int
	// 计算整个链表的非nil指针域个数
	for node.Next != nil {
		count++
		// 移动当前节点
		node = node.Next
	}
	return count
}

// 头插
func (node *LinkNode) InsertByHead(data interface{}) {
	if node == nil || data == nil {
		return
	}
	//创建新节点
	tempNode := new(LinkNode)
	tempNode.Data = data
	tempNode.Next = node.Next // 新节点指向原来的下个节点
	// 将头结点指向新节点
	node.Next = tempNode
}

// 尾插
func (node *LinkNode) InsertByTail(data interface{}) {
	if node == nil || data == nil {
		return
	}

	// 移动当前节点到尾结点
	for node.Next != nil { // node.next == nil 就是尾结点了
		node = node.Next
	}

	tempNode := new(LinkNode)
	tempNode.Data = data
	// 原尾结点指向新尾结点
	node.Next = tempNode
}

// 根据位置插入(从0开始)
func (node *LinkNode) InsertByIndex(index int, data interface{}) {
	if node == nil || data == nil {
		return
	}

	// 移动节点到index的前一个位置
	for i := 0; i < index; i++ {
		// 如果index超过了链表的长度即此时已经移到链表的末尾了, 直接返回
		// 这里判断node.next而不是判断node, next为空说明没有下个节点了, 无需再移动
		if node.Next == nil {
			return
		}
		// 移动节点
		node = node.Next
	}
	// 创建节点链接节点
	tempNode := new(LinkNode)
	tempNode.Data = data

	tempNode.Next = node.Next
	node.Next = tempNode
}

func OutCall_l() {
	l := new(LinkNode) // l就是头结点
	// 创建链表
	l.Create(1, 2, 3)
	l.Print()
	// 计算链表长度
	count := l.Length()
	fmt.Println("长度:", count)
	// 头插
	l.InsertByHead(0)
	l.Print()
	// 尾插
	l.InsertByTail(4)
	l.Print()
	// 插入(位置)
	l.InsertByIndex(5, 5)
	l.Print()
}
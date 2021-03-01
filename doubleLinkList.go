package main

import "fmt"

type DoubleLinkNode struct {
	Data interface{}
	Prev *DoubleLinkNode // 前指针
	Next *DoubleLinkNode // 后指针
}

// create
func (node *DoubleLinkNode) Create(data ...interface{}) {
	if node == nil || len(data) == 0 {
		return
	}

	for _, v := range data {
		// 创建新节点
		tempNode := new(DoubleLinkNode)
		tempNode.Data = v
		// 调整指针域指向
		tempNode.Prev = node
		node.Next = tempNode
		// 移动当前节点到新节点
		node = tempNode
	}
}

// print (正序打印)
func (node *DoubleLinkNode) Print() {
	if node == nil {
		fmt.Println("")
	}

	for node != nil {
		if node.Data != nil {
			fmt.Print(node.Data, " ")
		}
		node = node.Next
	}
	fmt.Println()
}

// print (倒序打印)
func (node *DoubleLinkNode) ReversePrint() {
	if node == nil {
		return
	}

	// 移动节点到末尾
	for node.Next != nil { // 循环结束后, node.next == nil, 即为末尾节点
		// 节点移动
		node = node.Next
	}
	// 从后往前移动节点
	for node.Prev != nil {
		if node.Data != nil {
			fmt.Print(node.Data, " ")
		}
		node = node.Prev
	}
	fmt.Println()
}

// length
func (node *DoubleLinkNode) Length() int {
	if node == nil {
		return -1
	}
	var count int
	for node.Next != nil {
		count++
		node = node.Next
	}
	return count
}

// insert(index, data)
func (node *DoubleLinkNode) InsertByIndex(index int, data interface{}) {
	if node == nil || data == nil || index < 0 {
		return
	}

	// 链表头结点算索引0, 其他从1开始
	// 将节点移动到index位置, 这里循环从[0, index)就是移动index次数
	for i := 0; i < index; i++ {
		// 索引超过链表长度
		if node.Next == nil {
			return
		}
		node = node.Next
	}

	// 链接新节点
	tempNode := new(DoubleLinkNode)
	tempNode.Data = data
	tempNode.Prev = node.Prev
	tempNode.Next = node

	// 注意先赋值node.Prev.next然后再赋值node.prev
	node.Prev.Next = tempNode
	node.Prev = tempNode
}

// delete(index) 这index > 0
func (node *DoubleLinkNode) DeleteByIndex(index int) {
	if node == nil || index < 0 {
		return
	}
	// 将当前节点移动到index处
	for i := 0; i < index; i++ {
		if node.Next == nil {
			return
		}
		node = node.Next
	}
	// 链接节点
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	// 销毁node
	node.Next = nil
	node.Prev = nil
	node.Data = nil
}

// destroy
func (node *DoubleLinkNode) Destroy() {
	if node == nil {
		return
	}
	node.Next.Destroy()
	node.Data = nil
	node.Prev = nil
	node.Next = nil
}

func OutCall_d() {
	// 头结点算索引0
	dl := new(DoubleLinkNode)
	dl.Create(1, 2, 3)
	fmt.Println(dl)
	// 打印
	dl.Print()
	//dl.ReversePrint()
	// 长度
	//count := dl.Length()
	//fmt.Println("length:", count)
	dl.InsertByIndex(2, 666)
	dl.Print()
	dl.ReversePrint()
	// 删除
	//dl.DeleteByIndex(2)
	//dl.Print()
	//dl.ReversePrint()
	dl.Destroy()
	fmt.Println("---双向链表销毁---")
	dl.Print()
	dl.ReversePrint()
}
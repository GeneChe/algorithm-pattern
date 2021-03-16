package base

import (
	"fmt"
	"reflect"
)

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
	//if node.Data != nil { // 正序打印
	//	fmt.Print(node.Data, " ")
	//}
	//node.Next.Print()
	//if node.Data != nil { // 倒序打印
	//	fmt.Print(node.Data, " ")
	//}

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
	if node == nil || data == nil || index < 0 {
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

// delete 位置
func (node *LinkNode) DeleteByIndex(index int) {
	if node == nil || index < 0 {
		return
	}
	// 移动节点到index前一个位置
	for i := 0; i < index; i++ {
		if node.Next == nil {
			return
		}
		node = node.Next
	}

	// 将前一个节点指针域指向后一个节点指针域指向的地址
	node.Next = node.Next.Next
}

// delete 数据
func (node *LinkNode) DeleteByData(data interface{}) {
	if node == nil || data == nil {
		return
	}

	// 遍历节点, 比较数据
	var preNode *LinkNode
	for node.Next != nil {
		preNode = node
		node = node.Next
		// 比较interface{}中的数据值和数据类型
		if reflect.TypeOf(node.Data) == reflect.TypeOf(data) && node.Data == data {
			preNode.Next = node.Next
			// 置空node
			node.Data = nil
			node.Next = nil
			// 有return是删除第一个相同的元素, 没有就是删除所有的元素
			return
		}
	}
}

// search (位置不计算头结点)
func (node *LinkNode) Search(data interface{}) int {
	if node == nil || data == nil {
		return -1
	}

	// 遍历节点
	var count int
	for node.Next != nil {
		node = node.Next
		// 比较两个接口类型的值是否相同
		if reflect.DeepEqual(node.Data, data) {
			return count
		}
		count++
	}

	return -1
}

// destroy
func (node *LinkNode) Destroy() {
	if node == nil {
		return
	}
	// 使用递归方式销毁链表
	node.Next.Destroy()
	node.Data = nil
	node.Next = nil
}

func OutCall_l() {
	l := new(LinkNode) // l就是头结点
	// 创建链表
	l.Create(1, 2, 3, 4, 5)
	l.Print()
	// 计算链表长度
	//count := l.Length()
	//fmt.Println("长度:", count)
	// 头插
	//l.InsertByHead(0)
	//l.Print()
	// 尾插
	//l.InsertByTail(4)
	//l.Print()
	// 插入(位置)
	//l.InsertByIndex(5, 5)
	//l.Print()
	// 删除(位置)
	//l.DeleteByIndex(2)
	//l.Print()
	// 删除(数据)
	//l.DeleteByData(2)
	//l.Print()
	index := l.Search(3)
	fmt.Println("index:", index)
	l.Destroy()
	fmt.Println("链表销毁")
	l.Print()
}
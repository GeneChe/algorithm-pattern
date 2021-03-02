package main

import "fmt"

// 链栈存储结构 --- 插入删除都在头部
type StackNode struct {
	Data interface{}
	Next *StackNode
}

// create
// 相当于头插法创建链表, 栈顶指针即为头结点(存储数据)
func CreateStack(data ...interface{}) *StackNode {
	if data == nil {
		return nil
	}
	// 栈顶指针
	var top *StackNode
	for _, v := range data {
		tempNode := new(StackNode)
		tempNode.Data = v
		// 如果栈顶指针有值, 新节点指向原来的栈顶; 没有就将新节点当成栈顶指针
		if top != nil {
			tempNode.Next = top
		}
		// 再移动栈顶指针到新节点
		top = tempNode
	}
	return top
}

// print
func PrintStack(s *StackNode) {
	if s == nil {
		return
	}
	for s != nil {
		fmt.Print(s.Data, " ")
		s = s.Next
	}
	fmt.Println()
}

// length
func LengthStack(s *StackNode) int {
	if s == nil {
		return -1
	}
	var count int
	for s != nil {
		count++
		s = s.Next
	}
	return count
}

// push 头插
func PushStack(top **StackNode, data interface{}) {
	if top == nil || *top == nil {
		return
	}
	if data == nil {
		return
	}

	tempNode := new(StackNode)
	tempNode.Data = data
	// 新节点指向原先栈顶
	tempNode.Next = *top
	*top = tempNode
}

// pop
func PopStack(top **StackNode) (data interface{}) {
	if top == nil || *top == nil {
		return nil
	}
	// 将原先栈顶的值返回, 调整栈顶指针
	data = (*top).Data
	*top = (*top).Next
	return
}

// clear
func ClearStack(top **StackNode) {
	if top == nil {
		return
	}
	// 直接释放栈顶指针, 整个栈就会被回收
	*top = nil
}

func OutCall_s() {
	// 创建
	stack := CreateStack(1, 2, 3, 4, 5) // stack代表栈顶指针
	fmt.Println(stack)
	// 打印
	PrintStack(stack)
	// 长度
	len := LengthStack(stack)
	fmt.Println("len:", len)
	// 入栈
	PushStack(&stack, 6)
	PrintStack(stack)
	PushStack(&stack, 7)
	PrintStack(stack)
	// 出栈
	data := PopStack(&stack)
	fmt.Println("pop value:", data)
	data = PopStack(&stack)
	fmt.Println("pop value:", data)
	// 清空栈
	fmt.Println("清空...")
	ClearStack(&stack)
	PrintStack(stack)
}

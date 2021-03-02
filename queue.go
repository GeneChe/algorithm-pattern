package main

import "fmt"

// 链表实现队列 --- 尾插头出
type QueueNode struct {
	Data interface{}
	Next *QueueNode
}

// create -- 从尾部插入
func CreateQueue(data ...interface{}) *QueueNode {
	if len(data) == 0 {
		return nil
	}

	// 队头, 队尾
	var front, rear *QueueNode
	for _, v := range data {
		tempNode := new(QueueNode)
		tempNode.Data = v
		if front == nil {
			front = tempNode
			rear = tempNode
		} else {
			rear.Next = tempNode
			// 更新队尾
			rear = tempNode
		}
	}
	// 返回队头
	return front
}

// print
func PrintQueue(queue *QueueNode) {
	if queue == nil {
		return
	}
	for queue != nil {
		fmt.Print(queue.Data, " ")
		queue = queue.Next
	}
	fmt.Println()
}

// length
func LengthQueue(queue *QueueNode) int {
	if queue == nil {
		return -1
	}
	var count int
	for queue != nil {
		count++
		queue = queue.Next
	}
	return count
}

// push 尾插
func PushQueue(queue *QueueNode, data interface{}) {
	if queue == nil {
		return
	}
	if data == nil {
		return
	}
	// 移动到队尾
	for queue.Next != nil {
		queue = queue.Next
	}

	tempNode := new(QueueNode)
	tempNode.Data = data
	queue.Next = tempNode
}

// pop 头部删除
func PopQueue(frontPtr **QueueNode) (data interface{}) {
	if frontPtr == nil || *frontPtr == nil {
		return nil
	}
	data = (*frontPtr).Data
	*frontPtr = (*frontPtr).Next
	return
}

func OutCall_q() {
	queue := CreateQueue(1, 2, 3, 4)
	fmt.Println(queue)
	PrintQueue(queue)
	count := LengthQueue(queue)
	fmt.Println("count:", count)
	PushQueue(queue, 5)
	PrintQueue(queue)
	PushQueue(queue, 6)
	PrintQueue(queue)
	data := PopQueue(&queue)
	fmt.Println("pop value:", data)
	data =  PopQueue(&queue)
	fmt.Println("pop value:", data)
	PrintQueue(queue)
}
package main

// 链表实现队列 --- 尾插头出
type QueueNode struct {
	Data interface{}
	Next *QueueNode
}

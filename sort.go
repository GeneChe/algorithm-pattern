package main

import "fmt"

// 排序的稳定性: 相同的数据, 原先的前后顺序在排序后没有改变则为稳定的, 改变了则为不稳定的

// 冒泡排序:
// 1. 外层控制行
// 2. 内层控制列
// 3. 比较相邻元素
// 4. 满足条件交换
func BubbleSort(arr []int) {
	var count int
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			count++
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("比较了:", count)
}

// 冒泡排序优化 - 针对有序序列减少比较次数
func BubbleSortAdvance(arr []int) {
	// 记录序列是否有序的标记
	var flag bool
	var count int
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			count++
			if arr[j] > arr[j+1] {
				flag = true
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		// 如果一轮下来没有数据交换则说明数列已经有序无需再循环下去
		if !flag {
			break
		} else {
			flag = false // 注意: 一轮比较完后重置flag
		}
	}
	fmt.Println("比较了:", count)
}

func OutCall_st() {
	arr := []int{1, 3, 5, 2, 8, 6, 4, 9}
	arr1 := []int{1, 2, 3, 4, 5, 8, 6, 9}
	BubbleSort(arr)
	fmt.Println(arr)
	BubbleSortAdvance(arr1)
	fmt.Println(arr1)
}
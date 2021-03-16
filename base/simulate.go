package base

import (
	"fmt"
)

func OutCall_sl() {
	// slice 模拟 栈, 队列
	stack := make([]int, 0)
	// push
	stack = append(stack,1)
	// pop
	_ = stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	// 检查栈空
	//len(stack) == 0

	// queue
	queue := make([]int, 0)
	// enter
	queue = append(queue,2)
	// out
	_ = queue[0]
	queue = queue[1:]

	// sort
	// sort.Ints()
	// sort.Strings()
	// sort.Slice()

	// math.MaxInt32
	// math.MinInt32

	var a = []int {1, 2, 3, 5, 6}
	i := 1
	// 删除a[i]
	copy(a[i:], a[i+1:])
	a = a[:len(a)-1]

	// make创建了长度，则通过索引赋值
	// make长度为0，则通过append()赋值

	// 常见技巧
	s := "12345" // s[0] byte
	n := int(s[0]-'0') // 1 int
	_ = byte(n + '0') // '1'

	// index := strStr("hellogenech", "gene")
	// fmt.Println(index)
	res := subsets([]int{1,2,4})
	fmt.Println(res)
}

// 查找字符串第一次出现的位置: 以当前字符开头字符串是否等于目标字符串
func strStr(src, search string) int {
	var lenSrc, lenSear = len(src), len(search)
	if lenSear > lenSrc {
		return -1
	}

	var i, j int
	for i = 0; i < lenSrc-lenSear+1; i++ {
		for j = 0; j < lenSear; j++ {
			if src[i+j] != search[j] {
				break
			}
		}
		// 循环完整走完
		if lenSear == j {
			return i
		}
	}
	return -1
}

// 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）
// 思路：这是一个典型的应用回溯法的题目，简单来说就是穷尽所有可能性
func subsets(num []int) [][]int {
	// 结果集
	result := make([][]int, 0)
	// 临时集
	var set []int
	// 回溯方法
	var back func(cur int, num []int)
	back = func(cur int, num []int) {
		// 当cur == num的长度时, 枚举完成, 添加临时集合到最终结果中
		if cur == len(num) {
			// 注意匿名函数使用外部局部变量问题
			result = append(result, append([]int(nil), set...))
			return
		}
		// cur之前的序列都已经枚举完了
		// 当前cur位置和之后的还没确定
		// 尝试取当前数
		set = append(set, num[cur])
		// 去确定剩下的数
		back(cur+1, num)
		// 再尝试不取当前的数
		set = set[:len(set)-1]
		// 去确定剩余的数
		back(cur+1, num)
	}
	back(0, num)
	return result
}
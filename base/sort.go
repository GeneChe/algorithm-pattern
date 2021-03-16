package base

import (
	"fmt"
	"math/rand"
	"time"
)

// 排序的稳定性: 相同的数据, 原先的前后顺序在排序后没有改变则为稳定的, 改变了则为不稳定的

// 冒泡排序: -- 稳定排序
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

// 选择排序 -- 不稳定排序
// 1. 在元素中找到最大值的下标
// 2. 跟剩余值中最后一个元素交换位置
func SelectSort(arr []int) {
	var index, count int
	for i := 0; i < len(arr)-1; i++ {
		index = 0 // 每次都用第一个值进行初始化
		for j := 1; j < len(arr)-i; j++ { // 所以j从下标1开始, 因为初始化是0无需在比较了
			count++
			if arr[j] > arr[index] {
				// 记录最大值下标
				index = j
			}
		}
		// 交换最大值到最后位置
		arr[index], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[index]
	}
	fmt.Println("比较了:", count)
}

// 插入排序 -- 稳定排序
// 算法设计的思路是，将数组划分成两部分，第一部分是有序的，第二部分是无序的
// 每次从无序部分取一个元素，将这个元素插入到有序部分，保持有序部分的有序性质
// 直到无序部分为空
func InsertSort(arr []int) {
	var count int
	// 默认第一个元素是有序的, 所以i从1开始
	for i := 1; i < len(arr); i++ {
		count++
		// 要插入的值
		vi := arr[i]
		// 当前位置的左边是有序数列得最右边
		j := i - 1
		// 比较要插入的值和有序序列的值, 小于有序序列, 则有序序列的向右移动一位, 直到找到最后要插入的位置
		for ; j >= 0 && vi < arr[j]; j-- {
			count++
			arr[j+1] = arr[j]
		}
		// 注意: 循环结束后j+1的位置才是要插入的位置(j--)
		if j+1 != i {
			arr[j+1] = vi
		}
	}
	fmt.Println("比较了:", count)
}

// 希尔排序 -- 不稳定排序
// 先取整数d1作为第一个增量，把数组的全部记录分组。所有距离为d1的倍数的记录放在同一个组中。
// 先在各组内进行直接插入排序；然后，取第二个增量d2 ... 直到增量为1成了一个分组
// 优化的地方: 这种比较相隔较远距离（称为增量）的数，使得数移动时能跨过多个元素，则进行一次比较就可能消除多个元素交换
// 增量的取值: 一般的初次取序列的一半为增量，以后每次减半，直到增量为1。
func ShellSort(arr []int) {
	// 控制步长
	for step := len(arr)/2; step > 0; step /= 2 {
		// 从step处开始循环
		for i := step; i < len(arr); i++ {
			// 开始插入排序 从 i-step开始, 初始为0
			for j := i - step; j >= 0 && arr[j+step] < arr[j]; j -= step { // j -= step相当于插入排序的j--, 依次向前比较
				arr[j], arr[j+step] = arr[j+step], arr[j]
			}
		}
	}
}

// 快速排序 -- 不稳定排序
// 快排可以在一次循环中（递归调用），找出某个元素的正确位置，并且该元素之后不需要任何移动
// 方法的基本思想是：
//1．先从数列中取出一个数作为基准数。
//2．分区过程，将比这个数大的数全放到它的右边，小于或等于它的数全放到它的左边。
//3．再对左右区间重复第二步，直到各区间只有一个数
func QuickSort(list []int, start, end int) {
	// 只剩一个元素时就返回了
	if start >= end {
		return
	}

	// 标记最左侧元素作为参考
	tmp := list[start]
	// 两个游标分别从两端相向移动，寻找合适的"支点"
	left := start
	right := end
	for left != right {
		// 右边的游标向左移动，直到找到比参考的元素值小的
		for list[right] >= tmp && left < right { // ** ???这个循环为什么不能跟下面的交换位置??? **
			right--
		}
		// 左侧游标向右移动，直到找到比参考元素值大的
		for list[left] <= tmp && left < right {
			left++
		}

		// 如果找到的两个游标位置不统一，就游标位置元素的值，并继续下一轮寻找
		// 此时交换的左右位置的值，右侧一定不大于左侧。可能相等但也会交换位置，所以才叫不稳定的排序算法
		if left < right {
			list[left], list[right] = list[right], list[left]
		}
	}
	// 这时的left位置已经是我们要找的支点了，交换位置
	list[start], list[left] = list[left], tmp

	// 按支点位置吧原数列分成两段，再各自逐步缩小范围排序
	QuickSort(list, start, left-1)
	QuickSort(list, left+1, end)
}

func QuickSort1(array []int, left, right int) {
	l := left
	r := right
	pivot := array[(left+right)/2]
	//for循环的目标是将比pivot小的数放在左边，比pivot大的数放在右边
	for l < r {
		//从pivot左边找到一个大于等于pivot的值
		for array[l] < pivot {
			l++
		}
		//从pivot右边找到小于等于pivot的值
		for array[r] > pivot {
			r--
		}
		//表明本次分解任务完成
		if l >= r {
			break
		}
		//交换
		array[l], array[r] = array[r], array[l]
		//优化
		if array[l] == pivot { // 因为上面交换了l和r位置的值, 所以这里是r--
			r--
		}
		if array[r] == pivot { // 同上
			l++
		}
	}
	//如果l==r，在移动下
	if l == r {
		l++
		r--
	}
	//向左递归
	if left < r {
		QuickSort1(array, left, r)
	}
	//向右递归
	if right > l {
		QuickSort1(array, l, right)
	}
}

// 原地交换，所以传入交换索引
func QuickSort2(nums []int, start, end int) {
	if start < end {
		// 分治法：divide
		pivot := partition(nums, start, end)
		QuickSort2(nums, 0, pivot-1)
		QuickSort2(nums, pivot+1, end)
	}
}
// 分区
func partition(nums []int, start, end int) int {
	// 选取最后一个元素作为基准pivot
	p := nums[end]
	i := start
	// 最后一个值就是基准所以不用比较
	for j := start; j < end; j++ {
		if nums[j] < p {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	// 把基准值换到中间
	nums[i], nums[end] = nums[end], nums[i]
	return i
}

// 堆排序 -- 不稳定排序
// 初次构建大根堆时有较多次的排序，所以不适合对少量元素进行排序
// 最大堆是一个完全二叉树。 并且要求每个结点的值必须大于他的两个子节点。 所以他的根结点一定是最大值。
// 但是左右结点大小不一定。
// 用数组表示的二叉树，可以这样表达： i的子节点下标为 2*i + 1 和 2 * i + 2.   i的父节点下标为 (i-1)/2。
func HeapSort(arr []int) {
	length := len(arr)
	// 由无序数组构成大根堆
	for i := length/2-1; i >= 0; i-- {
		sink(arr, i, length-1)
	}

	// 此时tree已经是个大根堆了。只需每次交换根节点和最后一个节点，并减少一个比较范围。再进行一轮比较
	// 根节点存储最大值, 将根节点跟叶子节点数据交换
	// 即交换a[0]跟a[length-1]
	// 然后把前段数组继续下沉保持数据结构, 如此循环
	for i := length-1; i >= 1; i-- {
		// 如果只剩根节点和左孩子节点，就可以提前结束了
		if i == 1 && arr[0] <= arr[i] {
			break
		}
		// 交换根节点和比较范围内最后一个节点的数值
		arr[0], arr[i] = arr[i], arr[0]
		// 这里递归的把较大值一层层提上来
		sink(arr, 0, i-1)
	}
}

func sink(arr []int, startNode, latestNode int) {
	var largerChild int
	leftChild := startNode*2 + 1
	rightChild := leftChild + 1

	// 子节点超过比较范围就跳出递归
	if leftChild >= latestNode {
		return
	}

	// 左右孩子节点中找到较大的，右孩子不能超出比较的范围
	if rightChild <= latestNode && arr[rightChild] > arr[leftChild] {
		largerChild = rightChild
	} else {
		largerChild = leftChild
	}

	// 此时startNode节点数值已经最大了，就不用再比下去了
	if arr[largerChild] <= arr[startNode] {
		return
	}

	// 到这里发现孩子节点数值比父节点大，所以交换位置，并继续比较子孙节点，直到把大鱼捞上来
	arr[startNode], arr[largerChild] = arr[largerChild], arr[startNode]
	sink(arr, largerChild, latestNode)
}

// 二分查找
// 注意: 前提是有序数据才能用二分查找
func BinarySearch(arr []int, num int) (bool, int) {
	start := 0
	end := len(arr)-1
	mid := (start+end)/2 // 基准数
	for i := 0; i < len(arr); i++ {
		// 如果num == arr[mid]则mid位置的树就是要找的值
		if num == arr[mid] {
			return true, mid
		} else if num > arr[mid] { // 说明num在mid的右边
			start = mid+1
		} else {
			end = mid-1
		}
		// 重新计算中间位置
		mid = (start+end)/2
	}
	return false, -1
}

// 变相排序 基于大量重复 值在某一范围
// 如高考分数等情况
func sample() {
	t := time.Now().Unix()
	fmt.Println("时间戳:", t)
	rand.Seed(t)
	s := make([]int, 10000)
	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(1000) // 0~999
	}
	fmt.Println(s)

	// 记录s中每个数出现的次数
	count := make(map[int]int)
	for _, v := range s {
		count[v]++
	}
	// 打印排序后的数据
	for i := 0; i < 1000; i++ { // s中值的范围
		for j := 0; j < count[i]; j++ { // s中值得次数
			fmt.Print(i, " ")
		}
	}
}

func OutCall_st() {
	arr := []int{1, 13, 5, 12, 8, 6, 4, 9, 11, 3, 7, 2, 3, 14, 10, 1}
	//arr1 := []int{1, 2, 3, 4, 5, 8, 6, 9}
	//BubbleSort(arr)
	//BubbleSortAdvance(arr1)
	//SelectSort(arr)
	//InsertSort(arr)
	//ShellSort(arr)
	QuickSort(arr, 0, len(arr)-1)
	//HeapSort(arr)
	fmt.Println(arr)
	//sample()

	// search
	num := 3
	res, idx := BinarySearch(arr, num)
	fmt.Println("找到", num, "结果:", res, "位置:", idx)
}
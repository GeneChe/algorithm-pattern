package base

import "fmt"

// 循环链表是尾结点的指针域指向第一个节点的地址(**注意不是头结点**)
// ** 所以第一个节点会被两个指针域指向(头结点, 尾结点) **
// 第一个位置插入删除新节点时, 注意需要调整尾结点和头结点的指向
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

// destroy 从头开始循环销毁每个节点, 最后再销毁头结点的指针域指向
func (node *CircularLinkNode) Destroy() {
	if node == nil {
		return
	}

	// 记录起始节点
	start := node.Next
	var tempNode *CircularLinkNode
	for start != nil {
		// 获取下一个节点
		tempNode = start.Next
		// 释放上一个节点
		start.Data = nil
		start.Next = nil

		start = tempNode
	}
	// 将头结点指向取出
	node.Next = nil
}

// xxxxx 未解决 xxxxx
// 约瑟夫环
// 类似问题1：n个人（编号0~(n-1）），从0开始报数，报到（m-1）的退出，剩下的人继续从0开始报数。求胜利者的编号。
// 类似问题2: n个人排成一圈。从某个人开始，按顺时针方向依次编号。从编号为1的人开始顺时针“一二一”报数，报到2的人退出圈子。
// 	这样不断循环下去，圈子里的人将不断减少。由于人的个数是有限的，因此最终会剩下一个人。试问最后剩下的人最开始的编号。
// 类似问题3: 一堆猴子都有编号，编号是1，2，3 ...m，这群猴子（m个）按照1-m的顺序围坐一圈，
// 	从第1开始数，每数到第N个，该猴子就要离开此圈，这样依次下来，直到圈中只剩下最后一只猴子，则该猴子为大王。

/*
	以总数8, 数3为例  m, n
	第一次 count = 3 删除一个后 长度 = 7
	第二次 count = 5 删除一个后 长度 = 6
	推到出 count = 上次count + 2, 得到此时count = 7 大于长度6.
	将推到后的count 和 长度 取余数得到新一轮的count值为1
	总结: count初始值为1, 每次循环增量2 (n - 1). 计算出的count值大于length时, 取余数作为新的count值. 如果余数为0, 则将count置为length
	注意: 长度是不断变化了, 所以count 不是 3 6 9 ...
*/

/*
Josephu问题为：设编号为1，2，…n的n个人围坐一圈，约定编号为k（1<=k<=n）的人从1开始报数，数到m 的那个人出列，
	它的下一位又从1开始报数，数到m的那个人又出列，依次类推，直到所有人出列为止，由此产生一个出队编号的序列
//小孩的结构体
type Boy struct {
	No int // 编号
	Next *Boy // 指向下一个小孩的指针[默认值是nil]
}

// 编写一个函数，构成单向的环形链表
// num ：表示小孩的个数
// *Boy : 返回该环形的链表的第一个小孩的指针
func AddBoy(num int) *Boy {

	first := &Boy{} //空结点
	curBoy := &Boy{} //空结点

	//判断
	if num < 1 	{
		fmt.Println("num的值不对")
		return first
	}
	//循环的构建这个环形链表
	for i := 1; i <= num; i++ {
		boy := &Boy{
			No : i,
		}
		//分析构成循环链表，需要一个辅助指针[帮忙的]
		//1. 因为第一个小孩比较特殊
		if i == 1 { //第一个小孩
			first = boy //不要动
			curBoy = boy
			curBoy.Next = first //
		} else {
			curBoy.Next = boy
			curBoy = boy
			curBoy.Next = first //构造环形链表
		}
	}
	return first

}

//显示单向的环形链表[遍历]
func ShowBoy(first *Boy) {

	//处理一下如果环形链表为空
	if first.Next == nil {
		fmt.Println("链表为空，没有小孩...")
		return
	}

	//创建一个指针，帮助遍历.[说明至少有一个小孩]
	curBoy := first
	for {
		fmt.Printf("小孩编号=%d ->", curBoy.No)
		//退出的条件?curBoy.Next == first
		if curBoy.Next == first {
			break
		}
		//curBoy移动到下一个
		curBoy = curBoy.Next
	}
}

func StartGame(first *Boy, startNo int, countNum int) {
	//1. 空的链表我们单独的处理
	if first.Next == nil {
		fmt.Println("空的链表，没有小孩")
		return
	}
	//留一个，判断 startNO <= 小孩的总数
	//2. 需要定义辅助指针，帮助我们删除小孩
	tail := first
	//3. 让tail执行环形链表的最后一个小孩,这个非常的重要
	//因为tail 在删除小孩时需要使用到.
	for {
		if tail.Next == first { //说明tail到了最后的小孩
			break
		}
		tail = tail.Next
	}
	//4. 让first 移动到 startNo [后面我们删除小孩，就以first为准]
	for i := 1; i <= startNo - 1; i++ {
		first = first.Next
		tail = tail.Next
	}
	fmt.Println()
	//5. 开始数 countNum, 然后就删除first 指向的小孩
	for {
		//开始数countNum-1次
		for i := 1; i <= countNum -1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("小孩编号为%d 出圈 \n", first.No)
		//删除first执行的小孩
		first = first.Next
		tail.Next = first
		//判断如果 tail == first, 圈子中只有一个小孩.
		if tail == first {
			break
		}
	}
	fmt.Printf("小孩小孩编号为%d 出圈 \n", first.No)

}

func main() {
	first := AddBoy(5)
	//显示
	ShowBoy(first)
	StartGame(first, 2, 3)
}
*/
func Josephu() {
	// 将数据存入循环链表中
	list := new(CircularLinkNode)
	list.Create(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
		17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32)
	fmt.Println("原始数据：")
	list.Print()
	fmt.Println("删除数据：")
	i := 1
	for list.Length() > 2 {
		i += 3
		if i > list.Length() {
			i = i % list.Length()
			// 注意这里的特殊情况考虑
			if i == 0 { // 整除了, i就等于list的长度
				i = list.Length()
			}
		}
		list.DeleteByIndex(i)
		list.Print()
	}
}


func OutCall_c() {
	//cl := new(CircularLinkNode)
	//cl.Create(1, 2, 3)
	//cl.Print()
	//count := cl.Length()
	//fmt.Println("长度:", count)
	//cl.InsertByIndex(4, 4)
	//cl.Print()
	//cl.DeleteByIndex(5)
	//cl.Print()
	//cl.Destroy()
	//fmt.Println("销毁:")
	//cl.Print()

	// 约瑟夫问题
	Josephu()
}
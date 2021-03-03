package main

import (
	"fmt"
)

// 根节点, 子节点, 子叶节点
type BinaryTreeNode struct {
	Data 	int
	LChild *BinaryTreeNode // 左子树
	RChild *BinaryTreeNode // 右子树
}

// create 根据前序遍历与中序遍历构建二叉树
// ** 我们能由二叉树前序, 中序等遍历得到顺序字符串 **
// ** 但我们并不能从某一种遍历方式推断出唯一的二叉树(这类似于hash运算) **
func CreateBinaryTree(preOrder, midOrder []int) *BinaryTreeNode {
	if len(preOrder) == 0 || len(midOrder) == 0 {
		return nil
	}
	// 在前序遍历中preOrder[0]就是根节点的值
	node := &BinaryTreeNode{preOrder[0], nil, nil}
	// 获取根节点的值在中序遍历中的索引, 这样这个索引之前的数据都是左节点, 之后的数据都是右节点的
	// 对应前序遍历就是 [1:index+1]的是左子树, [index+1:]之后的是右子树
	var index int
	for k, v := range midOrder {
		if v == preOrder[0] {
			index = k
			break
		}
	}
	// 递归创建左右节点
	node.LChild = CreateBinaryTree(preOrder[1:index+1], midOrder[:index])
	node.RChild = CreateBinaryTree(preOrder[index+1:], midOrder[index+1:])
	return node
}

// 先序遍历 - 根左右
func PreOrder(root *BinaryTreeNode) {
	// 递归退出条件
	if root == nil {
		return
	}

	fmt.Print(root.Data, " ")
	PreOrder(root.LChild)
	PreOrder(root.RChild)
}

// 中序遍历 - 左根右
func MidOrder(root *BinaryTreeNode) {
	if root == nil {
		return
	}

	MidOrder(root.LChild)
	fmt.Print(root.Data, " ")
	MidOrder(root.RChild)
}

// 后序遍历 - 左右根
func RearOrder(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	RearOrder(root.LChild)
	RearOrder(root.RChild)
	fmt.Print(root.Data, " ")
}

// maxDepth - 最大深度(根节点到最远 叶子节点 最长路径上的节点数)
func MaxDepth(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}
	// 当前节点深度计为 1
	// 递归计算左右子树的深度, 然后取其中的最大值. 再加上当前节点的深度1得到总深度
	lDepth := MaxDepth(root.LChild)
	rDepth := MaxDepth(root.RChild)
	if lDepth > rDepth {
		return lDepth + 1
	} else {
		return rDepth + 1
	}
}

// minDepth - 最小深度(根节点到最近 叶子节点 最短路径上的节点数)
// 这里有个小陷阱,不能跟上面反过来, 因为叶子节点是没有左右子树的节点
// 如果按上面的方案反过来, 在[1,2]这种情况下,会返回1, 而实际是2
// 所以要多一步判定,左子树为0,返回右子树深度+1,或者右子树为0,返回左子树深度+1
// 左右子树都为0,则返回1. 判断根节点没有左右子树计算深度的特殊情况(深度是计算到 叶子节点 的)
func MinDepth(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}

	lDepth := MinDepth(root.LChild)
	rDepth := MinDepth(root.RChild)

	// 判断根节点只有左子树或右子树计算到叶子的深度情况
	//if lDepth == 0 {
	//	return rDepth+1
	//} else if rDepth == 0 {
	//	return lDepth+1
	//}
	// 上面形式的简化
	if lDepth == 0 || rDepth == 0 {
		return lDepth + rDepth + 1
	}

	if lDepth < rDepth {
		return lDepth+1
	} else {
		return rDepth+1
	}
}

// leaf count
func LeafCount(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}
	// 没有左右子树即为叶子节点
	if root.LChild == nil && root.RChild == nil {
		return 1
	}

	lCount := LeafCount(root.LChild)
	rCount := LeafCount(root.RChild)

	return lCount + rCount
}

// search
func SearchTree(root *BinaryTreeNode, data int) bool {
	if root == nil {
		return false
	}

	if root.Data == data {
		return true
	}

	return SearchTree(root.LChild, data) || SearchTree(root.RChild, data)
}

// reverse 翻转二叉树
func ReverseTree(root *BinaryTreeNode) {
	if root == nil {
		return
	}

	// 翻转当前节点的左右子树
	root.LChild, root.RChild = root.RChild, root.LChild
	// 递归翻转左右子树
	ReverseTree(root.LChild)
	ReverseTree(root.RChild)
}

// copy - 创建一个新的二叉树, 修改里面的值并不会影响原来的二叉树
func CopyTree(dst, src *BinaryTreeNode) {
	if src == nil || dst == nil {
		return
	}

	// 复制当前节点
	dst.Data = src.Data
	var lChild, rChild *BinaryTreeNode
	// 这里需要创建新的节点, 而不是直接赋值src的
	if src.LChild != nil {
		lChild = new(BinaryTreeNode)
	}
	if src.RChild != nil {
		rChild = new(BinaryTreeNode)
	}
	dst.LChild = lChild
	dst.RChild = rChild

	CopyTree(lChild, src.LChild)
	CopyTree(rChild, src.RChild)
}

// destroy
func DestroyTree(root *BinaryTreeNode) {
	if root == nil {
		return
	}

	DestroyTree(root.LChild)
	root.LChild = nil
	DestroyTree(root.RChild)
	root.RChild = nil
	root.Data = 0
}

func OutCall_t() {
	preOrder := []int{1, 2, 4, 8, 9, 5, 3, 6, 7}
	midOrder := []int{8, 4, 9, 2, 5, 1, 6, 3, 7}
	tree := CreateBinaryTree(preOrder, midOrder)
	fmt.Println(tree)
	// 遍历二叉树
	fmt.Print("前序遍历:")
	PreOrder(tree)
	fmt.Println()
	fmt.Print("中序遍历:")
	MidOrder(tree)
	fmt.Println()
	fmt.Print("后序遍历:")
	RearOrder(tree)
	fmt.Println()
	// depth
	max := MaxDepth(tree)
	min := MinDepth(tree)
	fmt.Println("maxDepth:", max, " minDepth:", min)
	// leaf
	leaf := LeafCount(tree)
	fmt.Println("leaf count:", leaf)
	// search
	res := SearchTree(tree, 8)
	fmt.Println("num is exist:", res)
	// reverse
	ReverseTree(tree)
	fmt.Print("翻转后的树:")
	PreOrder(tree)
	fmt.Println()
	// copy
	newTree := new(BinaryTreeNode)
	CopyTree(newTree, tree)
	newTree.LChild.Data = 66
	fmt.Print("copy后的新树:")
	PreOrder(newTree)
	fmt.Println()
	// destroy
	DestroyTree(newTree)
	fmt.Print("销毁树:")
	PreOrder(newTree)
}
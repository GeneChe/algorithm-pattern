package main

/*
#include "stdlib.h"
*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

const Size = 8

type Slice struct {
	// 任何类型的指针值都可以转换为Pointer   unsafe.Pointer(指针)
	// Pointer可以转换为任何类型的指针值   (*float32)(Pointer)
	// uintptr is an integer type, 他能存储地址的值, 所以能进行指针运算, 相当于整形运算
	// uintptr可以转换为 Pointer
	// Pointer可以转换为 uintptr, 其他类型指针无法转换成uintptr
	// Pointer指针只能赋值指针类型, 无法通过*Pointer操作指针指向的内存和取其中的值
	Data unsafe.Pointer // 这里使用unsafe.Pointer(万能指针), 而不是interface{}, 是为了方便后面的指针偏移运算
	Len int
	Cap int
}

// create
func (s *Slice) Create(l, c int, data ...int) error {
	if len(data) == 0 {
		return errors.New("未添加数据")
	}
	if l < 0 || c < 0 || l > c || len(data) > l {
		return errors.New("参数错误")
	}

	if s == nil {
		return errors.New("对象不存在")
	}

	// 开辟空间
	// cannot use c * Size (type int) as type _Ctype_ulong in argument to _Cfunc__CMalloc
	// malloc是C的方法需要将go的int类型转换c的类型, C.ulong()强转
	// s.Data = C.malloc(c*Size)
	// 如果堆空间malloc开辟失败会返回nil(即内存地址编号为0的空间)
	s.Data = C.malloc(C.ulong(c) * Size)
	if s.Data == nil {
		return errors.New("malloc开辟空间失败")
	}
	s.Len = l
	s.Cap = c

	// 存储数据
	// 定义临时变量p保存s.Data的地址, 存储数据时操作p
	// 由于普通类型指针(包含Pointer)无法进行指针运算
	// 只有将普通类型指针转换成Pointer, 然后再转换成uintptr才能进行指针运算
	p := uintptr(s.Data)
	for _, v := range data {
		// uintptr是整形, 需要将其转换成pointer指针, 再转成普通指针才能赋值和取值
		*(*int)( unsafe.Pointer(p) ) = v
		p += Size
	}

	return nil
}

// print
func (s *Slice) Print() {
	if !s.IsValid() { // 避免s为空指针
		fmt.Println("[]")
		return
	}

	//遍历
	p := uintptr(s.Data)
	for i := 0; i < s.Len; i++ {
		fmt.Print(*(*int)( unsafe.Pointer(p) ), " ")
		p += Size
	}

	fmt.Println()
	return
}

// append
func (s *Slice) Append(data ...int) {
	// 参数判断
	if !s.IsValid() || len(data) == 0 {
		return
	}

	// 计算容量, 切片扩容后的容量是扩容前的两倍, 超过1024以后, 扩容1/4
	var flag bool
	for s.Len + len(data) > s.Cap {
		// 更新cap
		s.Cap *= 2
		// 标记
		flag = true
	}
	if flag {
		// 指针 = realloc(指针, 总容量)
		s.Data = C.realloc(s.Data, C.ulong(s.Cap) * 2 * Size)
	}

	// 将指针偏移到最后
	p := uintptr(s.Data)
	p += uintptr(s.Len) * Size // ** 注意: 这里p如果越界, 编译器不会报错 **

	// 赋值
	for _, v := range data {
		*(*int)(unsafe.Pointer(p)) = v
		p += Size
	}

	// 更新len
	s.Len += len(data)
}

// getData 根据索引获取元素
func (s *Slice) GetData(index int) int {
	// 判断是否有效
	if !s.IsValid() {
		return 0
	}
	// 判断索引范围, 避免越界. (此时越界运行并不会报错, 但这是非法使用)
	if index < 0 || index >= s.Len {
		return 0
	}

	// 根据下标偏移指针
	p := uintptr(s.Data)
	p += uintptr(index) * Size

	return *(*int)(unsafe.Pointer(p))
}

// searchData 根据元素获取索引, 返回-1表示查找失败
func (s *Slice) SearchData(data int) int {
	if !s.IsValid() {
		return -1
	}

	// 顺序比较
	p := uintptr(s.Data)
	for i := 0; i < s.Len; i++ {
		// 返回第一次出现data的下标
		if *(*int)(unsafe.Pointer(p)) == data {
			return i
		}
		p += Size
	}

	return -1
}

// 检车s或s.data是否为空
func (s *Slice) IsValid() bool {
	return s != nil && s.Data != nil
}

// delete
func (s *Slice) Delete(index int) {
	// 判断对象
	if !s.IsValid() {
		return
	}
	// 判断参数
	if index < 0 || index >= s.Len {
		return
	}
	// 将指针定位到索引位置
	p := uintptr(s.Data)
	p += uintptr(index) * Size
	// 将索引位置之后的数据往前移动
	for i := index; i < s.Len - 1; i++ {
		// 将后一个指针指向的值赋值给当前指针指向的值
		*(*int)(unsafe.Pointer(p)) = *(*int)(unsafe.Pointer(p+Size))
		// 移动一次指针
		p += Size
	}
	// 更新len
	s.Len--
}

// insert (暂不考虑扩容问题)
func (s *Slice) Insert(index, data int) {
	if !s.IsValid() {
		return
	}
	if index < 0 || index > s.Len {
		return
	}

	// 1. 将index位置后面的元素顺序往后移动一位
	// 将指针偏移到最后
	p := uintptr(s.Data)
	p += uintptr(s.Len) * Size
	// 从后往前遍历
	for i := s.Len; i > index; i-- {
		// 将前一个指针指向的值赋值给后一个指针指向的值
		*(*int)(unsafe.Pointer(p)) = *(*int)(unsafe.Pointer(p-Size))
		// 将指针前移一位
		p -= Size
	}

	// 2. 给index位置赋新的值
	p = uintptr(s.Data)
	p += uintptr(index) * Size
	*(*int)(unsafe.Pointer(p)) = data

	// 3. 长度++
	s.Len++
}

// destroy
func (s *Slice) Destroy() {
	// 释放堆内存空间
	C.free(s.Data)
	s.Data = nil
	s.Len = 0
	s.Cap = 0
}

func OutCall() {
	var s Slice
	s.Create(2, 2, 1, 2)
	fmt.Println(s)
	s.Print()
	s.Append(4, 5, 6)
	fmt.Println(s)
	s.Print()

	v := s.GetData(10)
	fmt.Println("值:", v)
	index := s.SearchData(5)
	fmt.Println("下标:", index)
	s.Delete(2)
	fmt.Println(s)
	s.Print()
	s.Delete(2)
	fmt.Println(s)
	s.Print()

	s.Insert(3, 3)
	fmt.Println(s)
	s.Print()
	s.Insert(2, 18)
	fmt.Println(s)
	s.Print()

	s.Destroy()
	fmt.Println(s)
}
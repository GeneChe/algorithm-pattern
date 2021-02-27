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
	s.Data = C.malloc(C.ulong(c) * Size)
	s.Len = l
	s.Cap = c

	// 存储数据
	// 定义临时变量p保存s.Data的地址, 存储数据时操作p
	// 由于普通类型指针(包含Pointer)无法进行指针运算
	// 只有将普通类型指针转换成Pointer, 然后再转换成uintptr才能进行指针运算
	p := uintptr(s.Data)
	for _, v := range data {
		// uintptr虽可以指针运算, 但go并不把它当成指针, 他无法拥有对象
		// 还需要将其转换成普通类型指针才能赋值
		*(*int)( unsafe.Pointer(p) ) = v
		p += Size
	}

	return nil
}

// print
func (s *Slice) Print() {
	if s == nil || s.Len == 0 { // 避免s为空指针
		fmt.Println("[]")
		return
	}

	//遍历
	p := uintptr(s.Data)
	for i := 0; i < s.Len; i++ {
		fmt.Print(*(*int)( unsafe.Pointer(p) ), " ")
		p += Size
	}

	return
}

func OutCall() {
	var s Slice
	s.Create(3, 5, 1, 2, 3)
	fmt.Println(s)
	s.Print()
}
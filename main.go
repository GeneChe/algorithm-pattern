package main

// 求最大公约数
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// 斐波纳契数列
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func main() {
	/*
		基础知识
		// slice 结构是由 1. 数据的内存地址(指针类型占用8字节) 2. len 有效数据长度 3. cap可扩容的有效容量 组成	24字节
		// 切片名是一个地址, 它指向其存储元素的首地址. 切片是引用传递
		var s []int
		var i int
		var f float64

		// unsafe.Sizeof() 计算数据类型在内存中占用的字节大小
		fmt.Println("float64占得内存大小:", unsafe.Sizeof(f))
		fmt.Println("int占得内存大小:", unsafe.Sizeof(i))
		fmt.Println("slice占得内存大小:", unsafe.Sizeof(s))
	*/

	/*
		获取接口变量实际的类型
		var i interface{}
		i = 10
		fmt.Printf("%T\n", i)

		// 1. 通过类型断言转换类型
		//if value, ok := i.(int); ok {
		//	fmt.Printf("%T\n", value)
		//}

		// 2. 通过反射获取接口中变量的类型
		t := reflect.TypeOf(i) // t 存储的是数据类型
		fmt.Println(t)	// int
		v := reflect.ValueOf(i)
		fmt.Println(v)
	 */

	//s := []int{1, 2, 3}
	//fmt.Println(s[1:1]) // [] 不报错

	// 内存地址是 无符号十六进制整形数据
	// 空指针可以调用方法
	/*
		var s Slice
		a := &s
		a = nil
		a.Append(1)
	*/

	// OutCall()
	// OutCall_l()
	// OutCall_d()
	//OutCall_c()
	//OutCall_s()
	//OutCall_q()
	//OutCall_t()
	OutCall_st()
}

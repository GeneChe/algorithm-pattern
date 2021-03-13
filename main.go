package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

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

// 错误重试
func waitForServer(url string) error {
	const timeout = time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("server not responding (%s);retrying...", err)
		time.Sleep(time.Second << tries)
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

// defer
func bigSlowOperation() {
	// 第一次遇到defer时会计算trace(""), 函数退出时会执行trace的返回值(一个函数)
	defer trace("bigSlowOperation")() // 注意最后面加上()
	time.Sleep(time.Second * 5)
}

// 统计函数执行时间
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func () {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

// T 实现了接口中的方法, 那么接口变量赋值就是 T或*T 类型的值
// *T 实现了接口中的方法, 那么接口变量赋值就是 *T 类型的值
type aa int

func (a aa) Write(p []byte) (int, error) {
	return 0, nil
}

func main() {
	// interface
	//var a aa
	//var w io.Writer
	//w = &a
	//w = a
	//fmt.Fprintln(w, "")

	// 接口值
	// 一个包含nil指针的接口不是nil接口
	//var buf *bytes.Buffer
	//w = buf
	//if w != nil {
		// 这个条件永远为true, 虽然w的动态值为nil, 但是其动态类型为(*bytes.Buffer), 所以w!=nil成立
		// 解决方法: 将buf定义为 io.Writer类型
		// nil的接口是 动态值和动态类型都为nil
	//}
	//sort.Sort()
	//sort.Strings()
	//sort.Reverse()
	//sort.IsSorted()

	// 转义url中的特殊字符&或?
	//url.QueryEscape("")
	//http.StatusOK
	/*
		var rmdirs []func()
		dirs := tempDirs()
		for i := 0; i < len(dirs); i++ {
			os.MkdirAll(dirs[i], 0755) // OK
			rmdirs = append(rmdirs, func() {
				os.RemoveAll(dirs[i]) // NOTE: incorrect!
				// 循环结束时, 变量i保存的是最后一次循环的值, 而不是每次循环的值
				// 解决方式就是 在每次循环时都定义一个局部变量, 常用与循环变量同名的局部变量 i := i
			})
		}
		**匿名函数内, 存的是外部局部变量的地址**

		defer语句中的函数会在 return语句更新返回值变量 后 再执行

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
	//OutCall_st()
	OutCall_h()

	// append nil 不出错
	//var a []int
	//a = nil
	//a = append(a, 1)
	//fmt.Println(a)

	// err retry
	//if err := waitForServer("www.baidu.com"); err != nil {
	//	log.Fatalf("Site is down: %v\n", err)
	//}
	//bigSlowOperation()
}
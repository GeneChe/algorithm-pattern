package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
)

func OutCall_r() {
	// 类型:
	// reflect.Type 	interface
	// reflect.Value    struct
	// 方法:
	// reflect.TypeOf()  // 返回接口中的 动态类型 (返回具体的类型)
	// reflect.ValueOf() // 返回接口中的 动态值
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w)) // *os.File
	// fmt.Printf 提供了一个缩写 %T 参数，内部使用 reflect.TypeOf 来输出
	// fmt 包的 %v 标志参数会对 reflect.Values 特殊处理

	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Println(v.String())
	t := v.Type()
	fmt.Println(t.String())
	// 逆操作 reflect.Value.Interface
	x := v.Interface()
	i := x.(int)
	fmt.Println(i)
	/*
		reflect.Value 和 interface{} 都能装载任意的值。
		所不同的是，一个空的接口隐藏了值内部的表示方式和所有方法，
		因此只有我们知道具体的动态类型才能使用类型断言来访问内部的值（就像上面那样），内部值我们没法访问。
		相比之下，一个 Value 则有很多方法来检查其内容
	*/

	// reflect.Value 的 Kind检查值类型
	// 空的 reflect.Value 的 kind 即为 Invalid
	k := v.Kind()
	fmt.Printf("%d\n", k)
	s := strconv.FormatInt(v.Int(), 2) // base是数字的进制
	fmt.Println(s)

	// 所有通过reflect.valueOf() 得到的value都是不可取地址的
	// reflect.ValueOf(&x).Elem() 解引用方式生成的，指向另一个变量，因此是可取地址的
	fmt.Println("&v: ", v.CanAddr())

	// 对于slice和array的方法
	// v.Len()
	// v.Index(1)

	// 对于struct的方法
	// v.NumField()
	// v.Type().Field(1).Name // 成员名
	// v.Field(1) // 成员值

	// 对于map
	// v.MapKeys()
	// v.MapIndex(key) // 参数是key

	// 对于ptr
	// v.IsNil()
	// v.Elem() // 指针指向的值

	// 对于接口
	// v.IsNil()
	// v.Elem()
}

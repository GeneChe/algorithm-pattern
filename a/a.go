package a

// 循环导入包, 编译时查不出, 运行时会报错
import "algorithm-pattern/b"

func AA() {
	b.BB()
}

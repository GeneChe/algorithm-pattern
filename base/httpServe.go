package base

import (
	"fmt"
	"log"
	"net/http"
)

type database map[string]int

// 指针类型是可比较的: 两个指针指向同一个变量，则这两个指针相等，或者两个指针同为nil，它们也相等。指针值可以与nil比较。
func OutCall_h() {
	// error
	//a, b := errors.New("hello"), errors.New("hello")
	// // if a == b { // false	两个指针指向的地址不同, 即指向的不是同一个变量, 所以不等
	//if a.Error() == b.Error() { // true
	//	fmt.Println("true")
	//} else {
	//	fmt.Printf("%T, %T", a, b)
	//}

	//io.WriteString()

	db := database{"shorts": 22, "socks":50}
	mu := http.NewServeMux()
	mu.HandleFunc("/list", db.list)
	mu.HandleFunc("/price", db.price) // web服务器在一个新的协程中调用每一个handler

	// http包提供了一个全局的serveMux实例 --- DefaultServeMux
	// http.HandleFunc("/price", db.price)
	// ListenAndServe的第二个参数传nil, 时默认使用DefaultServeMux,
	log.Fatal(http.ListenAndServe("localhost:8000", mu))
}

func (d database) list(w http.ResponseWriter, r *http.Request)  {
	// r.ParseForm()
	// r.Form
	for k, v := range d {
		fmt.Fprintf(w, "%s=%d\n", k, v)
	}
}

func (d database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	v, ok := d[item]
	if !ok {
		http.Error(w, "item no found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "%s=%d\n", item, v)
}
package main

import "fmt"

func main() {
	// make函数构造切片
	var a []string = make([]string, 5)
	var b []int     //此时的b是nil
	var e = []int{} //声明且初始化, 已经不是nil了，所以不能通过一个切片是否是nil的来判断是否是空的，需要通过len()
	a[0] = "sichuan"
	// b[0] = 1
	var c = [6]int{1, 2, 3, 4, 5, 6} //为啥数组的显示类型声明是这样
	b = c[1:5]
	// 基于数组得到切片
	fmt.Println(b)
	// 切片的底层实现原理
	// 切片不能直接比较
	// 切片的赋值拷贝
	// 切片的遍历
	// 切片的扩容
	var d string = "abcd"
	fmt.Printf(d)
	fmt.Println(e)
	// 追加多个元素，追加一个切片
	// 容量会变化，即动态扩容，应该尽量避免
}

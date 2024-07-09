package main

import (
	"fmt"
)

func main() {
	var a map[string]int
	//a["jianqiu"] = 2 //这种是错的，仅仅声明了，但还没有初始化，没有内存
	a = make(map[string]int, 8) //初始化
	a["jianqiu"] = 2
	fmt.Println(a)
	//声明，初始化, 加键值对
	//声明的同时初始化, 然后加键值对
	b := map[int]string{}
	b[2] = "jianqiu"
	fmt.Println(b)
	//如何判断一个map里面是否存在某一个键
	v, ok := b[2]
	if ok {
		fmt.Println(v)
	}
}

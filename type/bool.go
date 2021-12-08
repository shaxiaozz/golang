package main

import "fmt"

func main() {
	/*
		1、变量定义
			布尔类型：飙升真假
			标识符：bool
			字面量：true/false
			零值：false
		2、操作
			逻辑运算：与或非，&& || !
			关系运算：==,!=
	*/
	var zero bool
	// 简短声明变量
	isBoy := true
	isGirl := false
	fmt.Println(zero, isBoy, isGirl)
	fmt.Printf("%T %t %t \n", zero, isBoy, isGirl)

	// 定义常量
	const (
		A int  = 1
		B int  = 2
		C int  = 1
		D bool = true
		E bool = false
		F bool = true
	)
	// 逻辑运算
	fmt.Println(D && F)
	fmt.Println(D && E)
	fmt.Println(E && F)
	fmt.Println(D || E)
	fmt.Println(E || F)

	// 关系运算
	fmt.Println(A == B)
	fmt.Println(A == C)
	fmt.Println(A != B)
	fmt.Println(A != C)
}

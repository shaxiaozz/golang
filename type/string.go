package main

import "fmt"

func main() {
	// 定义字符串类型的变量，字面量有两种方式，双引号与反引号。反引号表示元字符(即字面量是什么就是什么，例如/n，/t，/r等)
	var name string = "zengfengjin"
	fmt.Printf("%T %s", name, name)

	// 字符索引(检索某个)
	desc := "abcdefg"
	fmt.Printf("\n %T %v", desc[0], desc[0])

	// 字符索引(检索多少，一个范围)
	fmt.Printf("\n %T %v \n", desc[0:3], desc[0:3])

	// 统计变量字符个数
	fmt.Println(len(desc))
}

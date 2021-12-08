package main

import "fmt"

func main() {
	// 作用域：定义标识符可以使用的范围
	// 在GO中用{}来定义作用域的范围
	// 使用原则：子语句块可以使用父语句块的标识符，父不能使用子的

	//简短声明变量
	name := "傻笑zz"
	//定义作用域
	{
		age := 21
		fmt.Println(name)
		fmt.Println(age)
	}
}

package main

import "fmt"

//定义常量变量(不可再次赋值的变量)
const (
	NAME   = "傻笑zz"
	AGE    = 21
	HIGHT  = 1.81
	GENDER = "男"
)

//定义枚举(枚举的意思即为：列举，将所有的情况都列举出来)
const (
	A = iota
	B
	C
)

func main() {
	//格式化输出常量
	fmt.Print("姓名: ", NAME, "\n", "年龄: ", AGE, "\n", "身高: ", HIGHT, "\n", "性别: ", GENDER, "\n")

	//输出枚举
	fmt.Println(A, B, C)
}

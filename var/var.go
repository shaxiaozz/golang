package main

import "fmt"

//定义包级别变量(手动指定变量类型)
var (
	/*
		1、定义年龄变量(age)，类型为int，初始值为21
		2、定义身高变量(hight)，类型为float64，初始值为1.81
		3、定义性别(gender)，类型为string，初始值为男
		4、定义程序版本(version)，类型为string，初始值为1.0
	*/
	age     int     = 21
	hight   float64 = 1.81
	gender  string  = "男"
	version string  = "1.0"
)

func main() {
	//简短声明变量(只能在函数内部使用)
	name, isStudent := "傻笑zz", false

	//格式化输出变量
	fmt.Print("姓名: ", name, "\n", "性别: ", gender, "\n", "年龄: ", age, "\n", "身高: ", hight, "\n", "是否为学生: ", isStudent)

	//定义不指定类型的变量(重新赋值)
	var name_1, gender_1, age_1, hight_1, isStudent_1 = name, gender, age, hight, isStudent

	//格式化输出变量
	fmt.Print("\n", "姓名: ", name_1, "\n", "性别: ", gender_1, "\n", "年龄: ", age_1, "\n", "身高: ", hight_1, "\n", "是否为学生: ", isStudent_1)
}

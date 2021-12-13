/*
根据成绩来判断学生的评分等级
1、100-90：为优秀
2、89-70：为良好
3、60-69：为及格
4、低于60：为不及格
*/
package main

import "fmt"

func main() {
	// 初始化变量
	var score int

	// 获取用户输入
	fmt.Print("请输入您的成绩：")
	fmt.Scan(&score)

	// switch判断
	switch {
	case score <= 100 && score >= 90:
		fmt.Println("您的成绩评定为：优秀")
	case score <= 89 && score >= 70:
		fmt.Println("您的成绩评定为：良好")
	case score <= 69 && score >= 60:
		fmt.Println("您的成绩评定为：及格")
	case score <= 60 && score >= 0:
		fmt.Println("您的成绩评定为：不及格")
	default:
		fmt.Println("请输入正确的成绩(0-100)")
	}
}

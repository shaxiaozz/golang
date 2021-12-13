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

	// 获取成绩
	fmt.Print("请输入您的成绩：")
	fmt.Scan(&score)

	// if判断
	if score >= 90 && score <= 100 {
		fmt.Println("您的成绩评定为：优秀")
	} else if score < 90 && score >= 70 {
		fmt.Println("您的成绩评定为：良好")
	} else if score < 70 && score >= 60 {
		fmt.Println("您的成绩评定为：及格")
	} else if score < 60 && score >= 0 {
		fmt.Println("您的成绩评定为：不及格")
	} else {
		fmt.Println("请输入正确的成绩(0-100)")
	}
}

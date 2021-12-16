/*
 99乘法表
 1X1=1
 1X2=2   2X2=4
 1X3=3   2X3=6  3X3=9
 ......
 1X9=9   ........................................... 9X9=81
*/

package main

import "fmt"

func main() {
	// 循环第一次，获取第一个1到9
	for line := 1; line <= 9; line++ {
		// 循环第二次，并根据line来做表达式。因此：当line为1时，第二次则只循环一次。当line为2时，第二次循环则循环两次
		for i := 1; i <= line; i++ {
			fmt.Print(i, "X", line, "=", line*i, "\t")
		}
		// 第一次循环后分行显示
		fmt.Println()
	}
}

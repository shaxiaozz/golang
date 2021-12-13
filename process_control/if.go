package main

import "fmt"

func main() {
	// if控制.
	// 老婆叫老公下班回来后，买10个包子，如果有卖西瓜的就买一个。
	var yes string
	fmt.Print("有卖西瓜的吗?(Y/N)")
	// 使用指针修改变量yes的值
	fmt.Scan(&yes)

	// if判断
	if yes == "Y" || yes == "y" {
		fmt.Println("\n老婆的想法：买10个包子和1个西瓜")
		fmt.Println("老公的想法：买1个西瓜")
	} else if yes == "N" || yes == "n" {
		fmt.Println("\n老婆的想法：买10个包子")
		fmt.Println("老公的想法：买10个包子")
	}
}

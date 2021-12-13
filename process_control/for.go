/*
 使用for 4种语法分别打印以下需求：
 1、计算1-100的总和
 2、计算1-100中基偶数的和
 3、遍历打印字符串中的字符
 4、for 死循环打印

 提前结束循环需求
 5、只打印1-10的偶数
 6、变量字符串"我是傻笑zz"，当循环到z的时候退出for循环
*/

package main

import (
	"fmt"
)

func main() {
	// 第一种需求及实现方式
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("第一种需求：")
	fmt.Println("1-100的总和为：", sum)

	// 第二种需求及实现方式
	jishu := 0
	oushu := 0
	i := 1
	for i <= 100 {
		// 根据对变量i的值取模进行判断
		if i%2 == 0 {
			oushu += i
		} else {
			jishu += i
		}
		i++
	}
	fmt.Println("第二种需求：")
	fmt.Println("1-100中偶数的和为：", oushu)
	fmt.Println("1-100中基数的和为：", jishu)

	// 第三种需求及实现方式
	name := "我是傻笑zz"
	fmt.Println("第三种需求：")
	for _, ch := range name {
		fmt.Println(string(ch))
		//fmt.Printf("\n%q", ch)
	}

	// 第四种需求及实现方式
	// fmt.Println("第四种需求：")
	// ii := 0
	// for {
	// 	fmt.Println(ii)
	// 	ii++
	// }

	// 第五个需求
	fmt.Println("第五种需求：")
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue // 退出本次循环
		}
		fmt.Println(i)
	}

	// 第六个需求
	fmt.Println("第六种需求：")
	myName := "我是傻笑zz"
	for _, ch := range myName {
		if string(ch) == "z" {
			break // 终止整个for循环
		}
		fmt.Println(string(ch))
	}
}

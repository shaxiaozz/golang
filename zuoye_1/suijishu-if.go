package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 初始化相关变量
	const NUMBER int = 5
	var intput, suijishu int

	// 初始化随机数种子,并生成0-100的随机数
	rand.Seed(time.Now().Unix())

	// 死循环，直到用户确定退出游戏
	for {
		// 生成随机数
		suijishu = rand.Intn(100)
		for i := 1; i <= NUMBER; i++ {
			// 获取用户输入
			fmt.Printf("请输入您的随机数：")
			fmt.Scan(&intput)

			// 主判断流程
			if intput == suijishu {
				fmt.Println("猜对了！！！")
				break
			} else if intput > suijishu {
				fmt.Printf("猜大了！！,你还有%d次机会\n", NUMBER-i)
			} else {
				fmt.Printf("猜小了！！,你还有%d次机会\n", NUMBER-i)
			}

			// 判断是否为最后一次机会
			if i == NUMBER {
				fmt.Println("太笨了！！5次机会都猜错。")
			}
		}

		// 获取用户输入，是否重新开始游戏
		var yes string
		fmt.Printf("是否重新开始游戏?(y/n)")
		fmt.Scan(&yes)
		// 当用户输入n时，退出死循环
		if yes != "y" {
			break
		}
	}

}

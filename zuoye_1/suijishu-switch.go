package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 初始化变量
	var intput, suijishu int
	const NUMBER int = 5

	// 初始化随机数种子
	rand.Seed(time.Now().Unix())

	for {
		suijishu = rand.Intn(100)
		for i := 1; i <= NUMBER; i++ {
			fmt.Print("请输入您的随机数：")
			fmt.Scan(&intput)

			switch {
			case intput == suijishu:
				fmt.Println("猜中了！！！")
				// 跳出循环
				goto END
			case intput > suijishu:
				fmt.Printf("猜大了，您还有%d次机会\n", NUMBER-i)
			default:
				fmt.Printf("猜小了，您还有%d次机会\n", NUMBER-i)
			}

			// 判断是否为最后一次机会
			switch i {
			case NUMBER:
				fmt.Println("太笨了，5次机会都猜错了！")
			}
		}
	END:
		var yes string
		fmt.Printf("是否重新开始游戏?(y/n)")
		fmt.Scan(&yes)
		if yes != "y" {
			break
		}
	}
}

package main

import "fmt"

func main() {
	// 整数类型
	// 标识符：int/int*/uint
	// 字面量：十进制，八进制，十六进制
	// 操作：算术运算(+,-,*,/,%,++,--)
	var age int = 21
	fmt.Printf("%T %d \n", age, age)

	// 算术运算(+,-,*,/,%,++,--)
	fmt.Println(1 + 2)
	fmt.Println(1 - 1)
	fmt.Println(1 * 9)
	fmt.Println(9 / 2)
	fmt.Println(9.0 / 2)
	fmt.Println(9 % 2)
	fmt.Println(10 % 2)

	// 自增，自减
	hight := 18
	hight++
	fmt.Println(hight)
	hight--
	fmt.Println(hight)

	// 关系运算(==，!=，>，<，>=，<=)
	fmt.Println("1 == 2吗", 1 == 2)
	fmt.Println("1 != 2吗", 1 != 2)
	fmt.Println("1 > 2吗", 1 > 2)
	fmt.Println("1 < 2吗", 1 < 2)
	fmt.Println("1 >= 2吗", 1 >= 2)
	fmt.Println("1 <= 2吗", 1 <= 2)

}

package main

// 导入fmt包，fmt包实现：格式化I/O输入输出
import "fmt"

//定义变量
var (
	name  string  = "傻笑zz"
	age   int     = 21
	hight float64 = 1.81
)

func main() {
	// 使用fmt.Println打印变量(默认换行)
	fmt.Println("以下是Println打印的内容，默认已换行")
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(hight)

	// 定义作用域
	{
		// 使用fmt.Print打印变量(无换行)
		fmt.Print("以下是Print打印的内容,默认不换行。如需换行需添加\\n")
		fmt.Print(name)
		fmt.Print(age)
		fmt.Print(hight)
	}

	// 定义作用域
	{
		// 使用fmt.Printf打印变量值与变量类型
		/*
			1、%T：打印变量类型
			2、%d：已整数类型打印
			3、%s：已字符串类型打印
			4、%f：已浮点数类型打印
			5、%t：已布尔值类型打印(真或假这个词)
		*/
		isBoy := true
		fmt.Printf("%T %T %T", name, age, hight)
		fmt.Printf("\n%s %d %f %t", name, age, hight, isBoy)
	}
}

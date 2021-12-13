package main

import (
	"flag"
	"fmt"
)

func main() {
	// 定义常量
	var (
		VERSION string
	)
	flag.StringVar(&VERSION, "version", "1.0", "1.0")
	flag.Parse()
	flag.Usage()
	fmt.Printf("version=%s\n", VERSION)

}

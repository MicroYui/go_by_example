package main

import (
	"fmt"
	"os"
)

func main() {

	// os.Args 提供了对原始命令行参数的访问。 请注意，片段中的第一个值是程序的路径，os.Args[1:] 包含程序的参数。
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}

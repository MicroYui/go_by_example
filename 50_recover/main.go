package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

func main() {
	// 在这个匿名函数中，调用 recover() 函数来捕获恐慌。recover() 会阻止程序的恐慌并返回恐慌时传入的值，
	// 如果没有发生恐慌，recover() 返回 nil。这里的代码块会在捕获到恐慌时执行，从而打印出错误信息。
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	fmt.Println("After mayPanic()")
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	currentDir, err := os.Getwd()
	check(err)
	fmt.Println(currentDir)

	filePath := filepath.Join(currentDir, "tmp", "dat")

	dat, err := os.ReadFile(filePath)
	check(err)
	fmt.Print(string(dat))

	f, err := os.Open(filePath)
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	o2, err := f.Seek(6, io.SeekStart)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	_, err = f.Seek(4, io.SeekCurrent)
	check(err)

	_, err = f.Seek(-10, io.SeekEnd)
	check(err)

	o3, err := f.Seek(6, io.SeekStart)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, io.SeekStart)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	for {
		line, isPrefix, err := r4.ReadLine()
		if err != nil {
			if err == io.EOF {
				break // 读取到文件末尾，退出循环
			}
			fmt.Println("读取失败:", err)
			return
		}

		// 处理读取的行
		fmt.Printf("读取的行: %s\n", line)

		// 如果 isPrefix 为 true，表示行长度超过了缓存大小，内容被截断了
		if isPrefix {
			fmt.Println("行长度超过了读取缓冲区的大小，内容被截断")
		}
	}
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close()
}

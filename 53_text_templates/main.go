package main

import (
	"os"
	"text/template"
)

// 这段代码使用了 Go 语言的 text/template 包来创建和执行文本模板。模板是一种文本格式，用于动态生成输出。
func main() {
	// 创建一个新的模板 t1，模板名为 "t1"
	t1 := template.New("t1")
	// 解析模板字符串 "Value is {{.}}\n"。{{.}} 是一个占位符，表示传递给模板的数据
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	// 帮助函数，确保模板解析成功。如果解析失败则会触发恐慌（panic）。
	// 这里将另外一个字符串 "Value: {{.}}\n" 添加到模板中。值得注意的是，这会覆盖之前的模板内容。
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	// 执行模板 t1，将 "some text" 传递给模板，输出结果为 Value: some text
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	// 定义了一个方便的函数 Create，用来创建命名的模板（通过 name）并解析模板字符串（通过 t）。同样使用 template.Must 来确保解析成功
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	t2 := Create("t2", "Name: {{.Name}}\n")

	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})

	// 创建了模板 t3，通过条件语句检查传入的值。
	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	// 创建模板 t4，用于遍历切片或数组。
	t4 := Create("t4",
		"Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout,
		[]string{
			"Go",
			"Rust",
			"C++",
			"C#",
		})
}

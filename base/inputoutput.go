package base

import (
	"bufio"
	"fmt"
	"os"
)

func Output() {
	// method 1
	os.Stdin.WriteString("hello world~")

	// method 2
	println("hello world")

	// method 3
	fmt.Println("hello world")
}

func Input() {
	//// 扫描从os.Stdin读入的文本，根据空格分隔，换行也被当作空格
	//func Scan(a ...any) (n int, err error)
	//
	//// 与Scan类似，但是遇到换行停止扫描
	//func Scanln(a ...any) (n int, err error)
	//
	//// 根据格式化的字符串扫描
	//func Scanf(format string, a ...any) (n int, err error)

	//var s, s2 string
	//fmt.Scanln(&s, &s2)
	//fmt.Println(s, s2)

	// buffer
	// read
	//scanner := bufio.NewScanner(os.Stdin)
	//scanner.Scan()
	//fmt.Println(scanner.Text())

	// write
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString("hello write")
	writer.Flush()
	fmt.Println(writer.Buffered())
}

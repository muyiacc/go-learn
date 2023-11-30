package base

import (
	"fmt"
	"os"
)

func Condition() {
	// if else
	a := 100
	b := 20
	if a > b {
		fmt.Println("hello muyiacc")
	}

	if a > b {
		fmt.Println("a > b")
	} else if a == b {
		fmt.Println("a = b")
	} else {
		fmt.Println("a < b")
	}

	// switch
	switch a {
	case 10:
		fmt.Println(a + 1)
	case 100:
		fmt.Println(a + 2)
		// break	// 跳出
		fallthrough // 执行完成，继续执行下一个分支
	case 1000:
		fmt.Println(a + 1000)
	default:
		fmt.Println(a + 3)
	}

	// label: 单独使用没有意义，一般配合使用goto,break,continue
	// goto: 跳转到另一个标签,goto不会正常自动退出
A:
	aaa := 10
	os.Exit(0)

	if aaa == 10 {
		goto A

	} else {
		fmt.Println(aaa)
	}

}

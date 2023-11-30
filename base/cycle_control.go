package base

import "fmt"

func Cycle() {
	// for： go中只有这一种循环，没有while
	fmt.Println("===for===")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 当成 while用
	fmt.Println("===while===")
	num := 1
	for num < 5 {
		num++
		fmt.Println(num)
	}

	// multiplication table
	for i := 1; i < 10; i++ {
		for j := i; j < 10; j++ {
			if i <= j {
				fmt.Printf("%d*%d=%2d  ", i, j, i*j)
			}
		}
		fmt.Println()
	}

	// for ... range: 更加方便的遍历数组，切片，字符串等
	sequence := "hello world"
	for index, value := range sequence {
		fmt.Println(index, value)
	}

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for key, value := range arr {
		fmt.Println(key, value)
	}

	// break:终止最内层循环

	// continue: 跳过最内层的本次循环
}

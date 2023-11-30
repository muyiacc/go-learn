package base

import "fmt"

func Loop() {

	// for i := 1; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	// var i = 1
	// for i <= 3 {
	// 	// fmt.Println("hello muyiacc")
	// 	i++
	// }

	arr := [3]int{122, 232, 353}
	for index, value := range arr {
		fmt.Println(index, value)
	}
}

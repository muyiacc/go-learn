package base

import (
	"cmp"
	"fmt"
)

/*
Varable
*/
func Varable() {

	// declaration method
	// var variable_name type
	var intNum int
	var str string
	var char byte
	fmt.Println(intNum)
	fmt.Println(str)
	fmt.Println(char)

	// the same types can be together
	var num1, num2, num3 int
	var num4, num5, num6 = 1, 2, 3
	fmt.Println(num1, num2, num3)
	fmt.Println(num4, num5, num6)

	// different types can be wrapped whit ()
	var (
		name string
		age  int
	)
	fmt.Println(name, age)

	// Short variable initialization: can omit the var keyword and the descendant type,
	// the type is left to compiler to infer on its own
	address := "earth"
	gender := false
	fmt.Println("address====>" + address)
	fmt.Println(gender)

	// but can't init nil, nil does not belong to any type

	// short variable declarations can also be bulk initialised
	a, b, c := "a", 123, 3.14
	fmt.Println(a, b, c)

	// 匿名变量
	a1, _ := 1, 2
	fmt.Println(a1)

	// swapping the values of two variables is straightforward with =
	swap1, swap2 := 1, 2
	fmt.Println(swap1, swap2)
	swap2, swap1 = swap1, swap2
	fmt.Println(swap1, swap2)

	// since the existence of unused inside a function will fail compilation, but  som variables are really not used,
	// you can use the anonymous variable _ to indicate that the variable can be ignored.
	ana, anb, _ := 1, 2, 3
	fmt.Println(ana, anb)

	/*
		compare
		since 1.21
		Added the min, max , and cmp packages for comparing variable values, all with generic support.

		func max[T cmp.Ordered](x T, y ...T) T
		func min[T cmp.Ordered](x T, y ...T) T
	*/
	maxValue := max(24, 5, 5, 78, -32, 53, 46)
	minValue := min(24, 5, 5, 78, -32, 53, 46)
	fmt.Println(maxValue)
	fmt.Println(minValue)

	// compare two values, equal to return 0 ,greater than return 1, less than return -1
	switch cmp.Compare(-1, 2) {
	case 0:
		fmt.Println(0)
	case 1:
		fmt.Println(1)
	case -1:
		fmt.Println(-1)
	}

	// datemine if it is less than
	less := cmp.Less(10, 20)
	fmt.Println(less)

}

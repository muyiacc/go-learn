package base

import "fmt"

func Constant() {
	// declare: must be initialised
	const msg = "It's a message"
	// msg = "msg" // Cannot assign to msg

	// const cmsg
	// fmt.Println(cmsg) 	// missing init expr for cmsg

	// bath declaration
	const (
		msg1 = "hello"
		msg2 = 123
		msg3 = true
	)

	// iota:usually used to denote an untyped integer ordinal a constant declaration,usually in parentheses
	const (
		num = iota
		num1
		num2
		num3
	)

	const (
		numn1   = iota * 2 // 0
		numn2              // 2
		numn3              // 4
		numn4   = 10
		numn5   = iota // 4
		numn6          // 5
		numn100        // 6
	)

	const (
		a = 1 // 1
		b     // 1
		c     // 1
		d     // 1

	)

	fmt.Println(num, num1, num2, num3) // 0 1 2 3 4 5
	// num = 10                                       // can't assign to num

	const aaa = 10
}

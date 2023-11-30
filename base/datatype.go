package base

import "fmt"

/*
data type set
*/
func Datatype() {
	// golang中的数据类型
	// refer site : https://golang.halfiisland.com/%E8%AF%AD%E8%A8%80%E5%85%A5%E9%97%A8/%E8%AF%AD%E6%B3%95%E5%9F%BA%E7%A1%80/10.datatype.html#%E6%B5%AE%E7%82%B9%E5%9E%8B

	// bool
	var abool bool
	fmt.Println("=====bool=====")
	fmt.Println(abool) // flase

	// integer
	var aint8 int8       // signed integer
	var auint8 uint8     // unsigned integer
	var aint int         // least 32bit signed integer
	var auint uint       // least 32bit unsigned integer
	var auintptr uintptr // a large enough integer
	fmt.Println("=====integer=====")
	fmt.Println(aint8, auint8, aint, auint, auintptr)

	// float
	var afloat32 float32 // 32bit float
	var afloat64 float64 // 64bit float
	fmt.Println("=====float=====")
	fmt.Println(afloat32, afloat64)

	afloat32 = 10.12345678912345
	afloat64 = 10.12345678912345
	fmt.Println(afloat32) // 10.123457
	fmt.Println(afloat64) // 10.12345678912345

	// complex
	var acomplex128 complex128 // 64的实数和虚数
	var acomplex64 complex64   // 32的实数和虚数
	fmt.Println("=====complex=====")
	fmt.Println(acomplex64, acomplex128)
	acomplex64 = 12 + 123456789i
	acomplex128 = 120 + 123456789i
	fmt.Println(acomplex64)  // (12+1.2345679e+08i)
	fmt.Println(acomplex128) // (120+1.23456789e+08i)

	// string
	var abyte byte     // equivalence unit8, can express ascii characters
	var arune rune     // equivalence int32, can express unicode characters
	var astring string // A string is a sequence of bytes, which can be converted to a byte type of [], i.e. a byte slice
	fmt.Println("=====string=====")
	fmt.Println(abyte, arune, astring)

	abyte = 123
	arune = 123456
	astring = "123456"
	fmt.Println(abyte)   // 123
	fmt.Println(arune)   // 123456
	fmt.Println(astring) // 123456

	// derive
	/*
		类型	例子
		数组	[5]int，长度为5的整型数组
		切片	[]float64，64位浮点数切片
		映射表	map[string]int，键为字符串类型，值为整型的映射表
		结构体	type Gopher struct{}，Gopher结构体
		指针	*int，整型指针
		函数	func()，一个没有参数，没有返回值的函数类型
		接口	type Gopher interface{}，Gopher接口
		通道	chan int，整型通道
	*/
	fmt.Println("=====derive=====")
	var aarray [3]int
	fmt.Println(aarray)

	aarray[0] = 1
	aarray[1] = 2
	aarray[2] = 3
	//aarray[3] = 3 	// failed to compile
	fmt.Println(aarray)

	var barray = [3]string{"hello", "world", "golang"}
	fmt.Println(barray)

	// zero value
	/*
		类型	零值
		数字类型	0
		布尔类型	false
		字符串类型	""
		数组	固定长度的对应类型的零值集合
		结构体	内部字段都是零值的结构体
		切片，映射表，函数，接口，通道，指针	nil
	*/

	// nil
	/*
		Go中的nil并不等同于其他语言的null，nil仅仅只是一些类型的零值，并且不属于任何类型，所以nil == nil这样的语句是无法通过编译的。
	*/

}

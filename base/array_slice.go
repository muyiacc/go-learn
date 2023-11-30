package base

import (
	"fmt"
)

func Array() {

	// 声明
	// go中数组是值类型，不是指向头部的指针
	// 初始化一个长度为5的整型数组
	// var arr1 [5]int

	// 元素初始化
	// arr2 := [5]int{1, 2, 3, 4, 5}

	// new函数获得一个指针
	// arr3 := new([5]int)

	// 初始化的长度必须是一个常量表达式
	// a := 5
	// arr4 := [a]int	// Invalid array bound 'a', must be a constant expression

	// 访问数组
	nums := [5]int{1, 2, 3, 4, 5}
	fmt.Println(nums[0])

	nums[0] = 0
	fmt.Println(nums[0])

	// 查看长度
	alen := len(nums)
	fmt.Println(alen)

	// 查看容量
	acap := cap(nums)
	fmt.Println(acap)

	// 切割，左闭右开
	anums1 := nums[:]
	anums2 := nums[1:]
	anums3 := nums[1:3]
	fmt.Println(anums1)
	fmt.Println(anums2)
	fmt.Println(anums3)
}

func Slice() {
	// 切片，可以理解为一个可以变动容量的数组，底层实现依旧是数组，是引用类型

	// 声明，初始化切片
	var nums1 []int // 默认值 nil
	nums2 := []int{1, 2, 3}
	nums3 := make([]int, 0, 0)
	nums4 := new([]int)

	fmt.Println("===切片===")
	fmt.Println("空切片 ==>", nums1)
	fmt.Println("初始化 ==>", nums2)
	fmt.Println("make建造 ==>", nums3)
	fmt.Println("new一个指针 ==>", nums4)

	// 对比初始化数组
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("===数组===")
	fmt.Println(arr)

	// 事实上，因为不知道事先的长度，所以切片在go中运用的更广一些，且后续使用过程中可能会频繁的插入和删除元素。

	// 使用
	// make的参数：类型，长度，容量。	长度代表着元素的个数，容量代表着最大存放元素的个数
	aslice := make([]int, 0, 0) // 创建一个切片

	fmt.Println("==切片的增删查改==")

	// 遍历切片，查看切片的内容

	fmt.Println(aslice[0])

	// 增加一个新的元素
	aslice = append(aslice, 1, 2, 3, 4, 5, 6, 7)

	fmt.Println(aslice, "")
	fmt.Println(len(aslice), cap(aslice)) // cap查看容量，len()查看长度

	// 插入元素
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// 从头部插入元素
	nums = append([]int{-1, 0}, nums...)
	fmt.Println(nums)
}

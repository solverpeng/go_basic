package main

import (
	"fmt"
	"strconv"
)

type (
	newInt int
)

func main() {
	// 自定义类型的转换也需要显式的转换
	var n int
	var m newInt = 1
	n = int(m)
	fmt.Println(n)

	// byte和uint8之间不需要显式转换，byte是uint8的类型别名
	var i uint8
	var j byte = 1
	i = j
	fmt.Println(i)

	// rune和int32之间不需要显式类型转换，rune是int32的类型别名
	var x int32
	var y rune = 22
	x = y
	fmt.Println(x)

	// 在相互兼容的两种类型之间进行转换
	var a float32 = 1.1
	b := int(a)
	fmt.Println(b)

	// 无法通过编译
	// var c bool = true
	// d := int(c)

	var f int = 65
	g := string(f)
	// 输出A
	fmt.Println(g)

	// Itoa is shorthand for FormatInt(int64(i), 10).
	g = strconv.Itoa(f)
	fmt.Println(g)

	// Atoi returns the result of ParseInt(s, 10, 0) converted to type int.
	f, _ = strconv.Atoi(g)
	fmt.Println(f)

}

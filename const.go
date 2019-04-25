package main

import "fmt"

// 常量的定义方式
const a int = 1

const b = 'A'

const (
	text   = "123"
	length = len(text)
	num    = b * 20
)

const i, j, k = 1, "2", '3'

const (
	text2, length2, num2 = "456", len(text2), k * 10
)

// 常量组，不提供默认值，则使用上行的表达式
const (
	aa = 'A'
	bb
)

//iota是常量计数器，从0开始，组中每定义一个常量自动递增1
const (
	cc = iota
	dd
)

// 每遇到一个const关键字，iota就会重置为0
const (
	ee = 1
	ff = iota
	gg
)

// 通过初始化规则与iota可以达到枚举的效果
// 星期枚举
const (
	// 第一个表达式不可省略
	MONDAY = iota
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
	SUNDAY
)

func main() {
	fmt.Println(aa) //65
	fmt.Println(bb) //65

	fmt.Println(cc) //0
	fmt.Println(dd) //1

	fmt.Println(ee) //1
	fmt.Println(ff) //1
	fmt.Println(gg) //2

}

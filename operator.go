package main

import "fmt"

/**
运算符
----------------
6： 0110
11: 1011
*/
func main() {
	// 一元运算符，优先级最高
	fmt.Println(^6)    //-7    按位取反，一个有符号位的 ^ 操作为 这个数+1的相反数
	fmt.Println(!true) //false 取反，只能作用与bool类型

	// 二元运算符，高优先级，优先级从高到低
	fmt.Println(6 * 11)     //乘
	fmt.Println(6 / 11)     //除
	fmt.Println(6 % 11)     //取余
	fmt.Println(2 << 10)    //2048 左移，转换为二进制后，各位向左移10
	fmt.Println(2048 >> 10) //2 右移，转换为二进制后，各位向后移10位，不足的部分舍弃
	fmt.Println(6 & 11)     //2 与，AND，二进制对应位，都为1，才为1
	fmt.Println(6 &^ 11)    //4 0100 位清空，AND NOT，若11 bit位上的数是0则结果位取6上对应位置的值，若bit位上为1则结果位取0
	fmt.Println(6 + 11)     //加
	fmt.Println(6 - 11)     //减
	fmt.Println(6 | 11)     //15 或，OR，转换为二进制后，x/y只要有一个1，结果位则取1
	fmt.Println(6 ^ 11)     //13 异或，转换为二进制后，x/y，对应位相反的情况下，结果位取1

	// 二元运算符，底优先级
	fmt.Println(2 == 2)
	fmt.Println(2 != 2)
	fmt.Println(2 < 2)
	fmt.Println(2 <= 2)
	fmt.Println(2 >= 2)
	fmt.Println(2 > 2)

	// && 和 ||，用于判断
	a := 1
	if a > 0 && 10/a > 1 {
		fmt.Println("ok")
	}

}

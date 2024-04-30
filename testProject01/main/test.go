package main

import (
	"fmt"
	"testProject01/utils"
)

func cal(num1 *int, num2 *int) int {
	return *num1 + *num2
}

func main() {
	fmt.Println("hello world")
	var sex = 12321
	fmt.Println(sex)
	//var n2 float32 = 3.14
	//var zhen bool = true
	fmt.Printf("%d\n", sex)
	var age = 18
	var ptr *int = &age
	fmt.Println(ptr)
	var wuhu int
	var err error
	// 键盘录入信息
	wuhu, err = fmt.Scanln(&age)
	if err != nil {
		fmt.Println("读取输入时出错:", err)
	}
	fmt.Println(wuhu, age)
	var sum int
	for i := 1; i <= 100; i++ {
		sum += i
		if sum > 100 {
			break
		}
	}
	fmt.Println(sum)
	utils.Plant()
}

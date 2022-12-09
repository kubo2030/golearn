package main

import "fmt"

func main() {
	/*
		遍历数组：依次获取 数组中的数据
		range 数组名：
		index ,value
	*/
	arr := [...]int{6, 2, 4, 9, 8, 3}
	// 1. for 遍历数组
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], "\t")
	}
	fmt.Println()

	// 2. rang 遍历数组
	sum := 0
	for _, value := range arr {
		fmt.Println(value)
		sum += value
	}
	fmt.Println(sum)

}

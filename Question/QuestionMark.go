package Question

import "fmt"

/**
这个 用法 定义了 题目 确认请求等
*/

func Question_init() int {
	fmt.Println("Do It!")
	fmt.Println("你是中国人吗")
	var chinese string
	fmt.Scanln(&chinese)
	if chinese == "是" {
		println("OK,初始化成功")
		return 1
	} else if chinese == "否" {
		println("你干嘛~")
		return 1
	} else {
		println("fuck you, please go out this app!")
		return 0
	}
}

func Question() int {
	var nit int = Question_init()
	if nit == 1 {
		println("中国人，你好")

		//Q1

		//Q2

		//Q3

		//Q4

		//Q5

		//Q6

		//Q7

		//Q8

		//Q9

		//Q10

	} else {
		println("滚")
		return 0
	}
	return 0
}

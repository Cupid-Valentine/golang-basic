package main

import "fmt"

// 函数外变量 不是全局变量  是包内部变量
var aa = 3
var ss = "kkk"
//函数外面不能这么写
//bb := true
var bb = true

//=======

//简写
var (
	aa1 = 3
	ss1 = "kkk"
	bb1 = true
)

func variable() {
	//变量名在前  类型在后
	//定义完成 就有初值
	var a int
	var s string
	//空字符串打印不出来
	fmt.Println(a, s)
	fmt.Printf("%d  %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}
func main() {
	variable()
	fmt.Println("hello world")
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)
}

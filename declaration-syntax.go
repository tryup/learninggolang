package main

import "fmt"
//type decl 参考https://blog.golang.org/gos-declaration-syntax学习
//https://golang.org/ref/spec
//https://tip.golang.org/ref/spec

func main(){

	var a int  //声明变量int a
	fmt.Println(a) //如果声明的时候没有初始化，则默认为类型的0值,详见zero-value
	a = 10  //赋值10给变量a
	fmt.Println(a)

	var a1,a2,a3 int //声明3个int变量
	a1,a2,a3 = 1,2,3
	fmt.Println(a1,a2,a3)

	var (
		a4 =4
		a5 = 5
	)

	fmt.Println(a4,a5)
	b:=11   //:= 声明和定义变量b 自动推导变量的类型是int
	fmt.Println(b)

	var c int = 12  //声明和定义变量int c
	fmt.Println(c)

	//可以总结出语法是var var_name [var_type]
	var pc *int = &c //声明定义int指针类型 pc
	fmt.Println(*pc) //所以打印pc内容需要解引用,有C/C++基础可以很容易理解。

	//那么可以很容易想到数组
	var pcs [10]*int //pcs指针数组数组len=10
	pcs[0] = &c
	*pcs[0] = 122
	fmt.Println(len(pcs))
	fmt.Println(*pcs[0])
	fmt.Println(c)

}
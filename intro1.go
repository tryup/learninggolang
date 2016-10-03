package main  //包名是main，每一个源文件必然属于一个包，详见后面对go包管理的分析

import "fmt"  //导入fmt包,如果导入的包不使用,则会编译错误。

import (
	"os"   // 另外一种语法syntax,用来批量导入包
	"math"
)

func firstfunc(lhs,rhs int)(rt1,rt2 int){ //函数定义 func关键词 函数名称(形参列表)(返回值列表) golang支持多返回值
	return lhs+rhs,lhs-rhs
}

//main类似C++main 作为程序的开始的默认函数名
func main() { //左括号匹配不能换行

	a,b := firstfunc(1,2)
	fmt.Println( "a:",a, "b:",b)
	fmt.Println("hi i am fobnn")
	fmt.Println( math.Sqrt(10))
	fmt.Println(os.Args)
}

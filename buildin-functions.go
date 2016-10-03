package main

import "fmt"
func slicemake(){
	var b []int
	b = make([]int,1) //构造长度和容量为1的切
	fmt.Println(b[0])
	fmt.Println(b[1]) //越界

	a := make([]int,1,5)

	fmt.Println(cap(a)) //调用内置函数cap获取切片的容量
	fmt.Println(a[0])
	fmt.Println(a[1]) //越界
	array := []int{1,2}

	a = array
	fmt.Println(a[1])

	make([]int,2,1)//len larger than cap in make([]int)
	make([]int,-1,-2)//negative len argument in make([]int)
	make([]int,0.5) //constant 0.5 truncated to integer 0

}
func main(){
	slicemake()
}

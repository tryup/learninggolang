##GoLang的预定义内置函数
从下面链接的预定一符号中的Functions中得知
https://tip.golang.org/ref/spec#Predeclared_identifiers
>+ append
>+ cap
>+ close
>+ complex
>+ copy
>+ delete
>+ imag
>+ len
>+ make
>+ new
>+ panic
>+ print
>+ println
>+ real
>+ recover

###make
>https://tip.golang.org/ref/spec#Making_slices_maps_and_channels

有点类似C++中的make_xxxx构造器,当然Golang中的实现肯定不同，通过make函数可以构造
切片，Map和channel。构造分为两个步骤，1分配内存，2初始化，跟C++很像，初始化为T的零值。
>```Call             Type T     Result
>   make(T, n)       slice      slice of type T with length n and capacity n
>   make(T, n, m)    slice      slice of type T with length n and capacity m
>  
>   make(T)          map        map of type T
>   make(T, n)       map        map of type T with initial space for n elements
>   
>   make(T)          channel    unbuffered channel of type T
>   make(T, n)       channel    buffered channel of type T, buffer size n

另外需要注意的是n和m参数必须是一个整数类型或者是无类型类型。
A constant size argument must be non-negative and representable by a value of type int. If both n and m are provided and are constant, then n must be no larger than m. If n is negative or larger than m at run time, a run-time panic occurs. 
>```
>s := make([]int, 10, 100)       // slice with len(s) == 10, cap(s) == 100
>s := make([]int, 1e3)           // slice with len(s) == cap(s) == 1000
>s := make([]int, 1<<63)         // illegal: len(s) is not representable by a value of type int
>s := make([]int, 10, 0)         // illegal: len(s) > cap(s)
>c := make(chan int, 10)         // channel with a buffer size of 10
>m := make(map[string]int, 100)  // map with initial space for 100 elements

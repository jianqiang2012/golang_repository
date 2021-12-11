package main

/*
1.变量
	Go语言的变量声明格式为：

    var 变量名 变量类型

	 var name string
同一个数据类型的变量可以省略保留以后一个
    var age,sum,sub,de int


批量申明

 var (
        a string
        b int
        c bool
        d float32
    )

2.初始化



    var 变量名 类型 = 表达式
	举个例子：

    var name string = "pprof.cn"
    var sex int = 1
	或者一次初始化多个变量

    var name, sex = "pprof.cn", 1

不想写类型的 可以上下文推导

	var name = "pprof.cn" // go 自己就知道是 string了

    var sex = 1  // go自己就知道是 int

还可以用冒号等号的局部申明方式，这是个骚操作


func main() {
    n := 10
    m := 200 // 此处声明局部变量m
    fmt.Println(m, n)
}


还有更大胆的匿名变量，我不想要的就用下划线

 x, _ := foo()
==========================
该收尾了 总结一下
 1  .函数外的每个语句都必须以关键字开始（var、const、func等）
 2.	:=不能使用在函数外
 3._多用于占位，表示忽略值


3.常量

const pi = 3.1415

多个

const (
a = 3.14159265
b = 10^8
)


你还可以这样创建变量 name，age := "golang",18
*/

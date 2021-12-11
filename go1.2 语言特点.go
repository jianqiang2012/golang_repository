package main

/*
1. Go 有这些特点
	1.自动立即回收。	// 有指针也有垃圾回收机制 （Java，python只能自动回收，coder不能直接操作内存）
    2.更丰富的内置类型。	// 除了内置的，还可以自定义类型 比如 interface{} 表示任意类型的
    3.函数多返回值。	// 函数返回值可以为空或者多个返回
    4.错误处理。  // 自定义的 panic
    5.匿名函数和闭包。  // 支持匿名函数和闭包
    6.类型和接口。  // 如上面的 interface就是接口
    7.并发编程。 //  通过chan来实现并发
    8.反射。	// 支持reflect
    9.语言交互性。 //支持和其他语言交互 python和公go的交互： https://www.cnblogs.com/dhcn/p/12044521.html
2. 源码标记 ： 都是以 '.go'结尾

3.Go的函数、变量、常量、自定义类型、包的命名规则：

    1.首字符可以是任意的Unicode字符或者下划线 // 数字开头不行 Unicode ～～unit8 ~~ [a-zA-Z0-9等]
    2.剩余字符可以是Unicode字符、下划线、数字
    3.字符长度不限制，看你心情吧
4.代码作用范围
	1.声明在函数内部，是受保护的，代码块内可以起作用，类似private
    2.声明在函数外部，是对当前包可见(包内所有.go文件都可见)的全局值，类似protect  //有点像 python的 local 但不完全
    3.声明在函数外部且首字母大写是所有包可见的全局值,类比Java的public，python的 gobal 就是说其他文件可以调用
5.申明方式
	var（声明变量，普通常量，指针等）
	const（声明常量，const x int  = 8）
	type（声明类型，struct，interface）
	func（声明函数，方法和 函数都用）
6.项目
	src：源代码文件
    pkg：包文件
    bin：相关bin文件
    <待补充>
7.go的编译


	<待补充>

8. 数据类型与值传递
	slice   -- 序列数组(最常用)
    map     -- 映射
    chan    -- 管道
	除了以上三个是指针类型可以引用传递，其余的类型都是值传递
9.常用的内置函数和简介
	 append          -- 用来追加元素到array、slice中,返回修改后的数组、slice
    close           -- 主要用来关闭channel（chan）
    delete            -- 从map中删除key对应的value; delete(变量，key)
    panic            -- 停止常规的goroutine  （panic和recover：用来做错误处理）
    recover         -- 允许程序定义goroutine的panic动作
    real            -- 返回complex的实部   （complex、real imag：用于创建和操作复数）
    imag            -- 返回complex的虚部
    make            -- 用来分配内存，返回Type本身(只能应用于slice, map, channel)
    new                -- 用来分配内存，主要用来分配值类型，比如int、struct。返回指向Type的指针
    cap                -- capacity是容量的意思，用于返回某个类型的最大容量（只能用于切片和 map）
    copy            -- 用于复制和连接slice，返回复制的数目
    len                -- 来求长度，比如string、array、slice、map、channel ，返回长度
    print、println     -- 底层打印函数，在部署环境中建议使用 fmt 包

10.接口
	  type error interface { //我们需要实现了Error()方法，返回值为String的都实现了error接口

            Error()    String

    		}

11.init函数
	1 init函数是用于程序执行前做包的初始化的函数，比如初始化包里的变量等
    2 每个包可以拥有多个init函数
    3 包的每个源文件也可以拥有多个init函数
    4 同一个包中多个init函数的执行顺序go语言没有明确的定义（自己尝试）
    5 不同包的init函数按照包导入的依赖关系决定该初始化函数的执行顺序
    6 init函数不能被其他函数调用，而是在main函数执行之前，自动被调用

12. 函数执行的总入口
 	// 一切Go语言程序的默认入口函数(主函数)：func main()
    函数体用｛｝一对括号包起来

    func main(){
        //函数体，自己写的
    }
13.比较 init和 main
	相同点：
        两个函数在定义时不能有任何的参数和返回值，且Go程序自动调用
    不同点：
        init可以应用于任意包中，且可以重复定义多个
        main函数只能用于main包中，且只能定义一个



*/

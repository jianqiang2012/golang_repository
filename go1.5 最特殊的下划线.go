package main
/*
下回线是用来忽略某些东西的

1  忽略import

	package main

import _ "./hello"  // ----不执行这个包的函数 ，只是加载一下这个包的 init函数-----

func main() {
    // hello.Print()
    //编译报错：./main.go:6:5: undefined: hello
}


2.句子当中

	----下划线意思是忽略这个变量，说白了就是用来占位的，不占位就报错了----

    比如os.Open，返回值为*os.File，error

    普通写法是f,err := os.Open("xxxxxxx")

    如果此时不需要知道返回的错误值

    就可以用f, _ := os.Open("xxxxxx")

    如此则忽略了error变量

*/

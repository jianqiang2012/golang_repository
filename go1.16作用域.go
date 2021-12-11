package main

import (
    "fmt"
)
// 同一级的作用域之间的变量是不可以共享的 这就是go 的静态作用域，shell的话可以访问同级别的作用域
func main()  {
    var name  string = "Python编程"
    fmt.Println("main 里面name",name )
    func01()
}
func func01()  {
    fmt.Println("func01的 name",name) // 直接报错 undefined: name


}
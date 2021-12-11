package main

import "fmt"

func set_data(x int)  {
 // recover 只能在defer里捕获异常信息
    defer func() {
        if err:= recover();
            err != nil{
            fmt.Println(err)}
    }()
    // 这里制造一个索引越界的危险
    var arr = [10] int{}

    arr[x] = 99
}

func main()  {
    set_data(20)
    fmt.Println("index  out of range..")

}

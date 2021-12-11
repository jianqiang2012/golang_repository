package main

import (
    "fmt"
)
// 这里划重点了 select的case执行不是顺序的而是随机的
//  select 只能用于信道
/*select{ case 1  代码块
        case 2  代码块
        case 3  代码块
        default 代码块
}   */

//举例子

func main()  {

    c1:= make(chan string ,1)
    c2 := make(chan string,2)
    c2 <- "hello"

   //  go makeTimeout(timeout,1)  当下面的select阻塞的时候我们设置阻塞时间
    select {
    case msg1 := <- c1:
        fmt.Println("c1 received",msg1)
    case msg2:= <- c2:
        fmt.Println("c2 recevied",msg2)

    default: // 这里必须写 避免死锁 deadlock
        fmt.Println("no data received.")


    }

}

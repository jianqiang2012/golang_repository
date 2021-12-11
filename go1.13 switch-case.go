package main

switch 表达式 {
    case  条件1:
        //代码块
    case 条件2:
        //代码块
    default:
        代码块
}


注意1  case 后的常量不能重复 就像python的 key 一样
注意2  switch case  中一般是满足一个case 就退出了，如果都不满足就执行 default 但是也可以穿透 fallthrough (表示和他最近的一个case被跳过，直接执行里面的代码块)
注意3   case 就可以跟函数 但是要注意函数的返回值类型必须和case保持一致
注意4   case 可以没有任何变量 表达式 函数

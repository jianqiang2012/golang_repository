
. 创建指针
	a :=1
	pa = &a // pa 就是个指针

	第二种方法
	pr := new(int)
	*pr = 1

	第三种方法
	a :=1
	var  pa *int
	pa = &a

2.指针申明后没有初始化的时候 值默认就是 nil //


3. 举例子说明 改变数组的值的方法  切片发和指针法
// 切片法
func modify(params []int)  {
    params[0] = 1000
} //

func main()  {
    a := [3]int{1,2,4}
    modify(a[:])
}  //


//指针法

func modify(pr *[3]int)  {
    pr[0] = 1000
} //

func main()  {
    a:= [3]int{1,2,3}
    modify(&a)

}




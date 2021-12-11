



//可以容纳若干元素但是不可以固定长度，切片的底层数据结构就是数组，所以切片是引用传递

arr := [...]int{1,2,3,4,5}
myslice:= arr[0:1]

切片还可以指定容量大小

myslice2:=arr[1:3]  // 输出的len= 2 cap = 4 就是从arr1,到arr[最后一个元素] 直接的个数[2-5]刚好就是4
myslice3:= arr[1:3:4] // 输出len= 2,cap =3  就是 2，3，4这三个计数就是3
这里的cap 只是切片的Cap不影响数组的cap


.直接申明切片

	var myslice4 []string
	var myslice5 []int
	var myslice6 = []int{} // 空切片

.使用make构造
	a:= make([]int,2,10)
	b:= make([]int,3) //

.更懒的方式
	a:=[]int{4:2} // len = 5 a= [0 0 0 0 2] cap = 5

关于len()和cap() 和 a 的区别

	a  就是公司
	cap 就是公司的工位
	len 就是当前公司有多少人

	由于切片是应用类型 所以当没有赋值的时候,他的默认值就是 nil

	这里我们比较一下数组和切片的异同
	数组容量固定 元素类型相同 切片容量不固定,元素类型相同 切片其实就是python的 list 我们可以使用append添加元素，使用delete删除元素

	my := []int{1}
	append(my,2) // 加一个2
	append(my,2,3) // 加多个 2，3
	append9(my,[]{7,8}...) // 加一个切片（... 是解包）






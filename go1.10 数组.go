

"golng 的数组是有固定长度的（和切片slice的区别也是这里）,可以是多个或者零个,长度的固定 ，这也导致了 GO中直接使用数组的情况不会很多"
 . 数组的申明
	var arr [3]int
 . 赋值如下
	arr[0]=1
	arr[1]=2
	arr[2]=3


	"还可以这样写"

	var arr [3]int = [3]int{1,2,3} // 注意等号左边的【3】int  就是一种数据类型
	arr := [3]int{1,2,3}

	"你可以来一个更厉害的操作"

	arr := [...]int{1,2,3} // ... 表示让go  自己去分配内存空间

	注意是:   [3]int 和 [4]int  都是数组，但是类型不同 可以用 %T 输出验证

	"""
	func main()  {
		arr01:= [3]int{1,2,3}
		arr02:= [4]int{1,2,3}
		fmt.Printf("%T--%T",arr01,arr02)

		}

	"""




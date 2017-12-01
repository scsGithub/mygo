package study

import (
	"fmt"
	"reflect" //
)

////结构 struct
// type person struct {
// 	Name string
// 	Age  int
// }
// type person struct {
// 	Name    string
// 	Age     int
// 	Contact struct {
// 		Phone, City string
// 	}
// }

/////method
//
//
//

////接口  interface
//
//
//

//反射 reflection
//
//
type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello() { //

	fmt.Println("我是接口拥有的方法 hello")
}

func Info(o interface{}) { //用于传进取一个空接口

	t := reflect.TypeOf(o) /////"reflect"包中的TypeOf 来获得接受到的接口类型
	fmt.Println("Type:", t.Name)

	v := reflect.ValueOf(o) /////"reflect"包中的ValueOf  来获得接受到的接口的字段
	fmt.Println("Fields:")

	for i := 0; i < t.NumField(); i++ { /////t.NumField()   t所拥有的字段数量

		f := t.Field(i)               ////根据 i取出具体字段
		val := v.Field(i).Interface() //////取出字段的值

		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)

	}
	//    for i := 0; < t.NumMethod(); i++ {////t.NumMethod();   取得方法的数量
	// 	m := t.Method(i) ///取得方法
	// 	fmt.Printf("%6s: %v\n",m.Name, m.Type)

	//    }

}

func Basis1() {

	// Test()

	////反射
	//
	//
	u := User{1, "xiaoming", 13}
	Info(u)

	//结构体 赋值 1
	// a := person{}
	// a.Name = "joe"
	// a.Age = 19
	// fmt.Println(a)
	// //结构体 赋值 2
	// a := person{
	// 	Name ： "joe"
	// 	Age ： 19
	// }
	// //结构体 赋值 3   在 person加取地址   方便调用操作
	// a := person{
	// 	Name ： "joe"
	// 	Age ： 19
	// }
	// a.Name = "ok"

	// fmt.Println(a)
	//结构体 赋值 4  匿名结构
	// a := struct {
	// 	Name string
	// 	Age  int
	// }{
	// 	Name: "joe",
	// 	Age:  19,
	// }
	// fmt.Println(a)

	// a := person{Name: "joe", Age: 19}
	// a.Contact.Phone = "1234567"
	// a.Contact.City = "beijing"
	// fmt.Println(a)

	//slice声明方法1
	// a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// fmt.Println(a)
	// // s1 := a[9] //取数组某位
	// s1 := a[5:10] //a[6 7 8 9 0]
	// fmt.Println(s1)

	//slice声明方法2 标准方式
	// s1 := make([]int, 3, 10) //参数1 类型   参数2 包含元素数  参数3 slice容量 初始容量 分配10个连续内存 当大于10时会大于一倍 重新分配一个20的连续内存块
	// fmt.Println(len(s1), cap(s1))  //长度  内存空间

	//slice获取底层数组的
	// a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	// sa := a[2:5]
	// fmt.Println(sa)
	// fmt.Println(string(sa))

	// //reslice   从一个slice中获取新的slice
	// a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	// sa := a[2:5]

	// fmt.Println(sa)
	// fmt.Println(string(sa))
	// fmt.Println(len(sa), cap(sa)) //sa长度是3 内存空间是从3开始 连续9个

	// sb := sa[2:9]
	// fmt.Println(sb)
	// fmt.Println(string(sb))

	// //Append //可以在slice尾部追加元素 将一个slice追加到另一个slice尾部
	// s1 := make([]int, 3, 6)

	// fmt.Printf("%p \n", s1)

	// s1 = append(s1, 1, 2, 3)
	// fmt.Printf("%v %p \n", s1, s1)

	// s1 = append(s1, 1, 2, 3)
	// fmt.Printf("%v %p \n", s1, s1)

	//map
	//类似字典  以key-value形式
	// var m map[int]string
	// m = make(map[int]string)
	// fmt.Println(m)

	// var m map[int]string = make(map[int]string)
	// fmt.Println(m)

	// m := make(map[int]string)
	// m[1] = "ok"
	// fmt.Println(m)
	// a := m[1]
	// fmt.Println(a)
	// delete(m, 1)
	// b := m[1]
	// fmt.Println(b)

	//复杂map
	// var m map[int]map[int]string
	// m = make(map[int]map[int]string)
	// m[1] = make(map[int]string)
	// m[1][1] = "ok"
	// a := m[1][1]
	// fmt.Println(a)

	//对键值进行检查
	// var m map[int]map[int]string
	// m = make(map[int]map[int]string)

	// a, ok := m[1][1] //判断键值对的存在
	// if !ok {
	// 	m[2] = make(map[int]string)
	// }
	// m[2][1] = "GOOD"
	// a, ok = m[2][1]
	// fmt.Println(a, ok)

	//对slice进行迭代  类似for in  i 是索引 int形  v 是slice的值得拷贝
	// for i,v:= range slice{
	//对 v 操作不改变 slice   改变slice 可以 slice[i] = ...
	// }
	//对map进行迭代 如上
	// for i,v:= range map{

	// }

	////不操作slice
	// sm := make([]map[int]string,5)
	// for _, v := range sm {
	// 	v = make(map[int]string,1)
	//     fmt.Println(v)
	// }
	// fmt.Println(sm)

	////操作slice
	// sm := make([]map[int]string, 5)
	// for i := range sm {
	// 	sm[i] = make(map[int]string, 1)
	// 	sm[i][1] = "ok"

	// 	fmt.Println(sm[i])
	// }
	// fmt.Println(sm)

	//map是无序的  对map间接排序
	// m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}

	// s := make([]int, len(m))
	// i := 0
	// for k, _ := range m {
	// 	s[i] = k
	// 	i++
	// }
	// fmt.Println(s)

	// sort.Ints(s)

	// fmt.Println(s)

	//
	//
	// m1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	// fmt.Println(m1)
	// m2 := make(map[string]int)

	// for k, v := range m1 {
	// 	m2[v] = k
	// }
	// fmt.Println(m2)

	//
	//
	//调用自定义函数
	//
	//  实验1
	// A("2", 3, 4, 5)
	//
	//
	//实验2 int传递
	// a, b := 1, 2
	// A(a, b)
	// fmt.Println(a, b)

	//
	//
	//实验3 slice 传递
	// s1 := []int{1, 2, 3, 4}
	// A(s1)
	// fmt.Println(s1)

	//
	//
	//指针传递

	//
	//
	//defer
	//实验一 逆序向上调用   打印顺序 a c b
	// fmt.Println("a")
	// defer fmt.Println("b")
	// defer fmt.Println("c")

	//实验2   打印  2 1 0
	// for i := 0; i < 3; i++ {
	// 	defer fmt.Println(i)
	// }

	//实验3
	// 打印 结果  3 3 3
	// for i := 0; i < 3; i++ {
	// 	defer func() {
	// 		fmt.Println(i)
	// 	}()
	// }

	//
	//实验4
	//defer   go没有异常机制 用panic/recover模式来处理错误

	//panic 在任何地方可引发   recover只能defer调用函数中生效

	// A()
	// B()
	// C()
}

//定义函数

// func  A(a int, b string,c int) (a int, b string,c int){

// 	a, b, c = 1, 2, 3
// 	return a, b, c
// }

////不定长变参
// func A(a ...int) {
// 	fmt.Println(a)
// }

// func A(b string, a ...int) {
// 	fmt.Println(b, a)
// }

//int形传递
// func A(s ...int) {
// 	s[0] = 3
// 	s[1] = 4
// 	fmt.Println(s)
// }

//slice传递
// func A(s []int) {
// 	s[0] = 3
// 	s[1] = 4
// 	s[2] = 8
// 	s[3] = 9
// 	fmt.Println(s)
// }

//指针
// func A(a *int) {
// 	*a = 88
// 	fmt.Println(*a)
// }

//
//实验4
//defer   go没有异常机制 用panic/recover模式来处理错误
////只执行 A  B
// func A() {
// 	fmt.Println("Func A")
// }

// func B() {

// 	panic("panic in B")
// }

// func C() {
// 	fmt.Println("Func c")
// }

//
//  先执行A 在执行C 在执行 fmt.Println("Recover in B")
// func A() {
// 	fmt.Println("Func A")
// }

// func B() {

// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println("Recover in B")
// 		}
// 	}()

// 	panic("panic in B")
// }

// func C() {
// 	fmt.Println("Func c")
// }

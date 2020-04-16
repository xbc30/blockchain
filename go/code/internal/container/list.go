package main

import (
	"container/list"
	"fmt"
	"strconv"
)
/*
list是一个双向链表。该结构具有链表的所有功能。
*/
func main01() {
	fmt.Println("入队列")
	//初始化一个list
	list1 := list.New()//初始化list
	//依次在尾部添加
	list1.PushBack(1)
	list1.PushBack("hello")
	list1.PushBack(3.14)
	list1.PushBack("江洲")
	//在头部添加
	list1.PushFront("888")


	//读取首元素
	frist := list1.Front().Value
	//fmt.Printf("%T",m)
	fmt.Println(frist)

	//获取最后一个入队列的元素
	last := list1.Back().Value
	fmt.Println(last)

}
func main() {
	list1 := list.New()
	for i := 0; i < 10; i++ {
		str :="hello"+strconv.Itoa(i)
		list1.PushBack(str)
	}
	list1.InsertAfter("你好", list1.Front())//插入数据：链表中首元素之后
	list1.InsertBefore("jiangzhou",list1.Back())//插入数据：链表中末尾元素之钱
	//fmt.Println(list1.Len())
	fmt.Println("遍历所用的元素")
	for i:=list1.Front();i!=nil;i=i.Next(){//依据链表进行
		fmt.Print(i.Value,"\t")
	}
	/*
	依次取出元素，同时删除原有数据
	*/
	fmt.Println()
	len := list1.Len() //获取原有长度
	for i := 0; i < len; i++ {
		de := list1.Front()
		fmt.Print(de.Value,"\t")
		list1.Remove(de) //出队列的首元素，同时在队列中将其删除
	}
}
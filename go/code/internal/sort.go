package main

import (
	"fmt"
	"sort"
)

func main()  {
	a:= sort.Float64Slice{1.1, 0.2, 3.3}
	b:= sort.IntSlice{1, 20, 3}
	c:= sort.StringSlice{"3.4", "0.5", "30.5"}
	d := sort.StringSlice{
		"啊",
		"博",
		"次",
		"得",
		"饿",
		"周",
	}

	sort.Float64s(a)
	sort.Ints(b)
	sort.Strings(c)
	sort.Strings(d) // 依次比较byte大小
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
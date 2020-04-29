package main

import "fmt"

func main() {

	//跟数组（arrays）不同，slices的类型跟所包含的元素类型一致（不是元素的数量）。使用内置的make命令，构建一个非零的长度的空slice对象。这里我们创建了一个包含了3个字符的字符串 。(初始化为零值zero-valued)
	s := make([]string, 3)
	fmt.Println("emp:", s)

	//我们可以像数组一样进行设置和读取操作。
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	//获取到的长度就是当时设置的长度。
	fmt.Println("len:", len(s))

	//相对于这些基本的操作，slices支持一些更加复杂的功能。有一个就是内置的append，可以在现有的slice对象上添加一个或多个值。注意要对返回的append对象重新赋值，以获取最新的添加了元素的slice对象。
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	//Slices也可以被复制。这里我们将s复制到了c，长度一致。
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	//Slices支持"slice"操作，语法为slice[low:high]（即截取slice中的某段值）。下面这段代码就会获取这些字符： s[2], s[3], 和 s[4]。
	l := s[2:5]
	fmt.Println("sl1:", l)

	//从开始截取到每5个字符（除了值）
	l = s[:5]
	fmt.Println("sl2:", l)

	//从第二个（包括）字符开始截取到最后一个
	l = s[2:]
	fmt.Println("sl3:", l)

	//我们可以将声明和赋值放在一行。
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	//Slices可以被组合成多维数组。里面一维的slices对象可以不等长，这一点跟多维数组不太一样。
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

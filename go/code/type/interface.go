package _type

import "fmt"

func Print(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("type is string,value is:%v\n", i.(string))
	case float64:
		fmt.Printf("type is float32,value is:%v\n", i.(float64))
	case int:
		fmt.Printf("type is int,value is:%v\n", i.(int))
	}
}
func main() {
	var i interface{}
	i = "hello"
	Print(i)
	i = 100
	Print(i)
	i = 1.29
	Print(i)
}
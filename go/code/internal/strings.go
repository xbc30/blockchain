package internal

import "fmt"
import "strings"

type User struct {
	Name string
	Age byte
}

func main() {

	a:= "21212p1kkj"
	b:=[]string{"asd","sdfg","dfg"}
	c:=append(b[:1], b[2:]...)
	d:= strings.Split(a, "p")
	f:=[...]User{
		{"asd", 10},
	}

	fmt.Println(d)
	fmt.Println(len(d))
	fmt.Println(strings.HasPrefix(a, "2"))
	fmt.Println(strings.HasSuffix(a, "j"))
	fmt.Println(strings.Index(a, "p"))
	fmt.Println(strings.Replace(a,"21", "qw", 2))
	fmt.Println(strings.ToUpper(a))
	fmt.Println(strings.Join(b, "-"))
	fmt.Println(strings.TrimRight(a, "kkj"))
	fmt.Println(c)
	fmt.Println(f)

	// fmt.Println(b.Contains("asd"))
	fmt.Println("Hello World")
}

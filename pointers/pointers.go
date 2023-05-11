package main

import "fmt"

func main() {
	i, j := 42, 1321

	p := &i
	fmt.Println(*p)
	fmt.Println("%T\n", p)

	*p = 21
	fmt.Println(i)
	p = &j
	*p = *p / 37
	fmt.Println(j)
}

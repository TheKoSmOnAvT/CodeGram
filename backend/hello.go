package backend

import "fmt"

type lib []string

func (l lib) print() {
	for key, val := range l {
		fmt.Println(key, " ", val)
	}
}

func change(x *int) {
	*x = *x * 10
}

func main() {

	var mass lib = lib{"f", "a"}
	mass.print()

	var first int = 4
	//var test *int
	//test = &first
	change(&first)
	fmt.Println(first)
}

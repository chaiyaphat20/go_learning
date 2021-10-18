package functions

import (
	"fmt"
	"strings"
)

var Fullname string

func add(x, y int) int {
	return x + y
}

func sum() int {
	return 5
}

func converter(title, email string) (string, string) {
	title = strings.ToUpper(title)
	return title, email
}

func test(numbers ...int) { // ...int  = Variadic function
	fmt.Println(numbers) //array
	fmt.Printf("type of number is %T \n", numbers)
	total := 0
	for _, v := range numbers {
		total += v
	}
	fmt.Println(total)
}

func Learn() {
	fmt.Println(add(sum(), 3))
	fmt.Println(converter("Hi", "Art"))
	test(10, 20, 30)
}

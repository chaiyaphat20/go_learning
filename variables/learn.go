package variables

import (
	"fmt"
	"strconv"
)

var fullname string //global
var email string = "codingthailand@gmail.com"
var c, python bool = true, false

const pi = 3

func Learn() {
	fmt.Println("-----------------func Learn()-----------------")
	fmt.Printf("Hello %s EMail %s \n", fullname, email)
	fmt.Println(c, python)

	companyName := "ArtThailand"
	isShow := true
	fmt.Println(companyName)
	fmt.Println(isShow)
	fmt.Printf("type of isShow is %T \n", isShow)

	age := 10
	f := float64(age)
	fmt.Printf("Type of f is %T \n", f)

	s := strconv.Itoa(pi) //convert number to string
	fmt.Println(s)

	fmt.Println("---------------------END----------------------")
}

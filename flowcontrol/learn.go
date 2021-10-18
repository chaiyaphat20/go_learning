package flowcontrol

import (
	"fmt"
	"io/ioutil"
	"runtime"
)

func Learn() {
	//if
	score := 10
	if score == 10 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	//switch case
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("MacOS")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Println("os")
	}

	b, err := ioutil.ReadFile("file.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		constent := string(b)
		fmt.Println(constent)
	}
}

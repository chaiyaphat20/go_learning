package flowcontrol

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"time"
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

	//for
	for j := 5; j >= 0; j-- {
		if j == 0 {
			fmt.Println("Boom!")
			break
		}
		fmt.Println(j)
		time.Sleep(1 * time.Second)
	}
}

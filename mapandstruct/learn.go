package mapandstruct

import "fmt"

func Learn() {
	//Maps (key, value)
	//var m map[struct]
	fmt.Println("------------------------Maps--------------------------------------")
	m := map[string]int{"token": 10, "access": 20}
	fmt.Println(m["token"])

	t := make(map[string]int)
	t["test"] = 200
	fmt.Println(t["test"])

	//for loop
	for key, v := range m {
		fmt.Printf("%v %v \n", key, v)
	}

	//delete map
	delete(m, "access")

	//add map
	m["show"] = 10000

	fmt.Println(m)

	//struct
	fmt.Println("------------------------struct--------------------------------------")
	type User struct { //ตัวเล็ก = private , ตัวใหญ่ = Public
		id       int
		username string
		password string
	}

	john := User{
		id:       1,
		username: "Chaiyaphat",
		password: "123456",
	}

	fmt.Println(john.id)
	fmt.Println(john.username)
	john.password = "cccccc"
	fmt.Println(john.password)

	//slice of struct
	users := []User{
		{id: 1, username: "xc", password: "1231231313"},
		{id: 2, username: "vv", password: "vvvvv"},
	}
	fmt.Println(users[0].password)

}

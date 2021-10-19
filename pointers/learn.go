package pointers

import "fmt"

func Learn() {
	x := 10
	fmt.Println(x)
	fmt.Println(&x) // & คือ get memberry address   //0xc0000180a8

	y := x
	fmt.Println(y)
	fmt.Println(&y)

	println("------pointer-----")
	p := &x //ชี้ไป x  same var p *int = &x
	fmt.Println(p)
	fmt.Println(*p) //เป็นการอ่านค่า x ผ่าน p pointer   (dereference)

	//change value
	*p = 20
	fmt.Println(*p)
	fmt.Println(p)

	b := student{name: "Bob"}
	fmt.Println(b)
	setName(&b)
	fmt.Println(b)

}

type student struct {
	name string
}

func setName(std *student) {
	std.name = "Boy"
}

/*
	& ถ้าวางหน้าตัวแปร เป็นการ อ่าน memory address
	* ถ้าวางไว้หน้าตัวแปร เป็นการดึงค่า ที่ memory address นั้นๆ  (dereference)
	* ถ้าวางไว้ หน้า type (ชนิดข้อมูล) ex.  *int   *student คือการเก็บค่า memory address ไว้  (ว่างๆคือ nil)

*/

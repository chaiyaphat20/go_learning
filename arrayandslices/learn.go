package arrayandslices

import "fmt"

func Learn() {
	//array
	var arr [2]string
	arr[0] = "MaJane"
	arr[1] = "ART"
	fmt.Println(arr)

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	//slice คือ ไม่ต้องระบุขนาด
	var s []int = arr2[1:2] //จำว่าอันแรก เอา จน ถึง อันสุดท้ายไม่เอา
	s2 := []int{1, 2, 3, 4}
	fmt.Println(s)
	fmt.Println(s2)
	fmt.Printf("len = %d cap=%d  of  %v \n", len(s2), cap(s2), s2)

	//concat
	x := []int{1, 2, 3, 4}
	y := []int{9, 10, 11}
	z := append(x, y...)
	fmt.Println(z)

}

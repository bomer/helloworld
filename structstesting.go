package main

import "fmt"

type Rectangle struct {
	length, width int
	name          string
}

func (r Rectangle) display() {
	fmt.Printf("Length = %d, Width = %d, Name = %s", r.length, r.width, r.name)
	fmt.Println()
	return
}

func main() {
	r1 := Rectangle{4, 5, "rec1"}
	r1.display()
	// fmt.Printf("Rectange 1", r1)
	// fmt.Println()

	r2 := Rectangle{length: 10, width: 20, name: "IM A BOX"}
	// fmt.Printf("Rectange 2 = %+v", r2)
	r2.display()

	r3 := new(Rectangle)
	r3.length = 2
	r3.width = 4
	r3.name = "hi there"
	// fmt.Println(*r3)
	r3.display()
}

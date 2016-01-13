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
	var r1 = Rectangle{4, 5, "rec1"}
	r1.display()

	r2 := Rectangle{length: 10, width: 20, name: "IM A BOX"}
	r2.display()

}

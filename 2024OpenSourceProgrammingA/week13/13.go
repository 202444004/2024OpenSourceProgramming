package main

import "fmt"

func main() {
	var student1 struct {
		id   int
		name string
		gpa  float32
	}
	student1.id = 20241234
	student1.name = "tae"
	student1.gpa = 2.5
	fmt.Println(student1.gpa)

	var student2 struct {
		id   int
		name string
		gpa  float32
	}
	student2.id = 20241235
	student2.name = "gan"
	student2.gpa = 1.5
	fmt.Println(student2.gpa)
}

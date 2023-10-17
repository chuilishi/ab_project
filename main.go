package main

import "fmt"

type User struct {
	name string
	Age  string
<<<<<<< HEAD
	id   int
}

func (u User) AddName(a int) {
	u.id = a
}
func main() {
	var user1 User
	user1.name = "John"
	println(user1.name)
=======
	Id   int
}

func main() {
	//u:=User{}
	fmt.Println("Hello world")
	fmt.Println("nihao")
	println("hah")

>>>>>>> main
}

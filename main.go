package main

type User struct {
	name string
	Age  string
	id   int
}

func (u User) AddName(a int) {
	u.id = a
}
func main() {
	var user1 User
	user1.name = "John"
	println(user1.name)
}

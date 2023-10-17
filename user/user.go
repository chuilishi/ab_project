package main

import (
	"fmt"
	"time"
)

type User struct {
	Name string
	Age  int
	Sex  string
}

func (u *User) setName(name string) {
	u.Name = name
}
func (u User) prin() {
	fmt.Print(u.Age, u.Sex, u.Name)
}
func main() {
	u := User{}

	u.Age = 1
	u.Sex = "nan1"
	u.setName("nihao")
	fmt.Println(1)
	u.prin()
	time.Sleep(time.Second * 5)

}

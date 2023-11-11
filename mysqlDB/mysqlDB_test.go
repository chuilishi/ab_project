package mysqlDB

import (
	"fmt"
	"testing"
)

func TestSendMessageToUser(t *testing.T) {
	InitGrom()

}
func TestFindUserMessage(t *testing.T) {
	InitGrom()
	messages, _ := FindUserMessage("3")
	for _, message := range messages {
		fmt.Println(message)
	}
}
func TestChangeUserStatus(t *testing.T) {
	InitGrom()
	ChangeUserStatus(1, "成功录取")

}
func TestPostProblem(t *testing.T) {
	InitGrom()
	PostProblem("2621iku823", "我想笑")
}
func TestAllUserStatus(t *testing.T) {
	InitGrom()
	user, _ := AllUserStatus()
	fmt.Println(user)
}

package utils

import (
	"fmt"
	"testing"
)

func TestSendMail(t *testing.T) {
	err := SendMail("l190543@outlook.com", "daijinyu")
	fmt.Println(err)
}

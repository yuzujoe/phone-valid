package sms

import (
	"fmt"
	"strings"
	"testing"
)

func TestPushSms(t *testing.T) {
	phoneNumber := "08071985692"
	fmt.Println(phoneNumber)

	phone := strings.Replace(phoneNumber, "0", "+81", 1)
	fmt.Println(phone)
}

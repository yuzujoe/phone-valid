package sms

import (
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

// PushSms SMS送信用のロジック
func PushSms(phoneNumber, code string) error {
	fmt.Println("create session")
	sess, err := session.NewSession()

	if err != nil {
		log.Fatalln(err)
	}

	svc := sns.New(sess)

	phone := strings.Replace(phoneNumber, "0", "+81", 1)

	input := &sns.PublishInput{
		Subject:     aws.String("test message"),
		Message:     aws.String("your auth code is" + code + "thats expire at 15 min"),
		PhoneNumber: aws.String(phone),
	}

	fmt.Println(input)

	_, err = svc.Publish(input)
	if err != nil {
		return err
	}

	return nil
}

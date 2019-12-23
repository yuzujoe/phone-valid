package sms

import (
	"fmt"
	"log
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
		Message:     aws.String("あなたの認証コードは" + code + "です"),
		PhoneNumber: aws.String(phone),
	}

	_, err = svc.Publish(input)
	if err != nil {
		return err
	}


	return nil
}

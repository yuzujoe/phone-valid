package sms

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

// PushSms SMS送信用のロジック
func PushSms(phoneNumber, code string) error {
	fmt.Println("create session")

	AccessKey := os.Getenv("AWS_ACCESS_KEY")
	SecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Credentials: credentials.NewStaticCredentials(AccessKey, SecretAccessKey, ""),
	})

	if err != nil {
		log.Fatalln(err)
	}

	svc := sns.New(sess)

	input := &sns.PublishInput{
		Message:     aws.String("messege" + code + "test"),
		PhoneNumber: aws.String(phoneNumber),
	}

	result, err := svc.Publish(input)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(result)

	return nil
}

package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"log"
)

func PushSms(phoneNumber string, code string) {
	fmt.Println("create session")

	sess := session.Must(session.NewSession())

	svc := sns.New(sess)
	fmt.Println(svc)

	params := &sns.PublishInput{
		Message:     aws.String("こちらの確認コードを入力してください" + code),
		PhoneNumber: aws.String(phoneNumber),
	}

	res, err := svc.Publish(params)
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println(res)
}

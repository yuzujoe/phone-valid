package sms

import (
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/sms"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/gin-gonic/gin"
)

var ErrGetAwsSession = errors.New("Failed get aws session")

// PushSms SMS送信用のロジック
func PushSms(c *gin.Context, phoneNumber, code string) error {
	fmt.Println("create session")
	sess, err := getAwsSession(c)
	if err != nil {
		return fmt.Errorf("get aws session error: %w", err)
	}

	svc := sns.New(sess, aws.NewConfig().WithRegion("ap-northeast-1"))

	phone := strings.Replace(phoneNumber, "0", "+81", 1)

	sms := &sms.SMS{}

	input := &sns.PublishInput{
		Message:     aws.String("your auth code is" + code + "thats expire at 15 min"),
		PhoneNumber: aws.String(phone),
	}

	fmt.Println(input)

	_, err = svc.Publish(input)
	if err != nil {
		return fmt.Errorf("auth code send sms error: %w", err)
	}

	fmt.Println("ok")

	return nil
}

func getAwsSession(c *gin.Context) (client.ConfigProvider, error) {
	sessPtr, ok := c.Get("AwsSession")
	if !ok {
		return nil, ErrGetAwsSession
	}

	sess := sessPtr.(client.ConfigProvider)
	if sess == nil {
		return nil, ErrGetAwsSession
	}

	fmt.Println(sess)

	return sess, nil
}

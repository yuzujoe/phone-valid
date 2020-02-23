package sms

import (
	"errors"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
	"github.com/gin-gonic/gin"
)

var ErrGetAwsSession = errors.New("Failed get aws session")

type SMS struct {
	Service snsiface.SNSAPI
}

// SMS送信用のロジック
func (s *SMS) sms(phoneNumber, code string) error {

	phone := strings.Replace(phoneNumber, "0", "+81", 1)

	input := &sns.PublishInput{
		Message:     aws.String("your auth code is" + code + "thats expire at 15 min"),
		PhoneNumber: aws.String(phone),
	}

	_, err := s.Service.Publish(input)
	if err != nil {
		return fmt.Errorf("auth code send sms error: %w", err)
	}

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

	return sess, nil
}

func SendSms(c *gin.Context, phoneNumber, code string) error {
	fmt.Println("create session")

	sess, err := getAwsSession(c)
	if err != nil {
		return fmt.Errorf("get aws session error: %w", err)
	}

	service := sns.New(sess, aws.NewConfig().WithRegion("ap-northeast-1"))
	sms := SMS{Service: service}
	return sms.sms(phoneNumber, code)
}

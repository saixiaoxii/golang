package data

import (
	"context"
	"customer/api/verifyCode"
	"regexp"

	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func GetVerifyCode(telephone string) (string, error) {
	pattern := `^1[3-9]\d{9}$`

	if !regexp.MustCompile(pattern).MatchString(telephone) {
		return "手机号格式错误", nil
	}

	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("localhost:9000"),
	)
	if err != nil {
		return "验证码服务连接失败", nil
	}
	defer conn.Close()

	client := verifyCode.NewVerifyCodeClient(conn)
	response, err := client.GetVerifyCode(context.Background(), &verifyCode.GetVerifyCodeRequest{
		Length: 6,
		Type:   1,
	})
	if err != nil {
		return "验证码服务调用失败", nil
	}

	return response.Code, nil
}

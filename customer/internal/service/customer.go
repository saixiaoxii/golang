package service

import (
	"context"
	"regexp"
	"time"

	"github.com/redis/go-redis/v9"

	pb "customer/api/customer"
	"customer/api/verifyCode"

	"github.com/go-kratos/kratos/v2/transport/grpc"
)

type CustomerService struct {
	pb.UnimplementedCustomerServer
}

func NewCustomerService() *CustomerService {
	return &CustomerService{}
}

func (s *CustomerService) GetVerifyCode(ctx context.Context, req *pb.GetVerifyCodeReq) (*pb.GetVerifyCodeResp, error) {
	// 验证手机号
	pattern := `^1[3-9]\d{9}$`

	if !regexp.MustCompile(pattern).MatchString(req.Telephone) {
		return &pb.GetVerifyCodeResp{
			Code:    1,
			Message: "手机号格式错误",
		}, nil
	}
	// 生成验证码
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("localhost:9000"),
	)
	if err != nil {
		return &pb.GetVerifyCodeResp{
			Code:    1,
			Message: "验证码服务连接失败",
		}, nil
	}
	defer conn.Close()
	// 调用验证码服务
	client := verifyCode.NewVerifyCodeClient(conn)
	response, err := client.GetVerifyCode(context.Background(), &verifyCode.GetVerifyCodeRequest{
		Length: 6,
		Type:   1,
	})
	if err != nil {
		return &pb.GetVerifyCodeResp{
			Code:    1,
			Message: "验证码服务调用失败",
		}, nil
	}
	//redis存储验证码
	options, err := redis.ParseURL("redis://192.168.29.130:6379/1?dial_timeout=5")
	if err != nil {
		return &pb.GetVerifyCodeResp{
			Code:    1,
			Message: "验证码存储失败",
		}, nil
	}
	rdb := redis.NewClient(options)
	status := rdb.Set(context.Background(), "customer-verify-code:"+req.Telephone, response.Code, 60*time.Second)
	if status.Err() != nil {
		return &pb.GetVerifyCodeResp{
			Code:    1,
			Message: "验证码存储失败Set",
		}, nil
	}
	return &pb.GetVerifyCodeResp{
		Code:           0,
		Message:        "验证码发送成功",
		VerifyCode:     response.Code,
		VerifyCodeTime: time.Now().Unix(),
		VerifyCodeLife: 60,
	}, nil
}

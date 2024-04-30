package service

import (
	"context"
	"math/rand"

	pb "verify-code/api/verifyCode"
)

type VerifyCodeService struct {
	pb.UnimplementedVerifyCodeServer
}

// 注册到集合
func NewVerifyCodeService() *VerifyCodeService {
	return &VerifyCodeService{}
}

func (s *VerifyCodeService) CreateVerifyCode(ctx context.Context, req *pb.CreateVerifyCodeRequest) (*pb.CreateVerifyCodeReply, error) {
	return &pb.CreateVerifyCodeReply{}, nil
}
func (s *VerifyCodeService) UpdateVerifyCode(ctx context.Context, req *pb.UpdateVerifyCodeRequest) (*pb.UpdateVerifyCodeReply, error) {
	return &pb.UpdateVerifyCodeReply{}, nil
}
func (s *VerifyCodeService) DeleteVerifyCode(ctx context.Context, req *pb.DeleteVerifyCodeRequest) (*pb.DeleteVerifyCodeReply, error) {
	return &pb.DeleteVerifyCodeReply{}, nil
}
func (s *VerifyCodeService) GetVerifyCode(ctx context.Context, req *pb.GetVerifyCodeRequest) (*pb.GetVerifyCodeReply, error) {
	return &pb.GetVerifyCodeReply{
		Code: RandCode(int(req.Length), req.Type),
	}, nil
}
func (s *VerifyCodeService) ListVerifyCode(ctx context.Context, req *pb.ListVerifyCodeRequest) (*pb.ListVerifyCodeReply, error) {
	return &pb.ListVerifyCodeReply{}, nil
}

// 开放的调用方法，区分类型
func RandCode(l int, t pb.TYPE) string {
	switch t {
	case pb.TYPE_DEFAULT:
		fallthrough
	case pb.TYPE_DIGIT:
		chars := "0123456789"
		return randCode(l, chars)
	case pb.TYPE_LETTER:
		chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
		return randCode(l, chars)
	case pb.TYPE_MIXED:
		chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
		return randCode(l, chars)
	default:

	}

	return ""
}

// 随机的核心方法
func randCode(l int, chars string) string {
	result := make([]byte, l)
	for i := 0; i < l; i++ {
		randIndex := rand.Intn(len(chars))
		result[i] = chars[randIndex]
	}
	return string(result)
}

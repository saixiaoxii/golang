package data

import (
	"context"
	"time"
)

type CustomerData struct {
	data *Data
}

func NewCustomerData(data *Data) *CustomerData {
	return &CustomerData{data: data}
}

// 设置验证码的方法
func (data CustomerData) SetVerifyCode(telephone, code string, ex int64) error {
	status := data.data.Rdb.Set(context.Background(), "CVC:"+telephone, code, time.Duration(ex)*time.Second)
	if _, err := status.Result(); err != nil {
		return err
	}

	return nil
}

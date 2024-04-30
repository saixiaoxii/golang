package data

type CustomerData struct {
	data *Data
}

func NewCustomerData(data *Data) *CustomerData {
	return &CustomerData{data: data}
}

// 设置验证码的方法
func (data CustomerData) SetVerifyCode(code string, ex int) {

}

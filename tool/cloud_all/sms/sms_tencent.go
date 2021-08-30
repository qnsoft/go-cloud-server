package sms

import (
	"encoding/json"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"strings"
)

type SmsTenCent struct {
	SecretID  string
	SecretKey string
}

/**
 * @Description: 发送短信
 * @receiver this
 * @param _sign 短信签名
 * @param phone_numbers 手机号,多个手机号之间用半角逗号隔开,
 * @param template_code 模板编号
 * @param template_param 模板参数
 * @return interface{}
 * @return error
 */
func (this *SmsTenCent) SendSMS(_sign, phone_numbers, template_code, template_param string) (interface{}, error) {
	var _data interface{}
	var _err error
	credential := common.NewCredential(
		this.SecretID,
		this.SecretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	client, _ := sms.NewClient(credential, "ap-beijing", cpf)
	request := sms.NewSendSmsRequest()
	var _numbersA []string
	for _, item := range strings.Split(phone_numbers, ",") {
		_numbersA = append(_numbersA, fmt.Sprintf("+86%s", item))
	}
	request.PhoneNumberSet = common.StringPtrs(_numbersA)
	request.SmsSdkAppId = common.StringPtr("1400563159")
	request.SignName = common.StringPtr(_sign)
	request.TemplateId = common.StringPtr(template_code)
	_param_model := struct {
		Code string `json:"code"`
		Time string `json:"time"`
	}{}
	json.Unmarshal([]byte(template_param), &_param_model)
	nr := []*string{}
	nr = append(nr, &_param_model.Code)
	nr = append(nr, &_param_model.Time)
	request.TemplateParamSet = nr
	_result, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
		_err = err
	}
	_data = map[string]interface{}{
		"result": _result.ToJsonString(),
	}
	return _data, _err
}

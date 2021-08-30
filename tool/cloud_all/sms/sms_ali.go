package sms

import (
	"fmt"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

type SmsAli struct {
	SecretID  string
	SecretKey string
}

/**
 * @Description: 发送短信
 * @receiver this
 * @param _sign 短信签名
 * @param phone_numbers 手机号
 * @param template_code 模板编号
 * @param template_param 模板参数
 * @return interface{}
 * @return error
 */
func (this *SmsAli) SendSMS(_sign, phone_numbers, template_code, template_param string) (interface{}, error) {
	var _data interface{}
	var _err error
	config := &openapi.Config{
		AccessKeyId:     tea.String(this.SecretID),
		AccessKeySecret: tea.String(this.SecretKey),
	}
	client, _ := dysmsapi20170525.NewClient(config)
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  &phone_numbers,  //接收短信的手机号码。支持对多个手机号码发送短信，手机号码之间以英文逗号（,）分隔。上限为1000个手机号码。批量调用相对于单条调用及时性稍有延迟。
		SignName:      &_sign,          //短信签名名称。请在控制台签名管理页面签名名称一列查看。
		TemplateCode:  &template_code,  //短信模板ID。请在控制台模板管理页面模板CODE一列查看。例如：SMS_153055065
		TemplateParam: &template_param, //短信模板变量对应的实际值，JSON格式。例如：{"code":"1111"}
	}
	_result, err := client.SendSms(sendSmsRequest)
	if err != nil {
		fmt.Print(err.Error())
		_err = err
	}
	_data = map[string]interface{}{
		"result": _result.Body,
	}
	return _data, _err
}

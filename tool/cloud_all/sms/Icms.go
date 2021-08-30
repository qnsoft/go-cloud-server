package sms

import ()

type ISMS interface {
	//发送短信
	SendSMS(_sign, phone_numbers, template_code, template_param string) (interface{}, error)
}

/*
混合云对象工厂
*/
func ObjectSMS(_clout int, SecretID, SecretKey string) ISMS {
	var smsObject ISMS
	switch _clout {
	case 0: //阿里云
		smsObject = &SmsAli{SecretID: SecretID, SecretKey: SecretKey}
	case 1: //腾讯云
		smsObject = &SmsTenCent{SecretID: SecretID, SecretKey: SecretKey}
	}
	return smsObject
}

package cloud_all

//云类型
const (
	//阿里云
	Ali = iota
	//腾讯云
	TenCent
	//华为云
	Huawei
	//百度云
	Baidu
)

//云授权
var (
	SecretID  string = ""
	SecretKey string = ""
)

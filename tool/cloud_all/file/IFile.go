package file
"github.com/qnsoft/cloud_server/tool/cloud_all"
/*
对象云存储接口
*/
type IFile interface {
	//查询存储桶列表
	GetBuckets() ([]interface{}, error)
	//上传对象
	UploadObject(local_file_path, server_bucket_name, server_directory, server_file string) error
	//获取上传对象列表
	GetObjectList(server_bucket_name string) ([]interface{}, error)
	//下载对象
	ObjectDownLoad(server_bucket_name, server_directory, server_file, local_file_path string) error
	//删除对象
	ObjectDelete(server_bucket_name, server_directory, server_file string) error
}

/*
混合云对象工厂
*/
func ObjectFile(_clout int) IFile {
	var fileObject IFile
	switch _clout {
	case 0://阿里云
		fileObject=&CloudAli{SecretID: cloud_all.SecretID, SecretKey: cloud_all.SecretKey,CloudUrl: "https://oss-cn-beijing.aliyuncs.com"}
	case 1://腾讯云
		fileObject=&CloudTencent{SecretID: cloud_all.SecretID, SecretKey: cloud_all.SecretKey,CloudUrl: "https://img1.qnsoft.net"}
	}
	return fileObject
}

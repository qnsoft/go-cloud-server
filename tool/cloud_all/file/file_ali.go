package file

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"net/http"
	"os"
)

/*
腾讯云操作对象
*/
type FileAli struct {
	SecretID  string
	SecretKey string
	//云存储域名
	CloudUrl string
}

/*
查询存储桶列表
*/
func (this *FileAli) GetBuckets() ([]interface{}, error) {
	var _list []interface{}
	var _err error
	client, err := oss.New(this.CloudUrl, this.SecretID, this.SecretKey)
	if err != nil {
		//fmt.Println("Error:", err)
		//os.Exit(-1)
		_err = err
	}
	// 列举存储空间。
	// 限定此次列举存储空间的个数为500。默认值为100，最大值为1000。
	lsRes, err := client.ListBuckets(oss.MaxKeys(500))
	if err != nil {
		fmt.Println("Error:", err)
		//os.Exit(-1)
		_err = err
	}
	for _i, bucket := range lsRes.Buckets {
		_list = append(_list, map[int]interface{}{_i: bucket.Name})
	}
	return _list, _err
}

/*
上传本地对象
*/
func (this *FileAli) UploadObject(local_file_path, server_bucket_name, server_directory, server_file string) error {
	var _err error
	// 创建OSSClient实例。
	client, err := oss.New(this.CloudUrl, this.SecretID, this.SecretKey)
	if err != nil {
		//fmt.Println("Error:", err)
		//os.Exit(-1)
		_err = err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(server_bucket_name)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 上传本地文件。
	err = bucket.PutObjectFromFile(server_directory+"/"+server_file, local_file_path)
	if err != nil {
		fmt.Println("Error:", err)
		//os.Exit(-1)
		_err = err
	}

	//
	//
	//u, _ := url.Parse(this.CloudUrl)
	//b := &cos.BaseURL{BucketURL: u}
	//c := cos.NewClient(b, &http.Client{
	//	Transport: &cos.AuthorizationTransport{
	//		SecretID:  this.SecretID,
	//		SecretKey: this.SecretKey,
	//	},
	//})
	//name := server_directory+"/"+server_file
	//// 1.通过字符串上传对象
	//f := strings.NewReader(server_directory)
	//
	//_, err := c.Object.Put(context.Background(), name, f, nil)
	//if err != nil {
	//	_err=err
	//}
	//// 2.通过本地文件上传对象
	//_, err = c.Object.PutFromFile(context.Background(), name, local_file_path, nil)
	//if err != nil {
	//	_err=err
	//}
	//// 3.通过文件流上传对象
	//fd, err := os.Open(local_file_path)
	//if err != nil {
	//	_err=err
	//}
	//defer fd.Close()
	//_, err = c.Object.Put(context.Background(), name, fd, nil)
	//if err != nil {
	//	_err=err
	//}
	return _err
}

/*
获取对象列表
*/
func (this *FileAli) GetObjectList(server_bucket_name string) ([]interface{}, error) {
	var _list []interface{}
	var _err error
	// 创建OSSClient实例。
	client, err := oss.New(this.CloudUrl, this.SecretID, this.SecretKey)
	if err != nil {
		//fmt.Println("Error:", err)
		//os.Exit(-1)
		_err = err
	}

	// 获取存储空间。
	bucketName := server_bucket_name
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		_err = err
	}

	continueToken := ""
	for {
		lsRes, err := bucket.ListObjectsV2(oss.ContinuationToken(continueToken))
		if err != nil {
			_err = err
		}

		// 打印列举结果。默认情况下，一次返回100条记录。
		for _, object := range lsRes.Objects {
			//fmt.Println(object.Key, object.Type, object.Size, object.ETag, object.LastModified, object.StorageClass)
			_list = append(_list, map[string]interface{}{"file": object.Key, "size": object.Size})
		}

		if lsRes.IsTruncated {
			continueToken = lsRes.NextContinuationToken
		} else {
			break
		}
	}
	return _list, _err
}

/*
对象下载
*/
func (this *FileAli) ObjectDownLoad(server_bucket_name, server_directory, server_file, local_file_path string) error {

	var _err error
	// 创建OSSClient实例。
	client, err := oss.New(this.CloudUrl, this.SecretID, this.SecretKey)
	if err != nil {
		//fmt.Println("Error:", err)
		//os.Exit(-1)
		_err = err
	}
	// 获取存储空间。
	bucket, err := client.Bucket(server_bucket_name)
	if err != nil {
		fmt.Println("Error:", err)
		_err = err
	}

	var retHeader http.Header
	// 下载yourObjectVersionId版本的文件到缓存。
	_, err = bucket.GetObject(server_directory, oss.VersionId(server_directory+"/"+server_file), oss.GetResponseHeader(&retHeader))
	if err != nil {
		fmt.Println("Error:", err)
		_err = err
	}

	return _err
}

/*
对象删除
*/
func (this *FileAli) ObjectDelete(server_bucket_name, server_directory, server_file string) error {
	var _err error
	// 创建OSSClient实例。
	client, err := oss.New(this.CloudUrl, this.SecretID, this.SecretKey)
	if err != nil {
		//fmt.Println("Error:", err)
		//os.Exit(-1)
		_err = err
	}
	bucketName := server_bucket_name
	objectName := server_directory + "/" + server_file

	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		_err = err
	}

	// 删除单个文件。objectName表示删除OSS文件时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	// 如需删除文件夹，请将objectName设置为对应的文件夹名称。如果文件夹非空，则需要将文件夹下的所有object删除后才能删除该文件夹。
	err = bucket.DeleteObject(objectName)
	if err != nil {
		fmt.Println("Error:", err)
		_err = err
	}

	return _err
}

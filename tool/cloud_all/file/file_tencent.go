package file

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

/*
腾讯云操作对象
*/
type FileTencent struct {
	SecretID  string
	SecretKey string
	//云存储域名
	CloudUrl string
}

/*
查询存储桶列表
*/
func (this *FileTencent) GetBuckets() ([]interface{}, error) {
	var _list []interface{}
	var _err error
	c := cos.NewClient(nil, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  this.SecretID,
			SecretKey: this.SecretKey,
		},
	})
	//查询存储桶列表
	s, _, err := c.Service.Get(context.Background())
	if err != nil {
		_err = err
	}

	for _i, b := range s.Buckets {
		//fmt.Printf("%#v\n", b)
		_list = append(_list, map[int]interface{}{_i: b})
	}
	return _list, _err
}

/*
上传本地对象
*/
func (this *FileTencent) UploadObject(local_file_path, server_bucket_name, server_directory, server_file string) error {
	var _err error
	u, _ := url.Parse(this.CloudUrl)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  this.SecretID,
			SecretKey: this.SecretKey,
		},
	})
	// 对象键（Key）是对象在存储桶中的唯一标识。
	// 例如，在对象的访问域名 `examplebucket-1250000000.cos.COS_REGION.myqcloud.com/test/objectPut.go` 中，对象键为 test/objectPut.go
	name := server_directory + "/" + server_file
	// 1.通过字符串上传对象
	f := strings.NewReader(server_directory)

	_, err := c.Object.Put(context.Background(), name, f, nil)
	if err != nil {
		_err = err
	}
	// 2.通过本地文件上传对象
	_, err = c.Object.PutFromFile(context.Background(), name, local_file_path, nil)
	if err != nil {
		_err = err
	}
	// 3.通过文件流上传对象
	fd, err := os.Open(local_file_path)
	if err != nil {
		_err = err
	}
	defer fd.Close()
	_, err = c.Object.Put(context.Background(), name, fd, nil)
	if err != nil {
		_err = err
	}
	return _err
}

/*
获取对象列表
*/
func (this *FileTencent) GetObjectList(server_bucket_name string) ([]interface{}, error) {
	var _list []interface{}
	var _err error
	u, _ := url.Parse(this.CloudUrl)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  this.SecretID,
			SecretKey: this.SecretKey,
		},
	})

	opt := &cos.BucketGetOptions{
		Prefix:  "",
		MaxKeys: 3,
	}
	v, _, err := c.Bucket.Get(context.Background(), opt)
	if err != nil {
		_err = err
	}
	for _, c := range v.Contents {
		//fmt.Printf("%s, %d\n", c.Key, c.Size)
		_list = append(_list, map[string]interface{}{"file": c.Key, "size": c.Size})
	}

	return _list, _err
}

/*
对象下载
*/
func (this *FileTencent) ObjectDownLoad(server_bucket_name, server_directory, server_file, local_file_path string) error {
	var _err error
	u, _ := url.Parse(this.CloudUrl)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  this.SecretID,
			SecretKey: this.SecretKey,
		},
	})
	// 1.通过响应体获取对象
	name := server_directory + "/" + server_file
	resp, err := c.Object.Get(context.Background(), name, nil)
	if err != nil {
		_err = err
	}
	bs, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("%s\n", string(bs))
	// 2.获取对象到本地文件
	_, err = c.Object.GetToFile(context.Background(), name, local_file_path, nil)
	if err != nil {
		_err = err
	}
	return _err
}

/*
对象删除
*/
func (this *FileTencent) ObjectDelete(server_bucket_name, server_directory, server_file string) error {
	var _err error
	u, _ := url.Parse(this.CloudUrl)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  this.SecretID,
			SecretKey: this.SecretKey,
		},
	})
	name := server_directory + "/" + server_file
	if server_file == "" {
		name = server_directory
	}
	_, err := c.Object.Delete(context.Background(), name)
	if err != nil {
		_err = err
	}
	return _err
}

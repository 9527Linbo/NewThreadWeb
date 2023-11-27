package utils

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
)

var client *oss.Client
var bucketName string

func InitOSS() error {

	bucketName = viper.GetString("OSS.BucketName")

	// 从环境变量中获取访问凭证。运行本代码示例之前，请确保已设置环境变量OSS_ACCESS_KEY_ID和OSS_ACCESS_KEY_SECRET。
	provider, err := oss.NewEnvironmentVariableCredentialsProvider()
	if err != nil {
		return err
	}

	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	client, err = oss.New("oss-cn-wulanchabu.aliyuncs.com", "", "", oss.SetCredentialsProvider(&provider))
	if err != nil {
		return err
	}
	return nil
}

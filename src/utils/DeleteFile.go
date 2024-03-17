package utils

import (
	"github.com/spf13/viper"
)

func DeleteIcon(fileuuid string) error {
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}

	path := viper.GetString("File.OSSIconPath")
	objectName := path + fileuuid
	// 删除头像。
	err = bucket.DeleteObject(objectName)
	if err != nil {
		return err
	}
	return nil
}

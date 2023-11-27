package utils

func Upload_Simple_File() error {
	// 填写存储空间名称，例如examplebucket。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}

	// 依次填写Object的完整路径（例如exampledir/exampleobject.txt）和本地文件的完整路径（例如D:\\localpath\\examplefile.txt）。
	err = bucket.PutObjectFromFile("TestDir/test.txt", "C:/Users/9527/Desktop/新建文本文档.txt")
	if err != nil {
		return err
	}
	return nil
}

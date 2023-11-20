package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path/filepath"
)

func InitMinioClient() *minio.Client {
	// 基本的配置信息
	endpoint := "127.0.0.1:9000"
	/*	accessKeyID := "minioadmin"
		secretAccessKey := "minioadmin"*/
	//这是由用户设置的两个密码
	accessKeyID := "admin"
	secretAccessKey := "admin123456"
	// 初始化一个minio客户端对象
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		log.Fatalf("初始化MinioClient错误：%s", err.Error())
	}
	return minioClient
}
func UploadFile(c *gin.Context) {
	// 创建客户端
	minioClient := InitMinioClient()
	// bucket名称
	bucketName := "mypic"
	//这里的bucket可以根据用户自己的需求进行更改
	ctx := context.Background()
	//ctx2,err:=context.WithDeadline(context.Background(),time.Now().Add(3*time.Hour))用于设置过期时间
	// 创建这个bucket
	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		// 检测这个bucket是否已经存在
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}
	data, err := c.FormFile("data")
	if err != nil {
		logrus.Info(err)
	}
	filename := filepath.Base(data.Filename) //Base 返回 path 的最后一个元素。 在提取最后一个元素之前，将删除尾随路径分隔符。 如果路径为空，则 Base 返回“.”。 如果路径完全由分隔符组成，则 Base 返回单个分隔符
	dst := "C:\\Users\\0\\Downloads\\" + filename
	err = c.SaveUploadedFile(data, dst)
	//ToDo:这里的dst为目标路径，这是一个在电脑本地的目录用于存放发送到minio服务器的文件
	if err != nil {
		logrus.Info(err)
	}
	/*
			ToDo:
			func UploadFile(c *gin.Context) {
			    // 获取上传的文件
			    file, _ := c.FormFile("data")
			    // 指定文件应该保存的路径
			    dst := "/path/to/save/" + file.Filename
			    // 保存文件
			    err := c.SaveUploadedFile(file, dst)
			    if err != nil {
			        log.Println(err)
			    }
			}
		在上面的代码中，我们首先从表单中获取上传的文件，然后指定文件应该保存的路径，最后调用c.SaveUploadedFile方法将文件保存到指定的路径。
		请注意，/path/to/save/应该替换为实际的路径，这个路径应该存在，并且服务器应该有权限写入这个路径。
	*/
	logrus.Info(filename)
	logrus.Info(dst)
	contentType := "multipart/form-data"
	//contentType := "multipart/form-data"这行代码定义了一个变量contentType，它的值是"multipart/form-data"。"multipart/form-data"是一种HTTP请求的内容类型，通常用于上传文件1

	fileInfo, err := os.Stat(dst)
	/*
		在Go语言中，os.Stat(dst)是一个函数调用，它返回一个描述文件dst的os.FileInfo值和一个error值12。
		如果调用成功，os.FileInfo值包含了文件的名称、大小、权限等信息12，error值为nil12。如果调用失败，例如文件不存在，os.FileInfo值为nil，error值描述了发生的错误12。
	*/
	if err == os.ErrNotExist { //这这是在对错误进行处理
		log.Printf("%s目标文件不存在", dst)
	}
	f, err := os.Open(dst)
	/*
		ToDo:
		在Go语言中，os.Open(dst)是一个函数调用，它打开一个文件并返回一个文件句柄和一个错误值12。如果调用成功，文件句柄可以用来读取文件的内容，错误值为nil12。如果调用失败，例如文件不存在，文件句柄为nil，错误值描述了发生的错误12。
		所以，f, err := os.Open(dst)这行代码的作用是打开文件dst，并将文件句柄存储在f变量中，将可能的错误存储在err变量中12
	*/
	if err != nil {
		return
	}
	//其过程为先将发送请求的文件保存到指定的地址dist，然后再从dist发送文件到minio服务器
	uploadInfo, err := minioClient.PutObject(ctx, bucketName, filename, f, fileInfo.Size(), minio.PutObjectOptions{ContentType: contentType}) //一个包含上传选项的结构，例如文件的内容类型.
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Successfully uploaded %s of size %d\n", filename, uploadInfo.Size)
}

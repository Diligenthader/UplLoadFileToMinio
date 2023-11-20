package main

import (
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"v1/routes"
)

func InitMinioClient() *minio.Client {
	// 基本的配置信息
	endpoint := "127.0.0.1:9000"
	accessKeyID := "admin"
	secretAccessKey := "admin12345"
	//对于这个用户名和密码需要自己现在minio服务端设置好，最后在这个函数内初始化时，就可以输入对应的用户名和密码.
	// 初始化一个minio客户端对象
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		log.Fatalf("初始化MinioClient错误：%s", err.Error())
	}
	return minioClient
}
func main() {
	/*	// 创建客户端
		minioClient := InitMinioClient()
		// bucket名称
		bucketName := "mypic"
		ctx := context.Background()
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
		}*/
	r := gin.Default()
	r = routes.SetRoute(r)
	r.Run()
	/*	// 需要上传文件的基本信息
		//objectName1 := "视频.mp4"
		objectName2 := "Picture.jpg"
		filePath2 := "C:\\Users\\0\\Pictures\\Screenshots\\Github.png"
		//filePath1 := "C:\\Users\\0\\Downloads\\Video\\E.mp4"
		contentType := "multipart/form-data"
		//fPath := filepath.Join(filePath, objectName)
		fileInfo, err := os.Stat(filePath2)
		if err == os.ErrNotExist {
			log.Printf("%s目标文件不存在", filePath2)
		}
		f, err := os.Open(filePath2)
		if err != nil {
			return
		}
		uploadInfo, err := minioClient.PutObject(ctx, bucketName, objectName2, f, fileInfo.Size(), minio.PutObjectOptions{ContentType: contentType})
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("Successfully uploaded %s of size %d\n", objectName2, uploadInfo.Size)*/
}

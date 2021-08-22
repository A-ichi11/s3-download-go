package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// 画像のPath
var filePath = "./sakura.jpeg"

// S3のバケット名
var bucket = "test-bucket-0814"

// key S3に保存するオブジェクトの名前になります
var key = "image/sakura"

// awsのリージョン名
var awsRegion = "ap-northeast-1"

func main() {

	// 画像を読み込みます
	imageFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	// 最後に画像ファイルを閉じます
	defer imageFile.Close()

	// sessionを作成します
	newSession := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// S3クライアントを作成します
	svc := s3.New(newSession, &aws.Config{
		Region: aws.String(awsRegion),
	})

	// S3にアップロードする内容をparamsに入れます
	params := &s3.GetObjectInput{
		// Bucket ダウンロードするS3のバケット名
		Bucket: aws.String(bucket),
		// Key ダウンロードするオブジェクト名
		Key: aws.String(key),
	}

	// S3からダウンロードします
	resp, err := svc.GetObject(params)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("S3からダウンロードが完了しました。")
	log.Println(resp)
}

// // GetS3Object は指定ファイルを取得する
// func GetS3Object(bucket, key string) (*bytes.Buffer, error) {
// 	svc := NewS3Client(GetS3LogBucketName())

// 	resp, err := svc.Client.GetObject(&s3.GetObjectInput{
// 		Bucket: aws.String(bucket),
// 		Key:    aws.String(key),
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	buf := new(bytes.Buffer)
// 	buf.ReadFrom(resp.Body)

// 	return buf, nil
// }

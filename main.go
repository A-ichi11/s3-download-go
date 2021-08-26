package main

import (
	"bytes"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// 作成する画像ファイルの名前
var filePath = "./sakura.jpeg"

// S3のバケット名
var bucket = "test-bucket-0814"

// S3のオブジェクト名
var key = "image/sakura.jpeg"

// awsのリージョン名
var awsRegion = "ap-northeast-1"

func main() {

	// ファイルを作成します。
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// sessionを作成します
	newSession := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// S3クライアントを作成します
	svc := s3.New(newSession, &aws.Config{
		Region: aws.String(awsRegion),
	})

	// S3からダウンロードする内容をparamsに入れます
	params := &s3.GetObjectInput{
		// Bucket ダウンロードするS3のバケット名
		Bucket: aws.String(bucket),
		// Key ダウンロードするオブジェクト名
		Key: aws.String(key),
	}

	// S3からダウンロードします
	getImage, err := svc.GetObject(params)
	if err != nil {
		log.Fatal(err)
	}

	//getImageをbytes.Buffer型に変換します
	buf := new(bytes.Buffer)
	buf.ReadFrom(getImage.Body)

	// ファイルに書き込みします。
	_, err = file.Write(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("S3からダウンロードが完了しました。")
}

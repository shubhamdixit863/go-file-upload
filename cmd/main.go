package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"goFileService/internal/api/handlers"
	"goFileService/internal/application"
	"log"
	"os"
)

func InitAws() *session.Session {
	config := aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_ACCESS_SECRET"), ""),
	}
	sess := session.Must(session.NewSession(&config))
	return sess

}

func main() {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	InitAws()

	r := gin.Default()
	config := cors.DefaultConfig()

	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Authorization"}
	fs := application.FileUploadServiceImpl{Aws: InitAws()}
	handler := handlers.Handler{
		FileUplaodService: &fs,
	}
	r.Use(cors.New(config))
	r.POST("/upload", handler.UploadFile)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"goFileService/internal/api/handlers"
	"goFileService/internal/application"
	"os"
)

func InitAws() *session.Session {
	config := aws.Config{
		Region:      aws.String(os.Getenv("REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_ACCESS_SECRET"), ""),
	}
	sess := session.Must(session.NewSession(&config))
	return sess

}

func main() {
	/*
		err := godotenv.Load("../.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}

	*/
	InitAws()

	r := gin.Default()
	config := cors.DefaultConfig()

	config.AllowOrigins = []string{"*"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Authorization", "Content-Type", "x-requested-with", "X-Requested-With"}
	fs := application.FileUploadServiceImpl{Aws: InitAws()}
	handler := handlers.Handler{
		FileUplaodService: &fs,
	}
	r.Use(cors.New(config))
	r.GET("/", handler.Status)
	r.POST("/upload", handler.UploadFile)
	r.Run("0.0.0.0:8090") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

package application

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestInitAws(t *testing.T) {
	env := os.Getenv("ENV")
	if env == "DEV" {
		err := godotenv.Load("../../.env")
		assert.Nil(t, err)
	}

	config := aws.Config{
		Region:      aws.String(os.Getenv("REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_ACCESS_SECRET"), ""),
	}
	sess := session.Must(session.NewSession(&config))
	assert.NotNil(t, sess)

}

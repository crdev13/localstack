package connection

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"
)

func CreateS3ClientWithoutCredentials(endpoint string) *s3.S3 {

	s := session.Must(session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials("foo", "var", ""),
		S3ForcePathStyle: aws.Bool(true),
		Region:           aws.String(endpoints.UsWest2RegionID),
		Endpoint:         aws.String(endpoint),
	}))

	return s3.New(s, &aws.Config{})

}

func CreateDynamoDBClientWithoutCredentials(endpoint string) *dynamodb.DynamoDB {

	s := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials("foo", "var", ""),
		Region:      aws.String(endpoints.UsWest2RegionID),
		Endpoint:    aws.String(endpoint),
	}))

	return dynamodb.New(s, &aws.Config{})

}

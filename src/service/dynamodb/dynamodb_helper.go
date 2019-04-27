package dynamodb

import (
	"github.com/agustin-sarasua/user-system/src/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	ddb "github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDBHelper interface {
}

type dynamoDBHelperImpl struct{}

func NewDynamoDBHelper() *dynamoDBHelperImpl {
	return &dynamoDBHelperImpl{}
}

func getDynamoDBDocumentClient() *ddb.DynamoDB {
	cfg := config.Configure("DEVELOPMENT")
	sess := session.Must(session.NewSession(&aws.Config{
		MaxRetries: aws.Int(3),
	}))
	return ddb.New(sess, aws.NewConfig().WithRegion(cfg.AwsRegion))
}

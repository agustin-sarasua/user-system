package dynamodb

import (
	"fmt"

	"github.com/agustin-sarasua/user-system/src/config"
	"github.com/agustin-sarasua/user-system/src/logger"
	"github.com/agustin-sarasua/user-system/src/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	ddb "github.com/aws/aws-sdk-go/service/dynamodb"
	ddba "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type dynamoDBHelperImpl struct {
	TableName   string
	TableExists bool
}

func NewDynamoDBHelper() *dynamoDBHelperImpl {
	return &dynamoDBHelperImpl{}
}

func (helper *dynamoDBHelperImpl) PutItem(user *model.User, creds *credentials.Credentials) *ddb.PutItemOutput {
	client := helper.getDynamoDBDocumentClient(creds)
	uMap, err := ddba.MarshalMap(user)
	if err != nil {
	}
	input := &ddb.PutItemInput{
		Item:      uMap,
		TableName: aws.String(helper.TableName),
	}
	out, err := client.PutItem(input)
	if err != nil {
	}
	return out
}

func (helper *dynamoDBHelperImpl) tableExists(creds *credentials.Credentials) *ddb.DescribeTableOutput {
	client := helper.getDynamoDBDocumentClient(creds)
	input := &ddb.DescribeTableInput{
		TableName: aws.String(helper.TableName),
	}
	out, err := client.DescribeTable(input)
	if err != nil {
		return nil
	}
	return out
}

func (helper *dynamoDBHelperImpl) createTable(creds *credentials.Credentials) *ddb.CreateTableOutput {
	client := helper.getDynamoDBDocumentClient(creds)
	input := &ddb.DescribeTableInput{
		TableName: aws.String(helper.TableName),
	}
	_, err := client.DescribeTable(input)
	if err == nil {
		logger.Info("Table already exists", nil)
	} else {
		input := &ddb.CreateTableInput{}
		input.SetTableName(helper.TableName)
		tc, err := client.CreateTable(input)
		if err != nil {
			logger.Error(fmt.Sprintf("Unable to create Table %s", helper.TableName), err)
			return nil
		} else {
			logger.Info(fmt.Sprintf("Table %s created", helper.TableName), nil)
			return tc
		}
	}
	return nil
}

func (helper *dynamoDBHelperImpl) getDynamoDBDocumentClient(creds *credentials.Credentials) *ddb.DynamoDB {
	sess := session.Must(session.NewSession(&aws.Config{
		MaxRetries: aws.Int(3),
	}))

	svc := ddb.New(sess, aws.NewConfig().WithRegion(config.Cfg.AwsRegion).WithCredentials(creds))
	return svc
}

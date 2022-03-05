package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/crdev13/localstack/golang/connection"
)

const DynamoDBEndpoint = "http://localhost:8000"

func main() {
	svc := connection.CreateDynamoDBClientWithoutCredentials(DynamoDBEndpoint)

	// createTable(svc)
	listTables(svc)
	createTable(svc)
	listTables(svc)
}

func listTables(svc *dynamodb.DynamoDB) {
	input := &dynamodb.ListTablesInput{}
	fmt.Printf("Tables:\n")

	// Get the list of tables
	result, err := svc.ListTables(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}
	for _, n := range result.TableNames {
		fmt.Println(*n)
	}

	// assign the last read tablename as the start for our next call to the ListTables function
	// the maximum number of table names returned in a call is 100 (default), which requires us to make
	// multiple calls to the ListTables function to retrieve all table names
	input.ExclusiveStartTableName = result.LastEvaluatedTableName

}

func createTable(svc *dynamodb.DynamoDB) {
	// snippet-start:[dynamodb.go.create_table.session]
	tableName := "Movies"

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Year"),
				AttributeType: aws.String("N"),
			},
			{
				AttributeName: aws.String("Title"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Year"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("Title"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}

	_, err := svc.CreateTable(input)
	if err != nil {
		log.Println("Got error calling CreateTable: %s", err)
	}

	fmt.Println("Created the table", tableName)
}

package repositories

import (
	"github.com/alejo-lapix/products-quote-go/pkg/quotes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type DynamoDBQuoteRepository struct {
	DynamoDB  *dynamodb.DynamoDB
	tableName *string
}

func NewDynamoDBQuoteRepository(db *dynamodb.DynamoDB) *DynamoDBQuoteRepository {
	return &DynamoDBQuoteRepository{
		DynamoDB:  db,
		tableName: aws.String("quotes"),
	}
}

func (repository *DynamoDBQuoteRepository) Store(quote *quotes.Quote) error {
	item, err := dynamodbattribute.MarshalMap(quote)

	if err != nil {
		return err
	}

	_, err = repository.DynamoDB.PutItem(&dynamodb.PutItemInput{
		ConditionExpression: aws.String("attribute_not_exists(id)"),
		Item:                item,
		TableName:           repository.tableName,
	})

	return err
}

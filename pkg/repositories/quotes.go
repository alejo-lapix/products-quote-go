package repositories

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/alejo-lapix/products-quote-go/pkg/quotes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"time"
)

type paginationKey struct {
	ID           *string `json:"id"`
	MonthAndYear *string `json:"monthAndYear"`
	CreatedAt    *string `json:"createdAt"`
}

type DynamoDBQuoteRepository struct {
	DynamoDB    *dynamodb.DynamoDB
	tableName   *string
	pageEncoder PageEncoder
}

func NewDynamoDBQuoteRepository(db *dynamodb.DynamoDB) *DynamoDBQuoteRepository {
	return &DynamoDBQuoteRepository{
		DynamoDB:    db,
		tableName:   aws.String("quotes"),
		pageEncoder: &DynamoEncoder{},
	}
}

func (repository *DynamoDBQuoteRepository) Find(ID *string) (*quotes.Quote, error) {
	item := &quotes.Quote{}
	output, err := repository.DynamoDB.GetItem(&dynamodb.GetItemInput{
		Key:       map[string]*dynamodb.AttributeValue{"id": {S: ID}},
		TableName: repository.tableName,
	})

	if err != nil {
		return nil, err
	}

	err = dynamodbattribute.UnmarshalMap(output.Item, item)

	if err != nil {
		return nil, err
	}

	return item, nil
}

func (repository *DynamoDBQuoteRepository) Paginate(year, month, lastEvaluatedKey *string) ([]*quotes.Quote, *string, error) {
	var key *string
	quoteList := make([]*quotes.Quote, 0)
	input := dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{":keyValue": {S: aws.String(fmt.Sprintf("%s-%s", *year, *month))}},
		IndexName:                 aws.String("monthAndYear-createdAt-index"),
		KeyConditionExpression:    aws.String("monthAndYear = :keyValue"),
		Limit:                     aws.Int64(20),
		ScanIndexForward:          aws.Bool(false),
		TableName:                 repository.tableName,
	}

	if lastEvaluatedKey != nil {
		key, err := repository.pageEncoder.Decode(*lastEvaluatedKey)

		if err != nil {
			return nil, nil, err
		}

		input.SetExclusiveStartKey(key)
	}

	items, err := repository.DynamoDB.Query(&input)

	if err != nil {
		return nil, nil, err
	}

	err = dynamodbattribute.UnmarshalListOfMaps(items.Items, &quoteList)

	if err != nil {
		return nil, nil, err
	}

	if items.LastEvaluatedKey != nil {
		key, err = repository.pageEncoder.Encode(items.LastEvaluatedKey)

		if err != nil {
			return nil, nil, err
		}
	}

	return quoteList, key, nil
}

func (repository *DynamoDBQuoteRepository) All() ([]*quotes.Quote, error) {
	return make([]*quotes.Quote, 0), nil
}

func (repository *DynamoDBQuoteRepository) Store(quote *quotes.Quote) error {
	item, err := dynamodbattribute.MarshalMap(quote)

	if err != nil {
		return err
	}

	createdAtTime, err := time.Parse(time.RFC3339, *quote.CreatedAt)

	if err != nil {
		return err
	}

	item["monthAndYear"] = &dynamodb.AttributeValue{
		S: aws.String(fmt.Sprintf("%d-%d", createdAtTime.Year(), int(createdAtTime.Month()))),
	}

	_, err = repository.DynamoDB.PutItem(&dynamodb.PutItemInput{
		ConditionExpression: aws.String("attribute_not_exists(id)"),
		Item:                item,
		TableName:           repository.tableName,
	})

	return err
}

type PageEncoder interface {
	Encode(map[string]*dynamodb.AttributeValue) (*string, error)
	Decode(string) (map[string]*dynamodb.AttributeValue, error)
}

type DynamoEncoder struct{}

func (encoder DynamoEncoder) Encode(input map[string]*dynamodb.AttributeValue) (*string, error) {
	temporalKey := &paginationKey{}
	err := dynamodbattribute.UnmarshalMap(input, temporalKey)

	if err != nil {
		return nil, EncodeError{originalError: err, isMarshallError: true}
	}

	content, err := json.Marshal(temporalKey)

	if err != nil {
		return nil, EncodeError{originalError: err, isJSonError: true}
	}

	key := base64.StdEncoding.EncodeToString([]byte(content))

	return &key, nil
}

func (encoder DynamoEncoder) Decode(input string) (map[string]*dynamodb.AttributeValue, error) {
	const errorMessage = "La página indicada no es valida"
	jsonItem, err := base64.StdEncoding.DecodeString(input)

	if err != nil {
		return nil, EncodeError{originalError: err, isBase64Error: true, message: errorMessage}
	}

	keyFromJson := &paginationKey{}
	err = json.Unmarshal([]byte(jsonItem), keyFromJson)

	if err != nil {
		return nil, EncodeError{originalError: err, isMarshallError: true}
	}

	key, err := dynamodbattribute.MarshalMap(keyFromJson)

	if err != nil {
		return nil, EncodeError{originalError: err, isMarshallError: true}
	}

	return key, nil
}

type EncodeError struct {
	message                                     string
	isJSonError, isBase64Error, isMarshallError bool
	originalError                               error
}

func (encodeError EncodeError) Error() string {
	return encodeError.message
}

func (encodeError EncodeError) OriginalError() error {
	return encodeError.originalError
}

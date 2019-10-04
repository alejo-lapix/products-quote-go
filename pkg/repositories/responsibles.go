package repositories

import (
	"github.com/alejo-lapix/products-quote-go/pkg/responsibles"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type DynamoDBUserRepository struct {
	DynamoDB  *dynamodb.DynamoDB
	tableName *string
}

func NewDynamoDBUserRepository(db *dynamodb.DynamoDB) *DynamoDBUserRepository {
	return &DynamoDBUserRepository{
		DynamoDB:  db,
		tableName: aws.String("users"),
	}
}

func (repository *DynamoDBUserRepository) Find(id *string) (*responsibles.User, error) {
	item := &responsibles.User{}
	output, err := repository.DynamoDB.GetItem(&dynamodb.GetItemInput{
		Key:       map[string]*dynamodb.AttributeValue{"id": {S: id}},
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

func (repository *DynamoDBUserRepository) FindMany(items []*string) ([]*responsibles.User, error) {
	list := make([]*responsibles.User, len(items))
	keys := make([]map[string]*dynamodb.AttributeValue, len(items))

	for index, item := range items {
		keys[index] = map[string]*dynamodb.AttributeValue{"id": {S: item}}
	}

	output, err := repository.DynamoDB.BatchGetItem(&dynamodb.BatchGetItemInput{
		RequestItems:           map[string]*dynamodb.KeysAndAttributes{*repository.tableName: {Keys: keys}},
		ReturnConsumedCapacity: nil,
	})

	if err != nil {
		return nil, err
	}

	err = dynamodbattribute.UnmarshalListOfMaps(output.Responses[*repository.tableName], &list)

	if err != nil {
		return nil, err
	}

	return list, nil
}

func (repository *DynamoDBUserRepository) FindByCategoryAndZone(categoryID, zoneID *string) ([]*responsibles.User, error) {
	items := make([]*responsibles.User, 0)
	output, err := repository.DynamoDB.Scan(&dynamodb.ScanInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":categoryId": {S: categoryID},
			":zoneId":      {S: zoneID},
		},
		FilterExpression: aws.String("contains(categoryIds, :categoryId) AND contains(zoneIds, :zoneId)"),
		TableName:        repository.tableName,
	})

	if err != nil {
		return nil, err
	}

	err = dynamodbattribute.UnmarshalListOfMaps(output.Items, &items)

	if err != nil {
		return nil, err
	}

	return items, nil
}

func (repository *DynamoDBUserRepository) Store(user *responsibles.User) error {
	item, err := dynamodbattribute.MarshalMap(user)

	if err != nil {
		return err
	}

	_, err = repository.DynamoDB.PutItem(&dynamodb.PutItemInput{
		Item:      item,
		TableName: repository.tableName,
	})

	return err
}

func (repository *DynamoDBUserRepository) Remove(ID *string) error {
	_, err := repository.DynamoDB.DeleteItem(&dynamodb.DeleteItemInput{
		Key:       map[string]*dynamodb.AttributeValue{"id": {S: ID}},
		TableName: repository.tableName,
	})

	return err
}

func (repository *DynamoDBUserRepository) All() ([]*responsibles.User, error) {
	items := make([]*responsibles.User, 0)
	output, err := repository.DynamoDB.Scan(&dynamodb.ScanInput{TableName: repository.tableName})

	if err != nil {
		return nil, err
	}

	err = dynamodbattribute.UnmarshalListOfMaps(output.Items, &items)

	if err != nil {
		return nil, err
	}

	return items, nil
}

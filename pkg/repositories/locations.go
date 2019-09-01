package repositories

import (
	"github.com/alejo-lapix/products-quote-go/pkg/locations"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type DynamoDBCountryRepository struct {
	DynamoDB  *dynamodb.DynamoDB
	tableName *string
}

func NewDynamoDBCountryRepository(db *dynamodb.DynamoDB) *DynamoDBCountryRepository {
	return &DynamoDBCountryRepository{
		DynamoDB:  db,
		tableName: aws.String("countries"),
	}
}

func NewDynamoDBZoneRepository(db *dynamodb.DynamoDB) *DynamoDBZoneRepository {
	return &DynamoDBZoneRepository{
		DynamoDB:                 db,
		tableName:                aws.String("zones"),
		zonesByProductRepository: NewDynamoDBZoneByProductIDRepository(db),
	}
}

func (repository *DynamoDBCountryRepository) Find(ID *string) (*locations.Country, error) {
	item := &locations.Country{}
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

func (repository *DynamoDBCountryRepository) All() ([]*locations.Country, error) {
	items := make([]*locations.Country, 0)
	output, err := repository.DynamoDB.Scan(&dynamodb.ScanInput{
		TableName: repository.tableName,
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

func (repository *DynamoDBCountryRepository) Remove(ID *string) error {
	_, err := repository.DynamoDB.DeleteItem(&dynamodb.DeleteItemInput{
		Key:       map[string]*dynamodb.AttributeValue{"id": {S: ID}},
		TableName: repository.tableName,
	})

	return err
}

func (repository *DynamoDBCountryRepository) Store(country *locations.Country) error {
	item, err := dynamodbattribute.MarshalMap(country)

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
func (repository *DynamoDBCountryRepository) Update(ID *string, country *locations.Country) error {
	item, err := dynamodbattribute.MarshalMap(country)

	if err != nil {
		return err
	}

	_, err = repository.DynamoDB.PutItem(&dynamodb.PutItemInput{
		ConditionExpression:       aws.String("id = :id"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{":id": {S: ID}},
		Item:                      item,
		TableName:                 repository.tableName,
	})

	return err
}

type DynamoDBZoneRepository struct {
	tableName                *string
	DynamoDB                 *dynamodb.DynamoDB
	zonesByProductRepository locations.ZonesByProductIDRepository
}

func (repository *DynamoDBZoneRepository) Find(ID *string) (*locations.Zone, error) {
	item := &locations.Zone{}
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

func (repository *DynamoDBZoneRepository) FindMany(items []*string) ([]*locations.Zone, error) {
	list := make([]*locations.Zone, len(items))
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

func (repository *DynamoDBZoneRepository) FindByProduct(ID *string) ([]*locations.Zone, error) {
	item, err := repository.zonesByProductRepository.Find(ID)

	if err != nil {
		return nil, err
	}

	return repository.FindMany(item.ZoneIDs)
}

func (repository *DynamoDBZoneRepository) FindByCountry(ID *string) ([]*locations.Zone, error) {
	list := make([]*locations.Zone, 0)

	output, err := repository.DynamoDB.Query(&dynamodb.QueryInput{
		IndexName:                 aws.String("countryId-index"),
		KeyConditionExpression:    aws.String("countryId = :countryId"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{"countryId": {S: ID}},
		TableName:                 repository.tableName,
	})

	if err != nil {
		return nil, err
	}

	err = dynamodbattribute.UnmarshalListOfMaps(output.Items, &list)

	if err != nil {
		return nil, err
	}

	return list, nil
}

func (repository *DynamoDBZoneRepository) Remove(ID *string) error {
	_, err := repository.DynamoDB.DeleteItem(&dynamodb.DeleteItemInput{
		Key:       map[string]*dynamodb.AttributeValue{"id": {S: ID}},
		TableName: repository.tableName,
	})

	return err
}

func (repository *DynamoDBZoneRepository) Store(zone *locations.Zone) error {
	item, err := dynamodbattribute.MarshalMap(zone)

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
func (repository *DynamoDBZoneRepository) Update(id *string, zone *locations.Zone) error {
	item, err := dynamodbattribute.MarshalMap(zone)

	if err != nil {
		return err
	}

	_, err = repository.DynamoDB.PutItem(&dynamodb.PutItemInput{
		ConditionExpression:       aws.String("id = :id"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{":id": {S: id}},
		Item:                      item,
		TableName:                 repository.tableName,
	})

	return err
}

func (repository *DynamoDBZoneRepository) ProductsIDsByZone(ID *string) ([]*string, error) {
	list, err := repository.zonesByProductRepository.FindByZone(ID)

	if err != nil {
		return nil, err
	}

	result := make([]*string, len(list))

	for index, item := range list {
		result[index] = item.ProductID
	}

	return result, nil
}

type dynamoDBZonesByProductIDRepository struct {
	DynamoDB  *dynamodb.DynamoDB
	tableName *string
}

func NewDynamoDBZoneByProductIDRepository(db *dynamodb.DynamoDB) *dynamoDBZonesByProductIDRepository {
	return &dynamoDBZonesByProductIDRepository{
		DynamoDB:  db,
		tableName: aws.String("zonesByProductId"),
	}
}

func (repository *dynamoDBZonesByProductIDRepository) Store(element *locations.ZonesByProductID) error {
	item, err := dynamodbattribute.MarshalMap(element)

	if err != nil {
		return err
	}

	_, err = repository.DynamoDB.PutItem(&dynamodb.PutItemInput{
		Item:      item,
		TableName: repository.tableName,
	})

	return err
}

func (repository *dynamoDBZonesByProductIDRepository) Remove(productID *string) error {
	_, err := repository.DynamoDB.DeleteItem(&dynamodb.DeleteItemInput{
		Key:       map[string]*dynamodb.AttributeValue{"productId": {S: productID}},
		TableName: repository.tableName,
	})

	return err
}

func (repository *dynamoDBZonesByProductIDRepository) Find(productID *string) (*locations.ZonesByProductID, error) {
	item := &locations.ZonesByProductID{}
	output, err := repository.DynamoDB.GetItem(&dynamodb.GetItemInput{
		Key:       map[string]*dynamodb.AttributeValue{"productId": {S: productID}},
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

func (repository *dynamoDBZonesByProductIDRepository) FindByZone(ID *string) ([]*locations.ZonesByProductID, error) {
	items := make([]*locations.ZonesByProductID, 0)
	output, err := repository.DynamoDB.Scan(&dynamodb.ScanInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{":zoneId": {S: ID}},
		FilterExpression:          aws.String("zoneIds CONTAINS :zoneId"),
		TableName:                 repository.tableName,
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

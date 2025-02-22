package repositories

import (
	"github.com/alejo-lapix/products-quote-go/pkg/locations"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"reflect"
	"testing"
)

func TestDynamoDBCountryRepository_All(t *testing.T) {
	type fields struct {
		DynamoDB  *dynamodb.DynamoDB
		tableName *string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*locations.Country
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &DynamoDBCountryRepository{
				DynamoDB:  tt.fields.DynamoDB,
				tableName: tt.fields.tableName,
			}
			got, err := repository.All()
			if (err != nil) != tt.wantErr {
				t.Errorf("All() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("All() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDynamoDBCountryRepository_Find(t *testing.T) {
	type fields struct {
		DynamoDB  *dynamodb.DynamoDB
		tableName *string
	}
	type args struct {
		ID *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *locations.Country
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &DynamoDBCountryRepository{
				DynamoDB:  tt.fields.DynamoDB,
				tableName: tt.fields.tableName,
			}
			got, err := repository.Find(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDynamoDBCountryRepository_Remove(t *testing.T) {
	type fields struct {
		DynamoDB  *dynamodb.DynamoDB
		tableName *string
	}
	type args struct {
		ID *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &DynamoDBCountryRepository{
				DynamoDB:  tt.fields.DynamoDB,
				tableName: tt.fields.tableName,
			}
			if err := repository.Remove(tt.args.ID); (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDynamoDBCountryRepository_Store(t *testing.T) {
	type fields struct {
		DynamoDB  *dynamodb.DynamoDB
		tableName *string
	}
	type args struct {
		country *locations.Country
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &DynamoDBCountryRepository{
				DynamoDB:  tt.fields.DynamoDB,
				tableName: tt.fields.tableName,
			}
			if err := repository.Store(tt.args.country); (err != nil) != tt.wantErr {
				t.Errorf("Store() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDynamoDBCountryRepository_Update(t *testing.T) {
	type fields struct {
		DynamoDB  *dynamodb.DynamoDB
		tableName *string
	}
	type args struct {
		ID      *string
		country *locations.Country
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &DynamoDBCountryRepository{
				DynamoDB:  tt.fields.DynamoDB,
				tableName: tt.fields.tableName,
			}
			if err := repository.Update(tt.args.ID, tt.args.country); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDynamoDBZoneRepository_Find(t *testing.T) {
	type fields struct {
		tableName                *string
		DynamoDB                 *dynamodb.DynamoDB
	}
	type args struct {
		ID *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *locations.Zone
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &DynamoDBZoneRepository{
				tableName:                tt.fields.tableName,
				DynamoDB:                 tt.fields.DynamoDB,
			}
			got, err := repository.Find(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDynamoDBZoneRepository_FindByCountry(t *testing.T) {
	type fields struct {
		tableName                *string
		DynamoDB                 *dynamodb.DynamoDB
	}
	type args struct {
		ID *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*locations.Zone
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &DynamoDBZoneRepository{
				tableName:                tt.fields.tableName,
				DynamoDB:                 tt.fields.DynamoDB,
			}
			got, err := repository.FindByCountry(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByCountry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByCountry() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDynamoDBZoneRepository_FindMany(t *testing.T) {
	type fields struct {
		tableName                *string
		DynamoDB                 *dynamodb.DynamoDB
	}
	type args struct {
		items []*string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*locations.Zone
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &DynamoDBZoneRepository{
				tableName:                tt.fields.tableName,
				DynamoDB:                 tt.fields.DynamoDB,
			}
			got, err := repository.FindMany(tt.args.items)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindMany() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindMany() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDynamoDBZoneRepository_Remove(t *testing.T) {
	type fields struct {
		tableName                *string
		DynamoDB                 *dynamodb.DynamoDB
	}
	type args struct {
		ID *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &DynamoDBZoneRepository{
				tableName:                tt.fields.tableName,
				DynamoDB:                 tt.fields.DynamoDB,
			}
			if err := repository.Remove(tt.args.ID); (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDynamoDBZoneRepository_Store(t *testing.T) {
	type fields struct {
		tableName                *string
		DynamoDB                 *dynamodb.DynamoDB
	}
	type args struct {
		zone *locations.Zone
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &DynamoDBZoneRepository{
				tableName:                tt.fields.tableName,
				DynamoDB:                 tt.fields.DynamoDB,
			}
			if err := repository.Store(tt.args.zone); (err != nil) != tt.wantErr {
				t.Errorf("Store() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDynamoDBZoneRepository_Update(t *testing.T) {
	type fields struct {
		tableName                *string
		DynamoDB                 *dynamodb.DynamoDB
	}
	type args struct {
		id   *string
		zone *locations.Zone
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := &DynamoDBZoneRepository{
				tableName:                tt.fields.tableName,
				DynamoDB:                 tt.fields.DynamoDB,
			}
			if err := repository.Update(tt.args.id, tt.args.zone); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func dynamoDBInstance() *dynamodb.DynamoDB {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})

	if err != nil {
		panic(err.Error())
	}

	return dynamodb.New(sess)
}

package repositories

import (
	"github.com/alejo-lapix/products-quote-go/pkg/responsibles"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"reflect"
	"testing"
)

func responsibleRepository() *DynamoDBUserRepository {
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-1")})
	if err != nil {
		panic(err)
	}

	return NewDynamoDBUserRepository(dynamodb.New(sess))
}

func TestDynamoDBUserRepository_FindByCategoryAndZone(t *testing.T) {
	type fields struct {
		Repository responsibles.UserRepository
	}
	type args struct {
		categoryID *string
		zoneID     *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*responsibles.User
		wantErr bool
	}{
		{
			name:    "Can get the users by the category and zone id",
			fields:  fields{responsibleRepository()},
			args:    args{categoryID: aws.String("any"), zoneID: aws.String("any")},
			want:    []*responsibles.User{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := tt.fields.Repository
			got, err := repository.FindByCategoryAndZone(tt.args.categoryID, tt.args.zoneID)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByCategoryAndZone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByCategoryAndZone() got = %v, want %v", got, tt.want)
			}
		})
	}
}

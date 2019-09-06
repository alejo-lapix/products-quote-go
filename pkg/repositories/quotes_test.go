package repositories

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"testing"
)

func repository() *DynamoDBQuoteRepository {
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-1")})
	if err != nil {
		panic(err)
	}

	return NewDynamoDBQuoteRepository(dynamodb.New(sess))
}

func TestDynamoDBQuoteRepository_Paginate(t *testing.T) {
	type args struct {
		year             *string
		month            *string
		lastEvaluatedKey *string
	}

	arguments := &args{
		year:             aws.String("2019"),
		month:            aws.String("8"),
		lastEvaluatedKey: nil,
	}

	tests := []struct {
		name         string
		args         *args
		wantErr      bool
		want         int
		nextPageNil  bool
		errorMessage string
	}{
		{
			name: "Invalid key provided",
			args: &args{
				year:             aws.String("2019"),
				month:            aws.String("8"),
				lastEvaluatedKey: aws.String("INVALID KEY"),
			},
			wantErr:      true,
			errorMessage: "La página indicada no es valida",
		},
		{
			name:        "Query all the items",
			args:        arguments,
			wantErr:     false,
			want:        2,
			nextPageNil: false,
		},
		{
			name:        "Query all the items left",
			args:        arguments,
			wantErr:     false,
			want:        1,
			nextPageNil: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repository := repository()
			got, got1, err := repository.Paginate(tt.args.year, tt.args.month, tt.args.lastEvaluatedKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("Paginate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				if tt.errorMessage != err.Error() {
					t.Errorf("Paginate() error = %v, wantErr %v", err.Error(), tt.errorMessage)
					return
				}

				return
			}

			if len(got) != tt.want {
				t.Errorf("Paginate() len(want)= %v, want %d", len(got), tt.want)
				return
			}

			if got1 == nil && !tt.nextPageNil {
				t.Errorf("Paginate() got1= nil, want a *string")
				return
			}

			tt.args.lastEvaluatedKey = got1
		})
	}
}

package dynamo_db

import "github.com/azam-akram/aws-apigateway-lambda-demo-go/model"

type Handler interface {
	Save(book *model.MyBook) error
	Update(book *model.MyBook) error
	UpdateAttributeByID(id, key, value string) error
	GetByID(id string) (*model.MyBook, error)
	DeleteByID(id string) error
}

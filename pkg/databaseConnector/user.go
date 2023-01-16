package databaseConnector

import (
	"context"

	"github.com/vishalrana9915/demo_app/pkg/constant"
)

// save record into database
func CreateNewUser(data interface{}) interface{} {

	result, err := Adapter.db.Collection(constant.USERCOLLECTION).InsertOne(context.TODO(), data)

	if err != nil {
		panic(error.Error(err))
	}

	return result
}

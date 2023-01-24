package databaseConnector

import (
	"context"

	"github.com/vishalrana9915/demo_app/pkg/constant"
	"github.com/vishalrana9915/demo_app/pkg/users/userInterface"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// save record into database
func CreateNewUser(data interface{}) interface{} {

	result, err := Adapter.db.Collection(constant.USERCOLLECTION).InsertOne(context.TODO(), data)

	if err != nil {
		panic(error.Error(err))
	}

	return result
}

// find user from query
func FindUser(query bson.D) (userInterface.Users, string) {
	var userInfo userInterface.Users
	err := Adapter.db.Collection(constant.USERCOLLECTION).FindOne(context.TODO(), query).Decode(&userInfo)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return userInterface.Users{}, constant.USER_NOT_EXIST
		}
		panic(error.Error(err))
	}
	return userInfo, ""
}

// function call to check if user exist with same email or not
func CheckUserExist(email string) bool {

	filter := bson.D{
		{
			"email", email,
		},
	}
	var userInfo userInterface.Users
	err := Adapter.db.Collection(constant.USERCOLLECTION).FindOne(context.TODO(), filter).Decode(&userInfo)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return false
		}
		panic(error.Error(err))
	}

	return true
}

// find user from query
func UpdateUser(query bson.M, update bson.M) (userInterface.Users, string) {

	var userInfo userInterface.Users
	err := Adapter.db.Collection(constant.USERCOLLECTION).FindOneAndUpdate(context.TODO(), query, update).Decode(&userInfo)

	if err != nil {

		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return userInterface.Users{}, constant.USER_NOT_EXIST
		}
		panic(err)
	}
	return userInfo, ""
}

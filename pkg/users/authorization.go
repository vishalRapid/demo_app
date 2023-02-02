package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vishalrana9915/demo_app/pkg/constant"
	"github.com/vishalrana9915/demo_app/pkg/databaseConnector"
	"github.com/vishalrana9915/demo_app/pkg/redisConnector"
	"github.com/vishalrana9915/demo_app/pkg/users/userInterface"
	"github.com/vishalrana9915/demo_app/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// a function handler to take care of user onboarding
func RegisterUser(c *gin.Context) {
	var body userInterface.Users
	decode := json.NewDecoder(c.Request.Body)
	if err := decode.Decode(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constant.PARSING_ERROR})
		c.Abort()
		return
	}

	// checking user in cache
	result := redisConnector.CheckHExist(body.EMAIL)

	if result {
		c.JSON(http.StatusBadRequest, gin.H{"error": constant.USER_EXIST})
		c.Abort()
		return
	}
	// check if user exist
	userExist := databaseConnector.CheckUserExist(body.EMAIL)

	if userExist == true {
		c.JSON(http.StatusBadRequest, gin.H{"error": constant.USER_EXIST})
		c.Abort()
		return
	}

	// hashed password
	hashedPassword := utils.HashPassword(body.PASSWORD)

	// mapping Values
	body.PASSWORD = hashedPassword
	body.ISVERIFIED = false
	body.CREATEDAT = time.Now()
	body.LASTUPDATEDAT = time.Now()
	body.ID = primitive.NewObjectID()

	databaseConnector.CreateNewUser(body)

	// setting up value for current user in cache
	go redisConnector.HSetValue(body.EMAIL, "id", string(body.ID.Hex()))

	responseUser := userInterface.UserProfile{
		ID:            body.ID,
		NAME:          body.NAME,
		EMAIL:         body.EMAIL,
		ISVERIFIED:    body.ISVERIFIED,
		CREATEDAT:     body.CREATEDAT,
		LASTUPDATEDAT: body.LASTUPDATEDAT,
	}

	// Convert struct to JSON
	userJSON, err_json := json.Marshal(responseUser)

	if err_json != nil {
		panic(err_json)
	}

	// setting up value in redis
	go redisConnector.SetValue(body.ID.Hex(), string(userJSON))

	c.JSON(http.StatusCreated, gin.H{"message": constant.SUCCESS, "response": responseUser})

}

// a function handler to take care of user login
func AuthenticateUser(c *gin.Context) {
	var body userInterface.LoginPayload
	decode := json.NewDecoder(c.Request.Body)
	if err := decode.Decode(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constant.PARSING_ERROR})
		c.Abort()
		return
	}

	// checking user in cache
	result := redisConnector.CheckHExist(body.EMAIL)

	if !result {
		c.JSON(http.StatusBadRequest, gin.H{"message": constant.USER_NOT_EXIST})
		c.Abort()
		return
	}

	// need to find user from db
	// check if user exist
	query := bson.D{
		{
			"email", body.EMAIL,
		},
	}
	userInfo, errr := databaseConnector.FindUser(query)
	if errr != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errr})
		c.Abort()
		return
	}

	passwordMatched := utils.ComparePassword(userInfo.PASSWORD, body.PASSWORD)

	if !passwordMatched {
		c.JSON(http.StatusBadRequest, gin.H{"message": constant.INVALID_CREDENTIAL})
		c.Abort()
		return
	}

	// generate Profile
	token := utils.GenerateJWT(userInfo.ID.Hex())

	c.JSON(http.StatusOK, gin.H{"message": constant.SUCCESS,
		"token": token,
		"userInfo": gin.H{
			"name": userInfo.NAME,
		},
	})

}

// fetch user profile
func FetchProfile(c *gin.Context) {
	// Get the JWT token from the headers
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": constant.MISSING_TOKEN})
		c.Abort()
		return
	}

	// validate token
	id, err := utils.ValidateToken(tokenString)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": constant.INVALID_REQUEST})
		c.Abort()
		return
	}

	userResponse, err := redisConnector.GetValue(id)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": constant.INVALID_REQUEST})
		c.Abort()
		return
	}

	var userInfo userInterface.Users
	err_marshal := json.Unmarshal([]byte(userResponse), &userInfo)
	if err_marshal != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": constant.INVALID_REQUEST})
		c.Abort()
		return
	}
	responseUser := userInterface.UserProfile{
		ID:            userInfo.ID,
		NAME:          userInfo.NAME,
		EMAIL:         userInfo.EMAIL,
		ISVERIFIED:    userInfo.ISVERIFIED,
		CREATEDAT:     userInfo.CREATEDAT,
		LASTUPDATEDAT: userInfo.LASTUPDATEDAT,
	}

	c.JSON(http.StatusOK, gin.H{"message": constant.SUCCESS,
		"profile": responseUser})
}

// UpdateTags will update tags for current user
func UpdateTags(c *gin.Context) {
	// checking user required auth
	userId, exist := c.Get("userId")
	if !exist {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "userId not found in context"})
		return
	}

	var profileTags userInterface.UserTags

	decoder := json.NewDecoder(c.Request.Body)

	if err := decoder.Decode(&profileTags); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constant.PARSING_ERROR})
		c.Abort()
		return
	}

	str, ok := userId.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": constant.PARSING_ERROR})
		c.Abort()
		return
	}

	objId, err_parsing := primitive.ObjectIDFromHex(str)

	if err_parsing != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constant.PARSING_ERROR})
		c.Abort()
		return
	}

	query := bson.M{
		"_id": objId,
	}

	update := bson.M{
		"$set": bson.M{"tags": profileTags.TAGS},
	}

	_, err_update := databaseConnector.UpdateUser(query, update)

	if err_update != "" {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": err_update,
		})
		c.Abort()
		return
	}

	tags, _ := json.Marshal(profileTags.TAGS)
	// setting up value in redis
	go redisConnector.SetValue(fmt.Sprintf("user_tags_%s", str), string(tags))

	c.JSON(http.StatusOK, gin.H{
		"message": constant.SUCCESS,
	})

}

// function to update profile for a user
func UpdateProfile(c *gin.Context) {
	userId, exist := c.Get("userId")

	if !exist {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "userId not found in context"})
		return
	}

	// serialise update profile interface
	var userProfile userInterface.UpdateProfile
	decoder := json.NewDecoder(c.Request.Body)

	if err := decoder.Decode(&userProfile); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": constant.INVALID_REQUEST,
		})
		return
	}

	// serialization
	id, err_parsing := primitive.ObjectIDFromHex(userId.(string))

	if err_parsing != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": constant.INVALID_REQUEST,
		})
		return
	}
	// need to update latest information

	query := bson.M{
		"_id": id,
	}

	update := bson.M{
		"$set": userProfile,
	}

	_, err_update := databaseConnector.UpdateUser(query, update)

	if err_update != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": constant.ERROR_UPDATE,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": constant.SUCCESS,
	})

}

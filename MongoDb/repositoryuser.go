package MongoDb

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserService interface {
	Create()
	Insert(user *IUser) bool
	GetAllUser() []IUser
	GenerateLastID() int32
}

type UserServiceStruct struct {
	db        *mongo.Client
	DBName    string
	TabelName string
}

func NewUserServiceStruct(db *mongo.Client, DBName string, TabelName string) IUserService {
	return &UserServiceStruct{db: db, DBName: DBName, TabelName: TabelName}
}

func (tsUser *UserServiceStruct) Create() {
	 
	dbase := tsUser.db.Database(tsUser.DBName)
	CollectionNames, err := dbase.ListCollectionNames(context.Background(), bson.D{})
	if err != nil {
		log.Println("The CollectionName  error :" + err.Error())
	}
	var StateCollectionName = false
	for _, CollectionName := range CollectionNames {
		if CollectionName == tsUser.TabelName {
			StateCollectionName = true
			break
		}
	}
	if !StateCollectionName {
		tsUser.db.Database(tsUser.DBName).CreateCollection(context.Background(), tsUser.TabelName)
	}
}

func (tsUser *UserServiceStruct) GenerateLastID() int32 {
	dbUser := tsUser.db.Database(tsUser.DBName)
	dbCollection := dbUser.Collection(tsUser.TabelName)

	total, err := dbCollection.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		//log.Fatal(" not found : ", err)
		return 1
	}
	return int32(total + 1)
}

func (tsUser *UserServiceStruct) GetAllUser() []IUser {
	dbUser := tsUser.db.Database(tsUser.DBName)
	dbCollection := dbUser.Collection(tsUser.TabelName)

	var IUsers []IUser

	Result, err := dbCollection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Println(" not found : ", err)
	}
	defer Result.Close(context.Background())

	for Result.Next(context.Background()) {

		var entity bson.M
		err := Result.Decode(&entity)
		if err != nil {
			log.Println(" Decode error : ", err)
		}
		userID, err := strconv.ParseInt(fmt.Sprint(entity["userID"]), 10, 64)
		if err != nil {
			log.Println("   error : ", err)
		}

		userFname := fmt.Sprint(entity["userFname"])
		userLname := fmt.Sprint(entity["userLname"])
		userMobile := fmt.Sprint(entity["userMobile"])

		var IUser IUser

		IUser.UserID = int(userID)
		IUser.UserFname = userFname
		IUser.UserLname = userLname
		IUser.UserMobile = userMobile

		IUsers = append(IUsers, IUser)
	}

	return IUsers
}

func (tsUser *UserServiceStruct) Insert(user *IUser) bool {
	dbUser := tsUser.db.Database(tsUser.DBName)
	dbCollection := dbUser.Collection(tsUser.TabelName)
	Result, err := dbCollection.InsertOne(context.Background(),
		bson.D{
			{Key: "userID", Value: user.UserID},
			{Key: "userFname", Value: user.UserFname},
			{Key: "userLname", Value: user.UserLname},
			{Key: "userMobile", Value: user.UserMobile},
		})
	if err != nil {
		fmt.Println(" Not Insert Record")
	}
	if Result.InsertedID != nil {
		fmt.Printf("%v, type = %T\n", Result.InsertedID, Result.InsertedID)
		return true
	}
	return false
}

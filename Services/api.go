package Services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"example.com/ansiblego/MongoDb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	UserID int `bson:"userID"`
}

var DBName = "db_user"
var TabelName = "tbl_user"

func Connection() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Println(err)
	}
	//defer client.Disconnect(ctx)

	return client
}

func setupCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(*w).Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, PUT, PATCH, POST, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Requested-With, Authorization")
}

func Home(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w, r)
	fmt.Fprintf(w, "Home Call ...")
}
func Create(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w, r)
	var client = Connection()
	var UserService MongoDb.IUserService = MongoDb.NewUserServiceStruct(client, DBName, TabelName)
	UserService.Create()
	fmt.Fprintf(w, "Create Call ...")
}
func GetAll(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w, r)
	var client = Connection()
	var UserService MongoDb.IUserService = MongoDb.NewUserServiceStruct(client, DBName, TabelName)
	userEntitys := UserService.GetAllUser()
	fmt.Fprintf(w, "%v", userEntitys)
}
func Insert(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w, r)
	defer r.Body.Close()

	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		fmt.Println(err)
	}

	userEntity := new(MongoDb.IUser)
	userEntity.UserID = user.UserID
	userEntity.UserFname = "Mohammad"
	userEntity.UserLname = "Rahimi"
	userEntity.UserMobile = "09115755339"

	var client = Connection()
	var UserService MongoDb.IUserService = MongoDb.NewUserServiceStruct(client, DBName, TabelName)
	state := UserService.Insert(userEntity)
	fmt.Fprintf(w, "Insert Call  And UserID : "+strconv.Itoa(user.UserID), " state = ", state)
}

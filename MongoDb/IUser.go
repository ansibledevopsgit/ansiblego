package MongoDb

type IUser struct {
	UserID       int     `bson:"userID"`
	UserFname    string    `bson:"userFname"`
	UserLname    string    `bson:"userLname"`
	UserMobile   string    `bson:"userMobile"`
}

func New(userID int, userFname string, userLname string, userMobile string ) *IUser {
	return &IUser{
		UserID:       userID,
		UserFname:    userFname,
		UserLname:    userLname,
		UserMobile:   userMobile,
	}
}

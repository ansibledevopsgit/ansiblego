package Services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)
type User struct {
    UserID int  `json:"userID"`
     
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
	fmt.Fprintf(w, "Create Call ...")
}
func GetAll(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w, r)
	fmt.Fprintf(w, "GetAll Call ...")
}
func Insert(w http.ResponseWriter, r *http.Request) {
	setupCORS(&w, r)
	defer r.Body.Close()

	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user) 
	if err != nil{ 
		fmt.Println(err)
	}
	userid := strconv.Itoa(user.UserID)
	fmt.Fprintf(w, "Insert Call  And UserID : "+  userid)
}

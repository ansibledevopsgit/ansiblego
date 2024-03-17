package main

import (
	"log"
	"net/http"
	//"example.com/ansiblego/Server"
)

// func setupCORS(w *http.ResponseWriter, req *http.Request) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
// 	(*w).Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, PUT, PATCH, POST, DELETE")
// 	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Requested-With, Authorization")
// }

func main() {

	log.Println("Start ...")
	//Server.Run()

	http.HandleFunc("/", Home)

	//ServerHost := "localhost:9000"
	errListen := http.ListenAndServe(":9000", nil)
	log.Println("Run ...")
	if errListen != nil {
		log.Println(errListen)
	}

}

func Home(w http.ResponseWriter, r *http.Request) {
	//setupCORS(&w, r)
	log.Println("Home Call Ok ...")
}

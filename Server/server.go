package Server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"example.com/ansiblego/Router"
)

type Configserver struct {
	Port string `json:"port"`
	Host string `json:"host"`
}

func Run() {

	log.Println(" Config Router ...")
	Router.Route()

	file, err := os.Open("./config.json")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	config := new(Configserver)
	errDecode := json.NewDecoder(file).Decode(config)
	if errDecode != nil {
		log.Println(errDecode)
	}

	fmt.Print(" server run")
	log.Println("Start http server.")

	ServerHost := config.Host+":"+config.Port
	errListen := http.ListenAndServe(ServerHost,nil)
	if errListen != nil {
		fmt.Println(errListen)
	}
}

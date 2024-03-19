package Server

import (
	"log"
	"net/http"

	"example.com/ansiblego/Router"
)

type Configserver struct {
	Port string `json:"port"`
	Host string `json:"host"`
}

func Run() {

	log.Println(" Config Router ...")
	Router.Route()
   
	errListen := http.ListenAndServe(":9000", nil)
	log.Println("Run http server...")
	if errListen != nil {
		log.Println(errListen)
	}
	
}

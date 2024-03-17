package Router

import (
	"net/http"

	"example.com/ansiblego/Services"
)

func Route() {

	http.HandleFunc("/", Services.Home)
	http.HandleFunc("/create", Services.Create)
	http.HandleFunc("/insert", Services.Insert)
	http.HandleFunc("/getall", Services.GetAll)
}

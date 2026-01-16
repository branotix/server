package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type user struct {
	Name       string `json:"name"`
	Satus      bool   `json:"satus"`
	SecreteKey string `json:"secretekey"`
}

var pulock user = user{
	Name:       "Pulock Kumar",
	Satus:      true,
	SecreteKey: "eXchange@@11",
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	jsondata, _ := json.Marshal(pulock)
	fmt.Fprintf(w, "%s", string(jsondata))

}
func main() {

	fmt.Println(pulock)
	jsondata, _ := json.Marshal(pulock)
	fmt.Println("this is json data ", string(jsondata))

	http.HandleFunc("/user", userHandler)
	fmt.Println("server is listining localhost:8080/user")
	http.ListenAndServe(":8080", nil)
}

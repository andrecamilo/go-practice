package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/thedevsaddam/govalidator"
)

type user struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Web      string `json:"web"`
	Age      int    `json:"age"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	user := &user{Username: "fadf", Email: "ssds@asaa.com", Web: "asasa.com", Age: 12}
	//var user user
	rules := govalidator.MapData{
		"username": []string{"required", "between:3,5"},
		"email":    []string{"required", "min:4", "max:20", "email"},
		"web":      []string{"url"},
		"age":      []string{"numeric_between:18,56"},
	}

	opts := govalidator.Options{
		Request: r,
		Data:    &user,
		Rules:   rules,
	}

	v := govalidator.New(opts)
	e := v.ValidateJSON()
	fmt.Println(user) // your incoming JSON data in Go data struct
	err := map[string]interface{}{"validationError": e}
	w.Header().Set("Content-type", "applciation/json")
	json.NewEncoder(w).Encode(err)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening on port: 9000")
	http.ListenAndServe(":9000", nil)
}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	serv := NewServer()

	err := http.ListenAndServe(":8090", serv)
	if err != nil {
		return
	}
}

func (serv *Server) route() {
	serv.HandleFunc("/users", users)
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Server struct {
	*mux.Router
	users []User
}

func NewServer() *Server {
	serv := &Server{
		Router: mux.NewRouter(),
		users:  []User{},
	}
	serv.route()
	return serv
}

func users(w http.ResponseWriter, req *http.Request) {
	jsonFile, err := os.Open("users.json")
	if err != nil {
		fmt.Println(err)
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			fmt.Println("Error while closing")
		}
	}(jsonFile)

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var users []User

	err1 := json.Unmarshal(byteValue, &users)
	if err1 != nil {
		return
	}
	usersInfo := "[\n"
	for i := 0; i < len(users); i++ {
		usersInfo += "{name: " + users[i].Name
		usersInfo += "email: " + users[i].Email + "},\n"
	}
	usersInfo += "]"
	_, err2 := fmt.Fprintf(w, usersInfo)
	if err2 != nil {
		return
	}
}

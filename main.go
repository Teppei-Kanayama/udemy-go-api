package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/Teppei-Kanayama/udemy-go-api/handlers"
	"gopkg.in/mgo.v2/bson"
)

// User holds data for a singlr user
type User struct {
	ID   bson.ObjectId `json:"id" storm:"id"`
	Name string        `json:"name"`
	Role string        `json:"role"`
}

const (
	dbPath = "users.db"
)

// errors
var (
	ErrRecordInvalid = errors.New("invalid")
)

func main() {
	http.HandleFunc("/users/", handlers.UsersRooter)
	http.HandleFunc("/users", handlers.UsersRooter)
	http.HandleFunc("/", handlers.RootHandler)
	err := http.ListenAndServe("localhost:11111", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

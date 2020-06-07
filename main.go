package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/Teppei-Kanayama/udemy-go-api/handlers"
	"gopkg.in/mgo.v2/bson"
)

type jsonResponse map[string]interface{}

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

// func usersGetAll(w http.ResponseWriter, r *http.Request) {
// 	users, err := All()
// 	if err != nil {
// 		postError(w, http.StatusInternalServerError)
// 		return
// 	}
// 	postBodyResponse(w, http.StatusOK, jsonResponse{"users": users})
// }
//
// func postBodyResponse(w http.ResponseWriter, code int, content jsonResponse) {
// 	if content != nil {
// 		js, err := json.Marshal(content)
// 		if err != nil {
// 			postError(w, http.StatusInternalServerError)
// 			return
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(code)
// 		w.Write(js)
// 		return
// 	}
// }

func main() {
	// http.HandleFunc("/users/", UsersRooter)
	// http.HandleFunc("/users", UsersRooter)
	http.HandleFunc("/", handlers.RootHandler)
	err := http.ListenAndServe("localhost:11111", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/Teppei-Kanayama/udemy-go-api/handlers"
	"github.com/asdine/storm"
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

// All retrieves all users from the database
func All() ([]User, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	users := []User{}
	if err != nil {
		return nil, err
	}
	return users, nil
}

// One returns a single user record from the database
func One(id bson.ObjectId) (*User, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	u := new(User)
	err = db.One("ID", id, u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Delete removes a given record from the database
func Delete(id bson.ObjectId) error {
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	u := new(User)
	err = db.One("ID", id, u)
	if err != nil {
		return err
	}
	return db.DeleteStruct(u)
}

// Save updates or creates a given record in the database
func (u *User) Save() error {
	if err := u.validate(); err != nil {
		return err
	}
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Save(u)
}

// validate makes sure that the record contains valid data
func (u *User) validate() error {
	if u.Name == "" {
		return ErrRecordInvalid
	}
	return nil
}

func usersGetAll(w http.ResponseWriter, r *http.Request) {
	users, err := All()
	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"users": users})
}

func postBodyResponse(w http.ResponseWriter, code int, content jsonResponse) {
	if content != nil {
		js, err := json.Marshal(content)
		if err != nil {
			postError(w, http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(js)
		return
	}
}

func main() {
	http.HandleFunc("/users/", UsersRooter)
	http.HandleFunc("/users", UsersRooter)
	http.HandleFunc("/", handlers.RootHandler)
	err := http.ListenAndServe("localhost:11111", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package user

import (
	"errors"

	"github.com/asdine/storm"
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

// All retrieves all users from the database
func All() ([]User, error) { // []Userは、要素がUser型のslice（Pythonで言うlist）
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()  // 遅延関数呼び出し（Allが終了する時に実行される）
	users := []User{} // ここたぶんゆくゆくは変えていくんだろうね
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
func (u *User) Save() error { // *UserのメソッドとしてSave()を定義している
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

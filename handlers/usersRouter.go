package handlers

import (
	"net/http"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

// UsersRooter is something
func UsersRooter(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSuffix(r.URL.Path, "/")
	if path == "/users" {
		switch r.Method {
		case http.MethodGet:
			usersGetAll(w, r)
			return
		case http.MethodPost:
			return
		default:
			postError(w, http.StatusMethodNotAllowed)
			return
		}
	}

	path = strings.TrimSuffix(path, "/users/")
	if !bson.IsObjectIdHex(path) {
		postError(w, http.StatusNotFound)
		return
	}

	// id := bson.ObjectIdHex(path)

	switch r.Method {
	case http.MethodGet:
		return
	case http.MethodPut:
		return
	case http.MethodPatch:
		return
	case http.MethodDelete:
		return
	default:
		postError(w, http.StatusMethodNotAllowed)
	}
}

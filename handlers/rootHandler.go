package handlers

import "net/http"

// RootHandler is something
func RootHandler(arg1 http.ResponseWriter, arg2 *http.Request) {
	if arg2.URL.Path != "/" {
		arg1.WriteHeader(http.StatusNotFound)
		arg1.Write([]byte("Noot Found \n"))
		return
	}
	arg1.WriteHeader(http.StatusOK)
	arg1.Write([]byte("Okkkkkkkkkkkay!\n"))
}

package routes

import (
	"fmt"
	"net/http"
)

func ListTables(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "ListTables Route\n")
}
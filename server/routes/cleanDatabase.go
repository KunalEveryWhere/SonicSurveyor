package routes

import (
	"fmt"
	"net/http"
)

func CleanDatabase(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w, "CleanDatabase Route\n")
}
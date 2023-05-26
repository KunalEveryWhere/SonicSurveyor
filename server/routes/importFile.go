package routes

import (
	"fmt"
	"net/http"
)

func ImportFile(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "ImportFile Route\n")
}
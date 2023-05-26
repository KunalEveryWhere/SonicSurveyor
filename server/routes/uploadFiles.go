package routes

import (
	"fmt"
	"net/http"
)

func UploadFiles(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "UploadFiles Route\n")
}
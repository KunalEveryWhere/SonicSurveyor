package routes

import (
	"fmt"
	"net/http"
)

func ExportTable(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "ExportTable Route\n")
}
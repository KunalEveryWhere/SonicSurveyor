package routes

import (
	"fmt"
	"net/http"
)

func DropTable(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "DropTable Route\n")
}
package routes

import (
	"fmt"
	"net/http"

	//User-Defined Packages
	// "sonicsurveyor.com/main/commands"
)

func DropTable(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "DropTable Route\n")
}
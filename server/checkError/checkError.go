package checkError

import (
	"fmt"
	"log"
	"net/http"
)

// Errors in payload, et cetera. These errors are not to cause a state of panic
func ExternalIssues(message string, w http.ResponseWriter, err error) {
	if (err != nil) {
		// ⭐️ TODO: Handle errors according to type
		http.Error(w, "Internal Server Error: "+message+": ", http.StatusInternalServerError)
		fmt.Println("Error:", err)
	}
}

// All Other Errors. These cause the state of panic.
func InternalIssues(message string, err error) {
	if (err != nil) {
		fmt.Println("Error: ", message, err);
		log.Fatalln("Error: ", err)
	}
}
package routes

import (
	"fmt"
	"net/http"

	//User-Defined Packages
	"sonicsurveyor.com/main/commands"
)

func ListTables(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is GET
	if r.Method != http.MethodGet {
		http.Error(w, "🟥 Method Not Allowed", http.StatusMethodNotAllowed)
		fmt.Println("🟥 Method Not Allowed. ONLY GET Method Allowed.")
		return
	}

		//Calling Clean-Database
		ch := make(chan string)
		go commands.ListTables(ch);
		receivedMessage := <- ch;
		
		//Returning values to client.
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, receivedMessage);
}
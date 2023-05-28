package routes

import (
	"fmt"
	"net/http"

	//User-Defined Packages
	"sonicsurveyor.com/main/commands"
)

func CleanDatabase(w http.ResponseWriter, r *http.Request){

	// Check if the request method is DELETE
	if r.Method != http.MethodDelete {
		http.Error(w, "🟥 Method Not Allowed", http.StatusMethodNotAllowed)
		fmt.Println("🟥 Method Not Allowed. ONLY DELETE Method Allowed.")
		return
	}

	//Calling Clean-Database
	ch := make(chan string)
	go commands.CleanDatabase(ch);
	receivedMessage := <- ch;
	
	//Returning values to client.
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, receivedMessage);
}
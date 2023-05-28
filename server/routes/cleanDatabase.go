package routes

import (
	"fmt"
	"net/http"

	//User-Defined Packages
	"sonicsurveyor.com/main/commands"
)

func CleanDatabase(w http.ResponseWriter, req *http.Request){

	//Calling Clean-Database
	ch := make(chan string)
	go commands.ImportOSM(ch);
	receivedMessage := <- ch;
	fmt.Println(receivedMessage)
	
	
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, receivedMessage);
}
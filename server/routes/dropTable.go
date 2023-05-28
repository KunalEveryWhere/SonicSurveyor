package routes

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"

	//User-Defined Packages
	"sonicsurveyor.com/main/commands"
)

type DropTableRequestBody struct {
	TableName  string `json:"tableName"`
}

func DropTable(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is DELETE
	if r.Method != http.MethodDelete {
		http.Error(w, "🟥 Method Not Allowed", http.StatusMethodNotAllowed)
		fmt.Println("🟥 Method Not Allowed. ONLY DELETE Method Allowed.")
		return
	}

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "🟥 Error reading request body", http.StatusBadRequest)
		fmt.Println("🟥 Error reading request body", err)
		return
	}

	// Unmarshal the JSON data
	var data DropTableRequestBody;
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "🟥 Error un-marshaling JSON data", http.StatusBadRequest)
		fmt.Println("🟥 Error un-marshaling JSON data", err)
		return
	}

	//Calling Drop-Table Command
	ch := make(chan string)
	go commands.DropTable(w, data.TableName, ch);
	receivedMessage := <- ch;
	
	//Returning values to client.
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, receivedMessage);
}
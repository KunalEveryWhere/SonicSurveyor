package routes

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"encoding/json"

	//User-Defined Packages
	"sonicsurveyor.com/main/commands"
	"sonicsurveyor.com/main/settings"
)

var WorkingFilesPathDir = settings.WorkingFilesPath;

type ImportFileFromJSONRequestBody struct {
	FileContents string `json:"fileContents"`
	NewFileName string `json:"newFileName"`
	GeneratedTableName string `json:"generatedTableName"`
}

func ImportFileFromJSON(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "🟥 Method Not Allowed", http.StatusMethodNotAllowed)
		fmt.Println("🟥 Method Not Allowed. ONLY POST Method Allowed.")
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
	var data ImportFileFromJSONRequestBody;
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "🟥 Error un-marshaling JSON data", http.StatusBadRequest)
		fmt.Println("🟥 Error un-marshaling JSON data", err)
		return
	}

	fileName := data.NewFileName;

	// Create a new file
	newFile, err := os.Create(WorkingFilesPathDir+"/"+fileName)
	if err != nil {
		http.Error(w, "🟥 Error creating file on the server", http.StatusInternalServerError)
		fmt.Println("🟥 Error creating file on the server", err);
		return
	}
	defer newFile.Close()

	// Write the request body contents to the new file
	_, err = newFile.Write([]byte(data.FileContents))
	if err != nil {
		http.Error(w, "🟥 Error writing file", http.StatusInternalServerError)
		fmt.Println("🟥 Error writing the OSM file on the server", err);
		return
	}

	// Import File to the Database
	ch := make(chan string)
	go commands.ImportFile(w, (WorkingFilesPathDir +"/"+ fileName), data.GeneratedTableName, ch);
	receivedMessage := <- ch;
	
	//Returning values to client.
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, receivedMessage);

}
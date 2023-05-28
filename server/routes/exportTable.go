package routes

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"os"
	"io"

	//User-Defined Packages
	"sonicsurveyor.com/main/commands"
	"sonicsurveyor.com/main/settings"
)

var WorkingFilesFolderPath = settings.WorkingFilesPath;

type ExportTableRequestBody struct {
	TableName  string `json:"tableName"`
}

func ExportTable(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is GET
	if r.Method != http.MethodGet {
		http.Error(w, "🟥 Method Not Allowed", http.StatusMethodNotAllowed)
		fmt.Println("🟥 Method Not Allowed. ONLY GET Method Allowed.")
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
	var data ExportTableRequestBody;
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "🟥 Error un-marshaling JSON data", http.StatusBadRequest)
		fmt.Println("🟥 Error un-marshaling JSON data", err)
		return
	}

	//Calling Export-Table Command
	ch := make(chan string)
	go commands.ExportTable(w, WorkingFilesFolderPath+"/", data.TableName+".geojson", data.TableName, ch);
	receivedMessage := <- ch;

	// Open the file on the server
	filePath :=  WorkingFilesFolderPath+"/"+data.TableName+".geojson";
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "🟥 Error finding file. It may not exist.", http.StatusInternalServerError)
		fmt.Println("🟥 Error finding/opening the exported file", err)
		return
	}
	defer file.Close()

	// Get the file's information
	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "🟥 Error getting file information", http.StatusInternalServerError)
		fmt.Println( "🟥 Error getting exported file information", err)
		return
	}

	// Set the response headers
	w.Header().Set("Content-Disposition", "attachment; filename="+fileInfo.Name())
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))

	// Send the file's content as the response body
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "🟥 Error sending file", http.StatusInternalServerError)
		fmt.Println("🟥 Error sending the exported file", err)
		return
	}

	//Returning values to client.
	fmt.Fprintf(w, receivedMessage);
}
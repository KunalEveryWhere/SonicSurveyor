package routes

import (
	"fmt"
	"net/http"
	"os"
	"io"

	//User-Defined Packages
	"sonicsurveyor.com/main/commands"
	"sonicsurveyor.com/main/settings"
)

var WorkingFilesPathLocation= settings.WorkingFilesPath;
const maxAllowedFileSize = 25 * 1024 * 1024 // 25MB

type ImportOSMFileRequestBody struct {
	EPSG  string `json:"EPSG"`
}

func ImportOSMFile(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "🟥 Method Not Allowed", http.StatusMethodNotAllowed)
		fmt.Println("🟥 Method Not Allowed. ONLY POST Method Allowed.")
		return
	}

	var data ImportOSMFileRequestBody;

	// Read Text-based Form value the request body
	data.EPSG = r.FormValue("EPSG");

	// Parse the form data
	err := r.ParseMultipartForm(maxAllowedFileSize) // Set the maximum file size to 25MB
	if err != nil {
		http.Error(w, "🟥 Error parsing form data", http.StatusInternalServerError)
		fmt.Println("🟥 Error parsing form data. File Size over 25MB not allowed.", err)
		return
	}

	//Upload the OSM DataFile to Server
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "🟥 Error retrieving file from form data", http.StatusBadRequest);
		fmt.Println("🟥 Error retrieving file from form data", err);
		return
	}
	defer file.Close()

	// Create a new file on the server
	newFile, err := os.Create(WorkingFilesPathLocation +"/"+ handler.Filename)
	if err != nil {
		http.Error(w, "🟥 Error creating file on the server", http.StatusInternalServerError)
		fmt.Println("🟥 Error creating file on the server", err);
		return
	}
	defer newFile.Close()

	// Copy the file data to the new file
	_, err = io.Copy(newFile, file)
	if err != nil {
		http.Error(w, "🟥 Error saving file", http.StatusInternalServerError)
		fmt.Println("🟥 Error saving file", err)
		return
	}

	// Import File to the Database
	ch := make(chan string)
	go commands.ImportOSM(w, (WorkingFilesPathLocation +"/"+ handler.Filename), data.EPSG, ch);
	receivedMessage := <- ch;
	
	//Returning values to client.
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, receivedMessage);

}

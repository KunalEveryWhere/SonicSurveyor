package routes

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"path/filepath"
	"regexp"
	"strings"

	//User-Defined Packages
	"sonicsurveyor.com/main/commands"
	"sonicsurveyor.com/main/settings"
)

var WorkingFilesPath = settings.WorkingFilesPath;

func ImportFile(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "🟥 Method Not Allowed", http.StatusMethodNotAllowed)
		fmt.Println("🟥 Method Not Allowed. ONLY POST Method Allowed.")
		return
	}

	// Parse the form data
	err := r.ParseMultipartForm(10 << 20) // Set the maximum file size to 10MB
	if err != nil {
		http.Error(w, "🟥 Error parsing form data", http.StatusInternalServerError)
		fmt.Println("🟥 Error parsing form data. File Size over 10MB not allowed.", err)
		return
	}

	// Get the file from the form data
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "🟥 Error retrieving file from form data", http.StatusBadRequest);
		fmt.Println("🟥 Error retrieving file from form data", err);
		return
	}
	defer file.Close()

	// Create a new file on the server
	newFile, err := os.Create(WorkingFilesPath +"/"+ handler.Filename)
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

	generatedTableName := removeExtensions(handler.Filename);

	// Import File to the Database
	ch := make(chan string)
	go commands.ImportFile(w, (WorkingFilesPath +"/"+ handler.Filename), generatedTableName, ch);
	receivedMessage := <- ch;
	
	//Returning values to client.
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, receivedMessage);

}


func removeExtensions(filename string) string {
		// Remove file extensions
		name := filepath.Base(filename)
		extension := filepath.Ext(name)
		name = name[:len(name)-len(extension)]
	
		// Remove non-alphabetic, non-numeric, and non-underscore characters
		reg := regexp.MustCompile("[^a-zA-Z0-9_]+")
		name = reg.ReplaceAllString(name, "")
	
		// Replace spaces with underscores
		name = strings.ReplaceAll(name, " ", "_")
	
		return name
}
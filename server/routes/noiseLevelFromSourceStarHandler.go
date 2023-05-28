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

var WorkingDirFilesPath = settings.WorkingFilesPath;
const maxFileSize = 25 * 1024 * 1024 // 25MB

type UploadFilesRequestBody struct {
	EPSG  string `json:"EPSG"`
	Temperature  string `json:"temperature"`
	Humidity  string `json:"humidity"`
}

func NoiseLevelFromSourceStarHandler(w http.ResponseWriter, r *http.Request) {
		// Check if the request method is POST
		if r.Method != http.MethodPost {
			http.Error(w, "🟥 Method Not Allowed", http.StatusMethodNotAllowed)
			fmt.Println("🟥 Method Not Allowed. ONLY POST Method Allowed.")
			return
		}

		var data UploadFilesRequestBody;

		// Read Text-based Form values the request body
		data.EPSG = r.FormValue("EPSG");
		data.Humidity = r.FormValue("humidity");
		data.Temperature = r.FormValue("temperature");
	
		// Parse the form data
		err := r.ParseMultipartForm(maxFileSize) // Set the maximum file size to 25MB
		if err != nil {
			http.Error(w, "🟥 Error parsing form data", http.StatusInternalServerError)
			fmt.Println("🟥 Error parsing form data. File Size over 25MB not allowed.", err)
			return
		}
	
		//Upload the OSM DataFile to Server
		fileName, err := UploadToServer("OSMDataFile", w, r);
		
		//Step 1: Import OSM Data
		// Import OSMData to the Database
		ch := make(chan string)
		go commands.ImportOSM(w, (WorkingDirFilesPath +"/"+ fileName), data.EPSG, ch);
		receivedMessage := <- ch;

		//Upload the POINT_SOURCE DataFile to Server
		fileName, err = UploadToServer("PointSourceDataFile", w, r);
	
		//Step 2: Import POINT_SOURCE Data
		// Import Point_SOURCE File to the Database
		ch = make(chan string)
		go commands.ImportFile(w, (WorkingDirFilesPath +"/"+ fileName), "POINT_SOURCE", ch);
		receivedMessage = <- ch;

		//Step 3: Calculate Triangles and Receivers
		ch = make(chan string)
		go commands.CalculateReceivers(w, "BUILDINGS", "POINT_SOURCE", ch);
		receivedMessage = <- ch;

		//Step 4: Calculate Noise Level from Source
		ch = make(chan string)
		go commands.CalculateNoiseLevelFromSource(w, "BUILDINGS", "POINT_SOURCE", "RECEIVERS", data.Humidity, data.Temperature, ch);
		receivedMessage = <- ch;

		//Step 5: Generate Noise Iso-Level
		ch = make(chan string)
		go commands.GenerateNoiseLevelIsoSurface(w, "LDAY_GEOM", ch);
		receivedMessage = <- ch;
		
		//Returning values to client.
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, receivedMessage);
}

func UploadToServer(parameterName string, w http.ResponseWriter, r *http.Request) (string, error){
	// Get the file from the form data
	file, handler, err := r.FormFile(parameterName)
	if err != nil {
		http.Error(w, "🟥 Error retrieving "+parameterName+" from form data", http.StatusBadRequest);
		fmt.Println("🟥 Error retrieving "+parameterName+" from form data", err);
		return "", err
	}
	defer file.Close()

	// Create a new file on the server
	newFile, err := os.Create(WorkingDirFilesPath +"/"+ handler.Filename)
	if err != nil {
		http.Error(w, "🟥 Error creating "+parameterName+" on the server", http.StatusInternalServerError)
		fmt.Println("🟥 Error creating "+parameterName+" on the server", err);
		return "", err
	}
	defer newFile.Close()

	// Copy the file data to the new file
	_, err = io.Copy(newFile, file)
	if err != nil {
		http.Error(w, "🟥 Error saving file", http.StatusInternalServerError)
		fmt.Println("🟥 Error saving file", err)
		return "", err
	}

	//Return values
	return handler.Filename, nil
}
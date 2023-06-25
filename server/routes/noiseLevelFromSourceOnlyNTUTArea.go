package routes

import (
	"fmt"
	"net/http"

	//User-Defined Packages
	"sonicsurveyor.com/main/commands"
	"sonicsurveyor.com/main/settings"
)

var WorkingDirFilePath = settings.WorkingFilesPath;
const maximumFileSize = 25 * 1024 * 1024 // 25MB

func NoiseLevelFromSourceOnlyNTUTArea(w http.ResponseWriter, r *http.Request) {
		// Check if the request method is POST
		if r.Method != http.MethodPost {
			http.Error(w, "🟥 Method Not Allowed", http.StatusMethodNotAllowed)
			fmt.Println("🟥 Method Not Allowed. ONLY POST Method Allowed.")
			return
		}

		var data UploadFilesRequestBody;

		// Read Text-based Form values the request body
		data.Humidity = r.FormValue("humidity");
		data.Temperature = r.FormValue("temperature");
	
		// Parse the form data
		err := r.ParseMultipartForm(maximumFileSize) // Set the maximum file size to 25MB
		if err != nil {
			http.Error(w, "🟥 Error parsing form data", http.StatusInternalServerError)
			fmt.Println("🟥 Error parsing form data. File Size over 25MB not allowed.", err)
			return
		}

		
		//Step 1: Import OSM Data
		// Import OSMData to the Database
		ch := make(chan string)
		go commands.ImportOSM(w, ("util/NTUT_OSM_20May2023.osm"), "3826", ch);
		receivedMessage := <- ch;

		//Step 2: Import Point-Source
		// (It should be uploaded from a different route already)

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
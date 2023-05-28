package routes

import (
	"fmt"
	"net/http"

	//User-Defined Packages
	//"sonicsurveyor.com/main/checkError"
)

func Description(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "This is the SonicSurveyor API. Below are its Endpoints and Descriptions\n\n/description \t This relays information about the various endpoints and their services in this API.\n/echoHeaders \t This echos back to the client all headers it has received. This can be used to check the state of the API (running or not-running).\n/noiseLevelFromSourceStarHandler \t This allows the client to upload data / files for processing. It takes in OSM Data, Point_Source geojson, EPSG SRID Code, humidity, temperature => Runs all required BG processes, and created the with the final 'CONTOURING_NOISE_MAP' table. NOTE: Because this single route combined and utilizes many functions, it might take some time to run. Be patient.\n/importFile \t This allows the client to import additional files onto the database.\n/importOSMFile \t This allows the client to upload an .osm file and add its contents to the database\n/importFileFromJSON \t This allows the client to upload .geojson data as in string, and add it to the database.\n/listTables \t\t This allows the client to list all the tables and its columns currently in the database.\n/exportTable \t This allows the client to export any table from the database (if exists)\n/dropTable \t\t This allows the client to delete any given table from the database.\n/cleanDatabase \t This allows the client to drop all tables and clean the database.")
}
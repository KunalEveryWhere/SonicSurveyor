package routes

import (
	"fmt"
	"net/http"

	//User-Defined Packages
	//"sonicsurveyor.com/main/checkError"
)

func Description(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "This is the SonicSurveyor API. Below are its Endpoints and Descriptions\n\n/description \t This relays information about the various endpoints and their services in this API.\n/echoHeaders \t This echos back to the client all headers it has received. This can be used to check the state of the API (running or not-running).\n/uploadFiles \t This allows the client to upload data / files for processing.\n/importFile \t This allows the client to import additional files onto the database.\n/listTables \t\t This allows the client to list all the tables and its columns currently in the database.\n/exportTable \t This allows the client to export any table from the database (if exists)\n/dropTable \t\t This allows the client to delete any given table from the database.\n/cleanDatabase \t This allows the client to drop all tables and clean the database.")
}
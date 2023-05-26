package routes

import (
	"fmt"
	"net/http"
)

func Description(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "This is the SonicSurveyor API. Below are its Endpoints and Descriptions\n\n/description \t This relays information about the various endpoints and their services in this API.\n/echoHeaders \t This echos back to the client all headers it has received. This can be used to check the state of the API (running or not-running).\n")
}
package createServer

import (
	"fmt"
	"net/http"
	"net"

	//cors
	"github.com/rs/cors"

	//User-defined Packages
	"sonicsurveyor.com/main/settings"
	"sonicsurveyor.com/main/checkError"
	"sonicsurveyor.com/main/routes"
)

var PORT string = settings.PORT;

func getIPv4Address() (string, error) {
	// Retrieve the list of network interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	// Iterate over the network interfaces
	for _, iface := range interfaces {
		// Filter out loopback and non-up interfaces
		if iface.Flags&net.FlagLoopback == 0 && iface.Flags&net.FlagUp != 0 {
			// Retrieve the addresses for the interface
			addrs, err := iface.Addrs()
			if err != nil {
				return "", err
			}

			// Iterate over the addresses
			for _, addr := range addrs {
				// Check if the address is an IPv4 address
				if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
					return ipNet.IP.String(), nil
				}
			}
		}
	}

	return "", fmt.Errorf("unable to determine IPv4 address")
}

func MainHandler() {
	mux := http.NewServeMux()

	//Create a new CORS middleware with custom options
	cors := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{
            http.MethodPost,
            http.MethodGet,
			http.MethodDelete,
        },
        AllowedHeaders:   []string{"*"},
        AllowCredentials: false,
    })


	// Register the handler function with the default ServeMux (HTTP request multiplexer)
	mux.HandleFunc("/", routes.Description)
	mux.HandleFunc("/cleanDatabase", (routes.CleanDatabase))
	mux.HandleFunc("/description", (routes.Description))
	mux.HandleFunc("/dropTable", (routes.DropTable))
	mux.HandleFunc("/echoHeaders", (routes.EchoHeaders))
	mux.HandleFunc("/exportTable", (routes.ExportTable))
	mux.HandleFunc("/importFile", (routes.ImportFile))
	mux.HandleFunc("/importFileFromJSON", (routes.ImportFileFromJSON))
	mux.HandleFunc("/importOSMFile", (routes.ImportOSMFile))
	mux.HandleFunc("/listTables", (routes.ListTables))
	mux.HandleFunc("/noiseLevelFromSourceOnlyNTUTArea", (routes.NoiseLevelFromSourceOnlyNTUTArea))
	mux.HandleFunc("/noiseLevelFromSourceStarHandler", (routes.NoiseLevelFromSourceStarHandler))

	ip, err := getIPv4Address();
	checkError.InternalIssues("Error whilst retrieving IPv4 Address", err)

	//Write the details of where the server started
	fmt.Println("Starting the Server on: "+ip+":"+PORT)
	fmt.Println("\tPress CTRL + C to Terminate. \n\n")

	// Wrap the default ServeMux with the CORS middleware
	handler := cors.Handler(mux)

	// Start the server on port specified in ../settings.PORT
    err = http.ListenAndServe(":"+PORT, handler)
	// err = server.ListenAndServeTLS("", "")
	checkError.InternalIssues("Error Listening & Serving the Server", err);
}



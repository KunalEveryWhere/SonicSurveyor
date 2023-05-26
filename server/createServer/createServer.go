package createServer

import (
	"fmt"
	"net/http"
	"net"

	//User-Defined Packages
	"sonicsurveyor.com/main/settings"
	"sonicsurveyor.com/main/checkError"
	"sonicsurveyor.com/main/routes"
)

var PORT string = settings.PORT;

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

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
	// Register the handler function with the default ServeMux (HTTP request multiplexer)
    http.HandleFunc("/", routes.Description)
	http.HandleFunc("/description", routes.Description)
    http.HandleFunc("/echoHeaders", routes.EchoHeaders)

	ip, err := getIPv4Address();
	checkError.InternalIssues("Error whilst retrieving IPv4 Address", err)

	//Write the details of where the server started
	fmt.Println("Starting the Server on: "+ip+":"+PORT)
	fmt.Println("\tPress CTRL + C to Terminate. \n\n")

	// Start the server on port specified in ../settings.PORT
    err = http.ListenAndServe(":"+PORT, nil)
	checkError.InternalIssues("Error Listening & Serving the Server", err);
}
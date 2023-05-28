package commands

import (
	"fmt"
	"os/exec"
	"net/http"

	//User-defined Packages
	"sonicsurveyor.com/main/settings"
	"sonicsurveyor.com/main/checkError"
)


func ImportOSM(w http.ResponseWriter, filepath string, SRID string, ch chan<- string) {
	//Get the path to Util
	pathToUtil := settings.PathToUtil;

	// Command to run to import OSM
	cmd := exec.Command(pathToUtil+"/bin/wps_scripts", "-w", "./", "-s", pathToUtil+"/noisemodelling/wps/Import_and_Export/Import_OSM.groovy", "-pathFile", filepath, "-targetSRID", SRID)

	// Run the command and get the output
	output, err := cmd.CombinedOutput();
	checkError.ExternalIssues("Error while Importing OSM File to Database\n"+(string(output)), w, err);

	// Print & send the output to channel
	fmt.Println("\n\n\nImport OSM File | Command output: [ ✅ successful ]\n", string(output))
	ch <- string(output);
}
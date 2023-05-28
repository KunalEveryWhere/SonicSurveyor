package commands

import (
	"fmt"
	"os/exec"
	"net/http"

	//User-defined Packages
	"sonicsurveyor.com/main/settings"
	"sonicsurveyor.com/main/checkError"
)


func GenerateNoiseLevelIsoSurface(w http.ResponseWriter, resultsTable string, ch chan<- string) {
	//Get the path to Util
	pathToUtil := settings.PathToUtil;

	// Command to run to generate noise level isosurfaces
	cmd := exec.Command(
		pathToUtil+"/bin/wps_scripts",
		"-w", "./",
		"-s", 
		pathToUtil+"/noisemodelling/wps/Acoustic_Tools/Create_Isosurface.groovy",
		"-resultTable", resultsTable)

	// Run the command and get the output
	output, err := cmd.CombinedOutput();
	checkError.ExternalIssues("Error while Generating Noise Level Iso-surfaces: \n"+(string(output)), w, err);

	// Print & send the output to channel
	fmt.Println("\n\n\nGenerate Noise Level Iso-surfaces | Command output: [ ✅ successful ]\n", string(output))
	ch <- string(output);
}
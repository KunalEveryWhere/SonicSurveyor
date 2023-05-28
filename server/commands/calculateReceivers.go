package commands

import (
	"fmt"
	"os/exec"
	"net/http"

	//User-defined Packages
	"sonicsurveyor.com/main/settings"
	"sonicsurveyor.com/main/checkError"
)


func CalculateReceivers(w http.ResponseWriter, buildingTableName string, sourceTableName string, ch chan<- string) {
	//Get the path to Util
	pathToUtil := settings.PathToUtil;

	// Command to run to calculate the receivers
	cmd := exec.Command(pathToUtil+"/bin/wps_scripts", "-w", "./", "-s", pathToUtil+"/noisemodelling/wps/Receivers/Delaunay_Grid.groovy", "-tableBuilding", buildingTableName, "-sourcesTableName", sourceTableName)

	// Run the command and get the output
	output, err := cmd.CombinedOutput();
	checkError.ExternalIssues("Error while Calculating the Receivers Table: \n"+(string(output)), w, err);

	// Print & send the output to channel
	fmt.Println("\n\n\nRECEIVERS Table | Command output: [ ✅ successful ]\n", string(output))
	ch <- string(output);
}
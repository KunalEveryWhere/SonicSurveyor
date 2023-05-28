package commands

import (
	"fmt"
	"os/exec"
	"net/http"

	//User-defined Packages
	"sonicsurveyor.com/main/settings"
	"sonicsurveyor.com/main/checkError"
)


func ImportFile(w http.ResponseWriter, filepath string, generatedTableName string, ch chan<- string) {
	//Get the path to Util
	pathToUtil := settings.PathToUtil;

	// Command to run to import file
	cmd := exec.Command(
		pathToUtil+"/bin/wps_scripts", "-w", 
		"./", "-s", 
		pathToUtil+"/noisemodelling/wps/Import_and_Export/Import_File.groovy",
		"-pathFile", filepath,
		"-tableName", generatedTableName);

	// Run the command and get the output
	output, err := cmd.CombinedOutput();
	checkError.ExternalIssues("Error while Importing File to Database\n"+(string(output)), w, err);

	// Print & send the output to channel
	msg := "Import File | Command output: [ ✅ successful ]\n";
	fmt.Println("\n\n\n"+msg, string(output))
	ch <- msg+string(output);
}
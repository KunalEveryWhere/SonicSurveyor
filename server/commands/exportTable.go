package commands

import (
	"fmt"
	"os/exec"
	"net/http"

	//User-defined Packages
	"sonicsurveyor.com/main/settings"
	"sonicsurveyor.com/main/checkError"
)


func ExportTable(w http.ResponseWriter, exportFilePath string, exportFileName string, exportTableName string, ch chan<- string) {
	//Get the path to Util
	pathToUtil := settings.PathToUtil;

	// Command to run to export the table
	cmd := exec.Command(
		pathToUtil+"/bin/wps_scripts", "-w", 
		"./", "-s", pathToUtil+"/noisemodelling/wps/Import_and_Export/Export_Table.groovy",
		"-exportPath", exportFilePath+exportFileName,
		"-tableToExport", exportTableName)

	// Run the command and get the output
	output, err := cmd.CombinedOutput();
	checkError.ExternalIssues("Error Exporting Table\n"+(string(output)), w, err);

	// Print & send the output to channel
	fmt.Println("\n\n\nExport Table | Command output: [ ✅ successful ]\n", string(output))
	ch <- string(output);
}
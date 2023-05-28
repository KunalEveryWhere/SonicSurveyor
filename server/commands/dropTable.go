package commands

import (
	"fmt"
	"os/exec"
	"net/http"

	//User-defined Packages
	"sonicsurveyor.com/main/settings"
	"sonicsurveyor.com/main/checkError"
)


func DropTable(w http.ResponseWriter, tableName string, ch chan<- string) {
	//Get the path to Util
	pathToUtil := settings.PathToUtil;

	// Command to run to drop a table
	cmd := exec.Command(pathToUtil+"/bin/wps_scripts", "-w", "./", "-s", pathToUtil+"/noisemodelling/wps/Database_Manager/Drop_a_Table.groovy", "-tableToDrop", tableName)

	// Run the command and get the output
	output, err := cmd.CombinedOutput()
	checkError.ExternalIssues("Error while Dropping table: "+tableName+" from Database\n"+(string(output)), w, err);

	// Print & send the output to channel
	fmt.Println("\n\n\nDrop Table | Command output: [ ✅ successful ]\n", string(output))
	ch <- string(output);
}
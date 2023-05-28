package commands

import (
	"fmt"
	"os/exec"

	//User-defined Packages
	"sonicsurveyor.com/main/settings"
	"sonicsurveyor.com/main/checkError"
)


func CleanDatabase(ch chan<- string) {
	//Get the path to Util
	pathToUtil := settings.PathToUtil;

	// Command to run to Clean Database
	cmd := exec.Command(pathToUtil+"/bin/wps_scripts", "-w", "./", "-s", pathToUtil+"/noisemodelling/wps/Database_Manager/Clean_Database.groovy", "-areYouSure", "true")

	// Run the command and get the output
	output, err := cmd.CombinedOutput()
	checkError.InternalIssues("Error in Cleaning the Database\n"+(string(output)), err);

	// Print & send the output to channel
	fmt.Println("\n\n\nClean Database | Command output: [ ✅ successful ]\n", string(output))
	ch <- string(output);
}
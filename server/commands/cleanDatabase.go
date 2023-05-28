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
	msg := "Clean Database | Command output: [ ✅ successful ]\n";
	fmt.Println("\n\n\n"+msg, string(output))
	ch <- (msg+string(output));
}
package commands

import (
	"fmt"
	"os/exec"

	//User-defined Packages
	"sonicsurveyor.com/main/settings"
	"sonicsurveyor.com/main/checkError"
)

func SetupPermissions(ch chan<- string){
	//Get the path to Util
	pathToUtil := settings.PathToUtil;
	
	// Command to run to Setup executable permissions on the wps_scripts
	cmd := exec.Command("sudo", "chmod", "755", pathToUtil+"/bin/wps_scripts")

	// Run the command and get the output
	output, err := cmd.CombinedOutput()
	checkError.InternalIssues("Error in setting up permissions\n"+(string(output)), err);

	// Print & send the output to channel
	fmt.Println("\n\n\nSetup Permissions | Command output: [ ✅ successful ]\n", string(output))
	ch <- string(output);
}
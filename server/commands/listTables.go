package commands

import (
	"fmt"
	"os/exec"

	//User-defined Packages
	"sonicsurveyor.com/main/settings"
	"sonicsurveyor.com/main/checkError"
)


func ListTables(ch chan<- string) {
	//Get the path to Util
	pathToUtil := settings.PathToUtil;
	
	// Command to run to all tables in the Database
	cmd := exec.Command(pathToUtil+"/bin/wps_scripts", "-w", "./", "-s", pathToUtil+"/noisemodelling/wps/Database_Manager/Display_Database.groovy", "-showColumns", "false")


	// Run the command and get the output
	output, err := cmd.CombinedOutput();
	checkError.InternalIssues("Error in listing all the tables in the database\n"+(string(output)), err);

	// Print & send the output to channel
	msg := "List Table | Command output: [ ✅ successful ] \n";
	fmt.Println("\n\n\n"+msg, string(output))
	ch <- msg+string(output);
}
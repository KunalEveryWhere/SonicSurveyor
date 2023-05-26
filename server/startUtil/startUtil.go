package startUtil

import (
	// "os"
	// "os/exec"
	// "log"
	"fmt"
	// "path/filepath"

	//User-defined packages
	"sonicsurveyor.com/main/settings"
	// "sonicsurveyor.com/main/checkError"
)

var pathToUtil string;

func init(){
	pathToUtil = settings.PathToUtil;
}

func StartUtil() {
	// // Create the sudo command with the script as an argument
	// cmd := exec.Command("sh", pathToUtil)

	// // Set the stdout and stderr to os.Stdout and os.Stderr
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	// // Run the command
	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	wd := settings.PathToUtil
	fmt.Println("Working Directory:", wd)
}
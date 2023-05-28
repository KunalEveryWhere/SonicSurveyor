package settings

import (
	// "path/filepath"
	"fmt"
	"log"
	"os"
)

// "os"
// "path/filepath"

var PORT string
var PathToUtil string;
var WorkingFilesPath string;

func init() {
	PORT = "26001"
	
	//For development
	PathToUtil = "/Users/kunaleverywhere/Documents/External Projects/SonicSurveyor/SonicSurveyor/server/util/NM_4.0.0_WO_GUI"
	WorkingFilesPath = "/Users/kunaleverywhere/Documents/External Projects/SonicSurveyor/SonicSurveyor/server/WorkingFiles"

	//For build
	//exePath, _ := os.Executable()
	//PathToUtil = filepath.Dir(exePath)+"/util/NM_4.0.0_WO_GUI"
	// WorkingFilesPath = filepath.Dir(exePath) + "/WorkingFiles"

	//Call required functions
	CreateFolders(WorkingFilesPath);
}

func CreateFolders(dirPath string){
	// Check if the directory already exists
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// Create the directory
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			fmt.Println("Error creating directory: ", err);
			log.Fatal("Error creating directory: ", err);
		}
	}
}
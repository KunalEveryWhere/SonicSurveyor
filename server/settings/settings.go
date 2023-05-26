package settings

import (
	// "os"
	// "path/filepath"
)

var PORT string
var PathToUtil string;

func init() {
	PORT = "26001"
	
	//For development
	PathToUtil = "./util/NM_4.0.0_WO_GUI"

	//For build
	//exePath, _ := os.Executable()
	//PathToUtil = filepath.Dir(exePath)
}

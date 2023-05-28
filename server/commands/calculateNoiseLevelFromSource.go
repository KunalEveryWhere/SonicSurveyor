package commands

import (
	"fmt"
	"os/exec"
	"net/http"

	//User-defined Packages
	"sonicsurveyor.com/main/settings"
	"sonicsurveyor.com/main/checkError"
)


func CalculateNoiseLevelFromSource(w http.ResponseWriter, buildingTableName string, sourceTableName string, receiverTableName, humidity string, temperature string, ch chan<- string) {
	//Get the path to Util
	pathToUtil := settings.PathToUtil;

	if(humidity == "" || temperature == ""){
		cmd := exec.Command(
			pathToUtil+"/bin/wps_scripts",
			"-w", "./",
			"-s", 
			pathToUtil+"/noisemodelling/wps/NoiseModelling/Noise_level_from_source.groovy",
			"-tableBuilding", buildingTableName,
			"-tableSources", sourceTableName,
			"-tableReceivers", receiverTableName,
			"-confSkipLevening", "true",
			"-confSkipLnight", "true",
			"-confSkipLden", "true")
	
		// Run the command and get the output
		output, err := cmd.CombinedOutput();
		checkError.ExternalIssues("Error while Calculating the Noise-Level from Source: \n"+(string(output)), w, err);
	
		// Print & send the output to channel
		fmt.Println("\n\n\nNoise Level From Source | Command output: [ ✅ successful ]\n", string(output))
		ch <- string(output);
	} else {
		// Command to run to calculate the noise-level from source
	cmd := exec.Command(
		pathToUtil+"/bin/wps_scripts",
		"-w", "./",
		"-s", 
		pathToUtil+"/noisemodelling/wps/NoiseModelling/Noise_level_from_source.groovy",
		"-tableBuilding", buildingTableName,
		"-tableSources", sourceTableName,
		"-tableReceivers", receiverTableName,
		"-confSkipLevening", "true",
		"-confSkipLnight", "true",
		"-confSkipLden", "true",
		"-confHumidity", humidity,
		"-confTemperature", temperature)

	// Run the command and get the output
	output, err := cmd.CombinedOutput();
	checkError.ExternalIssues("Error while Calculating the Noise-Level from Source: \n"+(string(output)), w, err);

	// Print & send the output to channel
	fmt.Println("\n\n\nNoise Level From Source | Command output: [ ✅ successful ]\n", string(output))
	ch <- string(output);
	}
}
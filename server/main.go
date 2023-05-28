package main

import (
	"fmt"
	"net/http"

	//User-defined packages
	"sonicsurveyor.com/main/createServer"
	"sonicsurveyor.com/main/commands"
)

func init() {

}

var w http.ResponseWriter;
var OSM_File string = "/Users/kunaleverywhere/Desktop/Pilot Study - Noise Modelling/1. Original OSM Data/NTUT OSM.osm";
var POINT_SOURCE string = "/Users/kunaleverywhere/Desktop/Pilot Study - Noise Modelling/4. Point Sources & SHP Data/Point_Source_Hypo.geojson";
var Generated_Table_Name = "POINT_SOURCE"
var Export_FilePath string = "/Users/kunaleverywhere/Desktop/";
var Export_FileName string = "HypoSP-NCM.geojson"

func main() {
	displayEntryMessage();

	//0. Initilize and give permissions
	ch := make(chan string)
	go commands.SetupPermissions(ch);
	_ = <- ch;

	//1. Start the Server
	createServer.MainHandler();
}

func displayEntryMessage(){
	fmt.Println("\n\n")
	fmt.Println(ASCIIArt())
	fmt.Println("\n\n")
}

func ASCIIArt() string {
	return ("░██████╗░█████╗░███╗░░██╗██╗░█████╗░░██████╗██╗░░░██╗██████╗░██╗░░░██╗███████╗██╗░░░██╗░█████╗░██████╗░\n██╔════╝██╔══██╗████╗░██║██║██╔══██╗██╔════╝██║░░░██║██╔══██╗██║░░░██║██╔════╝╚██╗░██╔╝██╔══██╗██╔══██╗\n╚█████╗░██║░░██║██╔██╗██║██║██║░░╚═╝╚█████╗░██║░░░██║██████╔╝╚██╗░██╔╝█████╗░░░╚████╔╝░██║░░██║██████╔╝\n░╚═══██╗██║░░██║██║╚████║██║██║░░██╗░╚═══██╗██║░░░██║██╔══██╗░╚████╔╝░██╔══╝░░░░╚██╔╝░░██║░░██║██╔══██╗\n██████╔╝╚█████╔╝██║░╚███║██║╚█████╔╝██████╔╝╚██████╔╝██║░░██║░░╚██╔╝░░███████╗░░░██║░░░╚█████╔╝██║░░██║\n╚═════╝░░╚════╝░╚═╝░░╚══╝╚═╝░╚════╝░╚═════╝░░╚═════╝░╚═╝░░╚═╝░░░╚═╝░░░╚══════╝░░░╚═╝░░░░╚════╝░╚═╝░░╚═╝\n\n\n░██████╗███████╗██████╗░██╗░░░██╗███████╗██████╗░  ██╗░░░██╗░░███╗░░\n██╔════╝██╔════╝██╔══██╗██║░░░██║██╔════╝██╔══██╗  ██║░░░██║░████║░░\n╚█████╗░█████╗░░██████╔╝╚██╗░██╔╝█████╗░░██████╔╝  ╚██╗░██╔╝██╔██║░░\n░╚═══██╗██╔══╝░░██╔══██╗░╚████╔╝░██╔══╝░░██╔══██╗  ░╚████╔╝░╚═╝██║░░\n██████╔╝███████╗██║░░██║░░╚██╔╝░░███████╗██║░░██║  ░░╚██╔╝░░███████╗\n╚═════╝░╚══════╝╚═╝░░╚═╝░░░╚═╝░░░╚══════╝╚═╝░░╚═╝  ░░░╚═╝░░░╚══════╝\n")
}

func stepsToBeFollowed(){
	// 1. Initilize and give permissions
	ch := make(chan string)
	go commands.SetupPermissions(ch);
	_ = <- ch;

	//2. Import a OSM File onto the DB
	ch = make(chan string)
	go commands.ImportOSM(w, OSM_File, "3826", ch);
	receivedMessage := <- ch;
	fmt.Println(receivedMessage)

	//3. Import a Point_Source File onto the DB
	ch = make(chan string)
	go commands.ImportFile(w, POINT_SOURCE, Generated_Table_Name, ch);
	receivedMessage = <- ch;
	fmt.Println(receivedMessage)

	//4. Calculate Triangles and Receivers
	ch = make(chan string)
	go commands.CalculateReceivers(w, "BUILDINGS", Generated_Table_Name, ch);
	receivedMessage = <- ch;
	fmt.Println(receivedMessage)

	//5. Calculate Noise Level from Source
	ch = make(chan string)
	go commands.CalculateNoiseLevelFromSource(w, "BUILDINGS", Generated_Table_Name, "RECEIVERS", "", "", ch);
	receivedMessage = <- ch;
	fmt.Println(receivedMessage)

	//6. Generate Noise Iso-Level
	ch = make(chan string)
	go commands.GenerateNoiseLevelIsoSurface(w, "LDAY_GEOM", ch);
	receivedMessage = <- ch;
	fmt.Println(receivedMessage)

	//7. List all the tables in the DB
	ch = make(chan string)
	go commands.ListTables(ch);
	receivedMessage = <- ch;
	fmt.Println(receivedMessage)

	//8. Export the final table
	ch = make(chan string)
	go commands.ExportTable(w, Export_FilePath, Export_FileName, "CONTOURING_NOISE_MAP", ch);
	receivedMessage = <- ch;
	fmt.Println(receivedMessage)
}
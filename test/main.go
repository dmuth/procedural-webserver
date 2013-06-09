/**
* This is a sample app that starts up the procedural webserver.
*/

package main

//import "fmt"

import log "github.com/dmuth/google-go-log4go"
import server "github.com/dmuth/procedural-webserver"


func main() {

	//
	// Parse our arguments and report them
	//
	config := ParseArgs()
	log.Infof(
		"Config: NumLinksMin: %d, NumLinksMax: %d, " +
		"NumImagesMin: %d, NumImagesMax: %d, Seed: %s", 
		config.NumLinksMin, config.NumLinksMax,
		config.NumImagesMin, config.NumImagesMax, 
		config.Seed)

	//
	// Now fire up the server and run it, forever.
	//
	server_object := server.NewServer(8080, 
		config.NumLinksMin, config.NumLinksMax, 
		config.NumImagesMin, config.NumImagesMax, 
		config.Seed)
	server_object.Start()

} // End of main()



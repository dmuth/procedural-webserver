
package main

//import "fmt"

import log "github.com/dmuth/google-go-log4go"


func main() {

	//
	// Parse our arguments and report them
	//
	config := Parse()
	log.Infof(
		"Config: NumLinksMin: %d, NumLinksMax: %d, " +
		"NumImagesMin: %d, NumImagesMax: %d, Seed: %s", 
		config.NumLinksMin, config.NumLinksMax,
		config.NumImagesMin, config.NumImagesMax, 
		config.Seed)

	//
	// Now fire up the server and run it, forever.
	//
	server_object := NewServer(config, 8080)
	server_object.Start()

}




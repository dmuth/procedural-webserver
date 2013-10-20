/**
* This script is used to just to run the server from the command line for 
* testing purposes.
*/
package main

import "flag"
import "fmt"
import "os"

import server "github.com/dmuth/procedural-webserver"


func main() {

	port := flag.Int("port", 8080, "Port to listen on")
	num_links_min := flag.Uint("num-links-min", 1, "Minimum number of links per page")
	num_links_max := flag.Uint("num-links-max", 10, "Maximum number of links per page")
	num_images_min := flag.Uint("num-images-min", 1, "Minimum number of images per page")
	num_images_max := flag.Uint("num-images-max", 10, "Maximum number of images per page")
	seed := flag.String("seed", "test_seed", "Seed to use for random values")
	h := flag.Bool("h", false, "To get this help")
	help := flag.Bool("help", false, "To get this help")

	flag.Parse()

	if (*h || *help) {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Printf("About to start server on localpost port %d... (^C to exit!)", *port)
	server_obj := server.NewServer(*port, *num_links_min, *num_links_max,
		*num_images_min, *num_images_max, 
		*seed)
	server_obj.Start()

} // End of main()


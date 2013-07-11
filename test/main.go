/**
* This script is used to just to run the server from the command line for 
* testing purposes.
*/
package main

import "fmt"
import server ".."

func main() {

	port := 8080
	fmt.Printf("About to start server on localpost port %d... (^C to exit!)", 8080)

	server_obj := server.NewServer(port, 10, 20, 10, 20, "test_seed")
	server_obj.Start()


} // End of main()


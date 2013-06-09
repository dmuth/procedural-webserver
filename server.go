
package server

import "fmt"
import "time"
import "net"
import "net/http"
import "strconv"

import log "github.com/dmuth/google-go-log4go"


//
// Configuration for the server
//
type Config struct {
	NumLinksMin uint
	NumLinksMax uint
	NumImagesMin uint
	NumImagesMax uint
	Seed string
}


type Server_struct struct {
	//
	// Our html structure
	//
	html Html_struct
	//
	// What port are we listening on?
	//
	port int
	//
	// Our listener, so we can stop the server
	//
	listener net.Listener
	//
	// Embedded configuration information
	//
	config Config
}


/**
* Instantiate a structure for our webserver.
* @param {int} port What port are we running on?
* @param {uint} num_links_min Min number of links per page
* @param {uint} num_links_max Max number of links per page
* @param {uint} num_images_min Min number of images per page
* @param {uint} num_images_max Max number of images per page
* @param {string} seed The seed which all random strings will be based off of
*
* @return {Server_struct} Our server structure
*
*/
func NewServer(port int, num_links_min uint, num_links_max uint, 
	num_images_min uint, num_images_max uint, 
	seed string) (retval Server_struct) {

	var listener net.Listener
	config := Config{ num_links_min, num_links_max,
		num_images_min, num_images_max, seed }
	html_struct := NewHtml(config)

	retval = Server_struct{ html_struct, port, listener, config }

	return(retval)

} // End of New()


/**
* Actually start our webserver.
* This will block forever, so we should run this as a goroutine 
* if we want to continue going in our program.
* 
*/
func (s *Server_struct) Start() {

	var err error
	s.listener, err = net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if (err != nil) {
		panic(err)
		return
	}

	http.Handle("/", s)

	http.Serve(s.listener, nil)

} // End of Start()


/**
* Stop our currently running server.
*/
func (s *Server_struct) Stop() {

	//
	// If we try to close something that is nil, we'll get a panic.
	// This is a race condition which can be triggered if we close 
	// the server immediately after starting it.
	//
	if (s.listener == nil) {
		//
		// I'm not sure if not closing a listener which hasn't yet opened 
		// can cause issues down the road, hence the Warning.
		//
		log.Warn("Stop(): Listener is currently nil, not closing.")

	} else {
		s.listener.Close()

	}

} // End of Stop()


/**
* Our responder handler.  This is used when serving up a page.
* Yes, this implements the http.Handler interface, as descirbed 
* at http://golang.org/pkg/net/http/#Handler
* I never thought I'd be implementing an interface this soon. Scary easy.
*
*/
func (s *Server_struct) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	uri := req.RequestURI
	log.Infof("Request Start: %s (%s %s)", req.URL.Path, uri, req.RemoteAddr)
	start_time := time.Now()

	code, _ := strconv.Atoi(req.FormValue("code"))
	log.Debugf("Code passed in: %d", code)

	delay := req.FormValue("delay")
	log.Debugf("Delay passed in: %s", delay)
	if (delay != "") {
		duration, _ := time.ParseDuration(delay)
		log.Debugf("Pausing for %s...", duration )
		time.Sleep(duration)
	}

	//
	// Set an error code if one was passed in.
	// We'll still send the content, since that's legitimate.
	//
	if (code != 0) {
		http.Error(res, "", code)
	}

	//
	// Our URI is our seed so we'll get the same content on repeated page 
	// loads. (aka procedural page generation!)
	//
	output := s.html.Html(uri)
	fmt.Fprintf(res, output)

	elapsed := time.Now().Sub(start_time)
	log.Infof("Request Complete in %.6f sec: %s (%s %s)",
		float64(elapsed.Nanoseconds()) / float64(1000000000),
		req.URL.Path, req.RequestURI, req.RemoteAddr)

} // End of ServeHTTP()



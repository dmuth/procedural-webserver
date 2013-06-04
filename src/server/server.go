/**
* This package is responsible for actually running the webserver.
*/

package server

import "fmt"
import "time"
import "net"
import "net/http"
import "strconv"

import log "github.com/dmuth/google-go-log4go"

import "../args"
import "../html"


type Server_struct struct {
	//
	// Our html structure
	//
	html html.Html_struct
	//
	// What port are we listening on?
	//
	port int
	//
	// Our listener, so we can stop the server
	//
	listener net.Listener
}


/**
* Instantiate a structure for our webserver.
* @param {Config} html_config Our configuration for HTML pages
* @param {int} port What port are we running on?
*/
func New(html_config args.Config, port int) (retval Server_struct) {

	var listener net.Listener
	html_struct := html.New(html_config)
	retval = Server_struct{html_struct, port, listener}

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
	s.listener.Close()
}


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



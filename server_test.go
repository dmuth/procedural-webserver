package server

import "fmt"
import "net/http"
import "io/ioutil"
import "regexp"
import "testing"

func TestServer(t *testing.T) {

	port := 8080
	server_obj := NewServer(port, 10, 20, 10, 20, "test_seed")

	go server_obj.Start()

	result := httpGet(fmt.Sprintf("http://localhost:%d", port))
	//fmt.Printf("%s\n", result)
	pattern := "mwhxoejfka"
	match, _ := regexp.MatchString(pattern, result)
	if !match {
		t.Errorf("could not find pattern '%s' in result '%s", pattern, result)
	}

	result = httpGet(fmt.Sprintf("http://localhost:%d/12345", port))
	//fmt.Printf("%s\n", result)
	pattern = "dspnjxhzkm"
	match, _ = regexp.MatchString(pattern, result)
	if !match {
		t.Errorf("could not find pattern '%s' in result '%s", pattern, result)
	}

	result = httpGet(fmt.Sprintf("http://localhost:%d/?foo=bar", port))
	//fmt.Printf("%s\n", result)
	pattern = "fsnckooszy"
	match, _ = regexp.MatchString(pattern, result)
	if !match {
		t.Errorf("could not find pattern '%s' in result '%s", pattern, result)
	}

	// TODO: Test a 404 response, new function should return a code as well

	server_obj.Stop()

}

/**
* Helper function to make GET requests and return the value.
 */
func httpGet(url string) (retval string) {

	resp, err := http.Get(url)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		return
	}

	return (fmt.Sprintf("%s", body))

} // End of httpGet()

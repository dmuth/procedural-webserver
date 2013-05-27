
package html

import "fmt"
import "time"

import "../args"
import rand "../random_sorta"
import log "github.com/dmuth/google-go-log4go"


/**
* Create a fake HTML page.
* 
* @param {args.Config} Our configuration structure
* @return {string} Our fake page with links and images
*/
func Html(config args.Config) (retval string) {

	//
	// Determine how many links go on this page
	//
	num_links := uint(1)
	if (config.NumLinksMin != config.NumLinksMax) {
		diff := config.NumLinksMax - config.NumLinksMin
		random := rand.New( uint(time.Now().Nanosecond()), uint(diff) )
		num_links = random.Intn()
		num_links += config.NumLinksMin
	}
	log.Debug(fmt.Sprintf("Number of links on page: %d", num_links))

	//
	// How many images on the page?
	//
	num_images := uint(1)
	if (config.NumImagesMin != config.NumImagesMax) {
		diff := config.NumImagesMax - config.NumImagesMin
		random := rand.New(uint(time.Now().Nanosecond()), uint(diff))
		num_images = random.Intn()
		num_images += uint(config.NumImagesMin)
	}
	log.Debug(fmt.Sprintf("Number of images on page: %d", num_images))

	//
	// Create the actual links and images
	//
	random := rand.New(10, uint(time.Now().Nanosecond()))

	for i:=uint(0); i<num_links; i++ {
		str := random.StringLowerN(10)
		retval += fmt.Sprintf("<a href=\"/%s\" >%s</a>\n", str, str);
	}

	for i:=uint(0); i<num_images; i++ {
		str := random.StringLowerN(10)
		retval += fmt.Sprintf("<img src=\"/%s.png\" alt=\"%s\" />\n", str, str);
	}

	retval = "<html>\n" + 
		"<head><title></title></head>\n" +
		"<body>\n" +
		retval + 
		"</body\n" +
		"</html>\n"

	return(retval)

} // End of Html()




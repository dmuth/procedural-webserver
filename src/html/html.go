
package html

import "fmt"

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
	// Create the actual links and images
	//
	retval += getLinks(config)
	retval += getImages(config)

	retval = "<html>\n" + 
		"<head><title></title></head>\n" +
		"<body>\n" +
		retval + 
		"</body\n" +
		"</html>\n"

	return(retval)

} // End of Html()


/**
* Generate our links. 
*/
func getLinks(config args.Config) (retval string) {

	//
	// Determine how many links go on this page
	//
	num_links := uint(1)
	if (config.NumLinksMin != config.NumLinksMax) {
		diff := config.NumLinksMax - config.NumLinksMin
		random := rand.New( config.Seed, uint(diff) )
		num_links = random.Intn()
		num_links += config.NumLinksMin
	}
	log.Debug(fmt.Sprintf("Number of links on page: %d", num_links))

	random := rand.New( config.Seed, 100 )
	for i:=uint(0); i<num_links; i++ {
		str := random.StringLowerN(10)
		retval += fmt.Sprintf("<a href=\"/%s\" >%s</a>\n", str, str);
	}

	return(retval);

} // End of getLinks()


/**
* Generate our images.
*/
func getImages(config args.Config) (retval string) {

	//
	// How many images on the page?
	//
	num_images := uint(1)
	if (config.NumImagesMin != config.NumImagesMax) {
		diff := config.NumImagesMax - config.NumImagesMin
		random := rand.New( config.Seed, uint(diff) )
		num_images = random.Intn()
		num_images += uint(config.NumImagesMin)
	}
	log.Debug(fmt.Sprintf("Number of images on page: %d", num_images))

	random := rand.New( config.Seed, 100 )
	for i:=uint(0); i<num_images; i++ {
		str := random.StringLowerN(10)
		retval += fmt.Sprintf(
			"<img src=\"/%s.png\" alt=\"%s\" title=\"%s\" />\n", 
			str, str, str);
	}

	return(retval)

} // End of getImages()




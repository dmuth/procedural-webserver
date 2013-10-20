
package server

import "fmt"

import log "github.com/dmuth/google-go-log4go"


type Html_struct struct {
	config Config
	random_num_links Random_struct
	num_links_diff uint
	random_num_images Random_struct
	num_images_diff uint
	random_chars Random_struct
}


/**
* Create a new variable/object based on our Html_struct structure.
*/
func NewHtml(config Config) (retval Html_struct) {

	num_links_diff := config.NumLinksMax - config.NumLinksMin
	random_num_links := NewRand()

	num_images_diff := config.NumImagesMax - config.NumImagesMin
	random_num_images := NewRand()

	random_chars := NewRand()

	retval = Html_struct{config,
		random_num_links, num_links_diff,
		random_num_images, num_images_diff,
		random_chars	}

	return(retval)

} // End of New()


/**
* Create a fake HTML page.
* 
* @param {string} seed Our seed for the page
*
* @return {string} Our fake page with links and images
*/
func (h *Html_struct) Html(seed string) (retval string) {

	//
	// Glue our base seed onto our URI that was passed in
	//
	seed += h.config.Seed

	//
	// Create the actual links and images
	//
	retval += "<a href=\"/\">[Home]</a><br/>\n"
	retval += "<br/>\n"
	retval += h.getLinks(seed, h.config)
	retval += "<br/>\n"
	retval += h.getImages(seed, h.config)
	retval += "<br/>\n"

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
*
* @param {string} seed Our seed
*
*/
func (h *Html_struct) getLinks(seed string, config Config) (retval string) {

	//
	// Determine how many links go on this page
	//
	num_links := uint(1)
	if (config.NumLinksMin != config.NumLinksMax) {
		num_links = h.random_num_links.Intn(seed, h.num_links_diff)
		num_links += config.NumLinksMin
		// Change the seed for the next iteration of the loop
		seed += "13"
	}
	log.Debug(fmt.Sprintf("Number of links on page: %d", num_links))

	for i:=uint(0); i<num_links; i++ {
		str := h.random_chars.StringLowerN(seed, 10)
		retval += fmt.Sprintf("<a href=\"/%s\" >%s</a><br/>\n", str, str);
		// Change the seed for the next iteration of the loop
		seed += "13"
	}

	return(retval);

} // End of getLinks()


/**
* @param {string} seed Our seed
*
* Generate our images.
*/
func (h *Html_struct) getImages(seed string, config Config) (retval string) {

	//
	// How many images on the page?
	//
	num_images := uint(1)
	if (config.NumImagesMin != config.NumImagesMax) {
		num_images = h.random_num_images.Intn(seed, h.num_images_diff)
		num_images += uint(config.NumImagesMin)
	}
	log.Debug(fmt.Sprintf("Number of images on page: %d", num_images))

	for i:=uint(0); i<num_images; i++ {
		str := h.random_chars.StringLowerN(seed, 10)
		retval += fmt.Sprintf(
			"<img src=\"/%s.png\" alt=\"%s\" title=\"%s\" /><br/>\n", 
			str, str, str);
		//
		// Increase the seed by a prime for the next iteration.
		// I used to do this by 1, but then we had a "staircase" effect
		// for subsequent numbers.
		//
		// And this is why I should not write a random number generator. Ever.
		//
		seed += "13"
	}

	return(retval)

} // End of getImages()




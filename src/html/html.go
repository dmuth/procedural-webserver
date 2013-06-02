
package html

import "fmt"

import "../args"
import rand "../random_sorta"
import log "github.com/dmuth/google-go-log4go"


type Html_struct struct {
	config args.Config
	random_num_links rand.Random_struct
	num_links_diff uint
	random_num_images rand.Random_struct
	num_images_diff uint
	random_chars rand.Random_struct
}


/**
* Create a new variable/object based on our Html_struct structure.
*/
func New(config args.Config) (retval Html_struct) {

	num_links_diff := config.NumLinksMax - config.NumLinksMin
	seed := config.Seed + 1
	random_num_links := rand.New(seed)

	num_images_diff := config.NumImagesMax - config.NumImagesMin
	seed = config.Seed + 2
	random_num_images := rand.New(seed)

	random_chars := rand.New(config.Seed)

	retval = Html_struct{config,
		random_num_links, num_links_diff,
		random_num_images, num_images_diff,
		random_chars	}

	return(retval)

} // End of New()


/**
* Create a fake HTML page.
* 
* @return {string} Our fake page with links and images
*/
func (h *Html_struct) Html() (retval string) {

	//
	// Create the actual links and images
	//
	retval += h.getLinks(h.config)
	retval += h.getImages(h.config)

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
func (h *Html_struct) getLinks(config args.Config) (retval string) {

	//
	// Determine how many links go on this page
	//
	num_links := uint(1)
	if (config.NumLinksMin != config.NumLinksMax) {
		num_links = h.random_num_links.Intn(h.num_links_diff)
		num_links += config.NumLinksMin
	}
	log.Debug(fmt.Sprintf("Number of links on page: %d", num_links))

	for i:=uint(0); i<num_links; i++ {
		str := h.random_chars.StringLowerN(10)
		retval += fmt.Sprintf("<a href=\"/%s\" >%s</a>\n", str, str);
	}

	return(retval);

} // End of getLinks()


/**
* Generate our images.
*/
func (h *Html_struct) getImages(config args.Config) (retval string) {

	//
	// How many images on the page?
	//
	num_images := uint(1)
	if (config.NumImagesMin != config.NumImagesMax) {
		num_images = h.random_num_images.Intn(h.num_images_diff)
		num_images += uint(config.NumImagesMin)
	}
	log.Debug(fmt.Sprintf("Number of images on page: %d", num_images))

	for i:=uint(0); i<num_images; i++ {
		str := h.random_chars.StringLowerN(10)
		retval += fmt.Sprintf(
			"<img src=\"/%s.png\" alt=\"%s\" title=\"%s\" />\n", 
			str, str, str);
	}

	return(retval)

} // End of getImages()




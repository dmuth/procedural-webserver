
package args

import "flag"
import "os"


//
// Configuration for what was passed in on the command line.
//
type config struct {
	NumLinksMin int
	NumLinksMax int
	NumImagesMin int
	NumImagesMax int
	//MaxLevels int,
	//NumKeyValuePairs
}


/**
* Parse our command line arguments.
* @return {config} Our configuration info
*/
func Parse() (retval config) {

	retval = config{0, 0, 0, 0}
	flag.IntVar(&retval.NumLinksMin, "num-links-min", 1, 
		"Minimum number of links per page")
	flag.IntVar(&retval.NumLinksMax, "num-links-max", 1,
		"Maximum number of links per page")
	flag.IntVar(&retval.NumImagesMin, "num-images-min", 1,
		"Minimum number of image links per page")
	flag.IntVar(&retval.NumImagesMax, "num-images-max", 1,
		"Maximum number of image links per page")
	h := flag.Bool("h", false, "To get this help")
	help := flag.Bool("help", false, "To get this help")
	flag.Parse()

	if (*h || *help) {
		flag.PrintDefaults()
		os.Exit(1)
	}

	return(retval)

} // End of Parse()



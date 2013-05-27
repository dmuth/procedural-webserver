
package args

import "flag"
import "fmt"
import "os"


//
// Configuration for what was passed in on the command line.
//
type Config struct {
	NumLinksMin uint
	NumLinksMax uint
	NumImagesMin uint
	NumImagesMax uint
	//MaxLevels int,
	//NumKeyValuePairs
}


/**
* Parse our command line arguments.
* @return {config} Our configuration info
*/
func Parse() (retval Config) {

	retval = Config{0, 0, 0, 0}
	flag.UintVar(&retval.NumLinksMin, "num-links-min", 1, 
		"Minimum number of links per page")
	flag.UintVar(&retval.NumLinksMax, "num-links-max", 0,
		"Maximum number of links per page")
	flag.UintVar(&retval.NumImagesMin, "num-images-min", 1,
		"Minimum number of image links per page")
	flag.UintVar(&retval.NumImagesMax, "num-images-max", 0,
		"Maximum number of image links per page")
	h := flag.Bool("h", false, "To get this help")
	help := flag.Bool("help", false, "To get this help")
	flag.Parse()

	if (retval.NumLinksMax == 0) {
		retval.NumLinksMax = 1
	}

	if (retval.NumImagesMax == 0) {
		retval.NumImagesMax = 1
	}

	if (retval.NumLinksMax < retval.NumLinksMin) {
		panic(fmt.Sprintf(
			"Max num links (%d) is less than min num links (%d)!",
			retval.NumLinksMax, retval.NumLinksMin))
	}

	if (retval.NumImagesMax < retval.NumImagesMin) {
		panic(fmt.Sprintf(
			"Max num images (%d) is less than min num images (%d)!",
			retval.NumImagesMax, retval.NumImagesMin))
	}

	if (*h || *help) {
		flag.PrintDefaults()
		os.Exit(1)
	}

	return(retval)

} // End of Parse()



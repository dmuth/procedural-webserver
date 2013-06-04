
package args

import "flag"
import "fmt"
import "os"

import log "github.com/dmuth/google-go-log4go"


//
// Configuration for what was passed in on the command line.
//
type Config struct {
	NumLinksMin uint
	NumLinksMax uint
	NumImagesMin uint
	NumImagesMax uint
	Seed string
	//MaxLevels int,
	//NumKeyValuePairs
}


/**
* Parse our command line arguments.
* @return {config} Our configuration info
*/
func Parse() (retval Config) {

	retval = Config{0, 0, 0, 0, ""}

	flag.StringVar(&retval.Seed, "seed", "generic_seed",
		"Random seed to start with. This provides deterministic " +
		"behavior between runs, which is great for testing purposes. " +
		"If not specified, the default will be used. ")
	flag.UintVar(&retval.NumLinksMin, "num-links-min", 1,
		"Minimum number of links per page")
	flag.UintVar(&retval.NumLinksMax, "num-links-max", 2,
		"Maximum number of links per page")
	flag.UintVar(&retval.NumImagesMin, "num-images-min", 1,
		"Minimum number of image links per page")
	flag.UintVar(&retval.NumImagesMax, "num-images-max", 2,
		"Maximum number of image links per page")

	h := flag.Bool("h", false, "To get this help")
	help := flag.Bool("help", false, "To get this help")

	debug_level := flag.String("debug-level", "info", "Set the debug level")

	flag.Parse()

	log.SetLevelString(*debug_level)
	log.Error("Debug level: " + *debug_level)

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



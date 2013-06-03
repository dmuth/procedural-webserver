

package random_sorta

import "bytes"
import "crypto/md5"
import "encoding/binary"
import "fmt"

//import log "github.com/dmuth/google-go-log4go"


type Random_struct struct {
}


/**
* Create and return a new structure.
* 
*	We're specifying this at creation time so that the bitmask 
*	only needs to be created once.
*/
func New() (retval Random_struct) {

	retval = Random_struct{}

	return(retval)

} // End of New()


/**
* Core function that actually grabs the next "random" integer.
*
* @param {string} seed A seed to supply.
*
* @return {int} An integer
*/
func (r *Random_struct) int(seed string) (retval uint) {

	//
	// Create a hash based on our current seed
	//
	hash := md5.New()
	//hash.Write([]byte(fmt.Sprintf("%d", seed)))
// TEST
	hash.Write([]byte(seed))
	md5_value := hash.Sum(nil)
// TEST
	//fmt.Println("MD5:", fmt.Sprintf("%x", md5_value))

	//
	// Grab 8 bytes 
	//
	buf := bytes.NewBuffer(md5_value)

	//
	// I'm not yet clear on why I can't use 32 bits, but anyway 
	// here's where I'm going to grab the least significant 
	// 32 bits and store them in retval
	//
	var retval64 uint64
	binary.Read(buf, binary.LittleEndian, &retval64)
	retval = uint(retval64)
	//fmt.Printf("%x, %d\n", md5_value, retval) // Debugging

	return(retval)

} // End of int()


/**
* Return a random number between 1 and n
*
* @param {string} seed A seed to supply.
*
* @return {integer} retval The random value
*/
func (r *Random_struct) Intn(seed string, max uint) (retval uint) {

	if (max == 0) {
		panic("Max can't be == 0!")
	}

	retval = r.int(seed)
	bitmask := getBitmask(max)
	retval = retval & bitmask

	//
	// If the value is too big (e.g. 32 when the max is 17), call ourself
	// again and hope we get lucky.
	// (And I hope this never causes a stack overflow...)
	//
	if (retval >= max) {
// TEST
		hash := md5.New()
		hash.Write([]byte(seed))
//fmt.Printf("%x\n", hash.Sum(nil))
//fmt.Printf("%x\n", hash.Sum(nil))
		//md5_value := hash.Sum(nil)
		//seed = fmt.Sprintf("%x", md5_value)
		seed = fmt.Sprintf("%x", hash.Sum(nil))

// TEST
//fmt.Println("TEST", fmt.Sprintf("%x", seed), retval, max)
		retval = r.Intn(seed, max)

	}

	return(retval)

} // End of intn()


/**
* Create a bitmask from our max value.  This is for extracting that 
* value from an MD5 hash.
*
* @param {int} max Our maximum random value
* @return {int} A value which is 2*n-1.
*
* @TODO: In the future, I may want to address performance issues.  
*	I can think of a few ways:
*	- Cache results (could get out of control on memory usage, though)
*	- Require a max number to be specificed 
*/
func getBitmask(max uint) (retval uint) {

	retval = 1
	for i:=1; i<64; i++ {
		retval *= 2
		if (retval >= max) {
			break
		}
	}

	retval--

	return(retval)

} // End of getBitmask()


/**
* Return a lowercase string of num characters.
*
* @param {string} seed A seed to supply.
* @param {int} num How many characters do we want?
*
* @return {string} The random string
*/
func (r *Random_struct) StringLowerN(seed string, num_chars uint) (retval string) {

	//
	// Loop through our integers until we get something in the 
	// first 26 numbers.
	//
	for {

		num := r.Intn(seed, 26)
		num += 97
		//fmt.Printf("%d: %c\n", num, num) // Debugging
		retval = retval + fmt.Sprintf("%c", num)

		seed = r.mutateSeed(seed)

		if (uint(len(retval)) >= num_chars) {
			break
		}

	}

	return(retval)

} // End of StringLowerN()


/**
* Return a mixed case string of num characters.
*
* @param {string} seed A seed to supply.
* @param {int} num How many characters do we want?
*
* @return {string} The random string
*/
func (r *Random_struct) StringN(seed string, num_chars uint) (retval string) {

	//
	// Loop through our integers until we get something in the 
	// first 52 numbers.
	//
	for {

		num := r.Intn(seed, 52) 
		if (num <= 25) {
			num += 65
		} else {
			num -= 26
			num += 97
		}
		//fmt.Printf("%d: %c\n", num, num)

		retval = retval + fmt.Sprintf("%c", num)

		seed = r.mutateSeed(seed)

		if (uint(len(retval)) >= num_chars) {
			break
		}

	}

	return(retval)

} // End of StringLowerN()


/**
* Alter our seed in some way so that we get randomness on future passes.
*
* @param {string} seed The old seed
*
* @return {string} The mutated seed
*/
func (r *Random_struct) mutateSeed(seed string) (retval string) {

	//
	// Okay, this is silly.
	// I may change this up in the future and maybe have a counter in the structure,
	// but I did want seed mutation in a centralized place, so here it is!
	//
	retval = seed + "1"

	return(retval)

} // End of mutateSeed()




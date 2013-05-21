

package random_sorta

import "bytes"
import "crypto/md5"
import "encoding/binary"
import "fmt"

import log "github.com/dmuth/google-go-log4go"


type random_struct struct {
	//
	// The results of the previous hash.
	//
	seed string
	//
	// We will generate this number or lower.
	//
	max uint64
	//
	// Our bitmask for pulling values from MD5 responses
	//
	bitmask uint64
}


/**
* Create and return a new structure.
* 
* @param {string} seed An optional seed to supply.
* @param {integer} max All random numbers are less than this.
*	We're specifying this at creation time so that the bitmask 
*	only needs to be created once.
*/
func New(seed uint64, max uint64) (retval random_struct) {

	if (max == 0) {
		panic("Max can't be == 0!")
	}

	bitmask := getBitmask(max)

	retval = random_struct{fmt.Sprintf("%d", seed), max, bitmask}

	return(retval)

} // End of New()


/**
* Core function that actually grabs the next "random" integer.
* @return {int} An integer
*/
func (r *random_struct) int() (retval uint64) {

	//
	// Create a hash based on our current seed
	//
	hash := md5.New()
	hash.Write([]byte(r.seed))
	md5_value := hash.Sum(nil)
	//fmt.Println("MD5:", fmt.Sprintf("%x", md5_value))

	//
	// Make the current hash our new seed
	//
	r.seed = string(md5_value)

	//
	// Grab 8 bytes and put them into a uint64
	//
	// @todo In the future, I should store the remaining 8 bytes for 
	// the next call
	//
	buf := bytes.NewBuffer(md5_value)
	binary.Read(buf, binary.LittleEndian, &retval)
	//fmt.Println(md5_value, retval)

	return(retval)

} // End of int()


/**
* Return a random number between 1 and n
* @return {integer} retval The random value
*/
func (r *random_struct) Intn() (retval uint64) {

	retval = r.int()
	retval = retval & r.bitmask

	//
	// If the value is too big (e.g. 32 when the max is 17), call ourself
	// again and hope we get lucky.
	// (And I hope this never causes a stack overflow...)
	//
	if (retval >= r.max) {
		retval = r.Intn()
	}

	return(retval)

} // End of intn()


/**
* Read a request off of a channel, generate a random value, and write 
* it back out.
*
* @param {integer} How many random numbers do we want back?
* @param {chan int} out The channel to write results out to
*/
func (r *random_struct) IntnChannel(in chan uint64, out chan []uint64) {

	log.Info("Spawned IntNChannel()")

	for {
		var retval []uint64
		num_random := <-in

		for i:=uint64(0); i<num_random; i++ {
			num := r.Intn()
			retval = append(retval, num)
		}

		out <- retval

	}

} // End of IntnChannel()


/**
* Create a bitmask from our max value.  This is for extracting that 
* value from an MD5 hash.
*
* @param {uint64} max Our maximum random value
* @return {uint64} A value which is 2*n-1.
*
* @TODO: In the future, I may want to address performance issues.  
*	I can think of a few ways:
*	- Cache results (could get out of control on memory usage, though)
*	- Require a max number to be specificed 
*/
func getBitmask(max uint64) (retval uint64) {

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
* @param {int} num How many characters do we want?
* @return {string} The random string
*/
func (r *random_struct) StringLowerN(num_chars int) (retval string) {

	//
	// Loop through our integers until we get something in the 
	// first 26 numbers.
	//
	for {

		num := r.Intn() & 31
		if (num <= 25) {
			retval = retval + fmt.Sprintf("%c", num + 97)
		}

		if (len(retval) >= num_chars) {
			break
		}

	}

	return(retval)

} // End of StringLowerN()


/**
* Return a mixed case string of num characters.
* @param {int} num How many characters do we want?
* @return {string} The random string
*/
func (r *random_struct) StringN(num_chars int) (retval string) {

	//
	// Loop through our integers until we get something in the 
	// first 52 numbers.
	//
	for {

		num := r.Intn() & 63
		if (num <= 51) {

			if (num <= 25) {
				num += 65
			} else {
				num -= 26
				num += 97
			}
			//fmt.Printf("%d: %c\n", num, num)

			retval = retval + fmt.Sprintf("%c", num)

		}

		if (len(retval) >= num_chars) {
			break
		}

	}

	return(retval)

} // End of StringLowerN()







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
	max uint
	//
	// Our bitmask for pulling values from MD5 responses
	//
	bitmask uint
}


/**
* Create and return a new structure.
* 
* @param {string} seed An optional seed to supply.
* @param {integer} max All random numbers are less than this.
*	We're specifying this at creation time so that the bitmask 
*	only needs to be created once.
*/
func New(seed uint, max uint) (retval random_struct) {

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
func (r *random_struct) int() (retval uint) {

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
	// Grab 8 bytes 
	//
	// @todo In the future, maybe I should store the remaining 8 bytes for 
	// the next call
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
	//fmt.Println(md5_value, retval)

	return(retval)

} // End of int()


/**
* Return a random number between 1 and n
* @return {integer} retval The random value
*/
func (r *random_struct) Intn() (retval uint) {

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
* @param {chan uint} How many random numbers do we want back?
* @param {chan uint} out The channel to write results out to
*/
func (r *random_struct) IntnChannel(in chan uint, out chan []uint) {

	log.Info("Spawned IntNChannel()")

	for {
		var retval []uint
		num_random := <-in

		for i:=uint(0) ; i<num_random; i++ {
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
* @param {int} num How many characters do we want?
* @return {string} The random string
*/
func (r *random_struct) StringLowerN(num_chars uint) (retval string) {

	//
	// Loop through our integers until we get something in the 
	// first 26 numbers.
	//
	for {

		num := r.Intn() & 31
		if (num <= 25) {
			retval = retval + fmt.Sprintf("%c", num + 97)
		}

		if (uint(len(retval)) >= num_chars) {
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
func (r *random_struct) StringN(num_chars uint) (retval string) {

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

		if (uint(len(retval)) >= num_chars) {
			break
		}

	}

	return(retval)

} // End of StringLowerN()





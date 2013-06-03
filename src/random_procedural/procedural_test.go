
package random_procedural

//import "fmt"
import "testing"


func TestStringN(t *testing.T) {

	seed := "/foobar"
	random := New()

	output := random.StringN(seed, 1)
	expected := "b"
	if (output != expected) {
			t.Errorf("Got: %s, expected: %s", output, expected)
	}

	output = random.StringN(seed, 40)
	expected = "bhCXnZlRMvuYHpYtobOMNixvEpiiKLRSPGkgBOVU"
	if (output != expected) {
			t.Errorf("Got: %s, expected: %s", output, expected)
	}

	output = random.StringN(seed, 10)
	expected = "bhCXnZlRMv"
	if (output != expected) {
			t.Errorf("Got: %s, expected: %s", output, expected)
	}

} // End of TestStringLowerN()


func TestStringLowerN(t *testing.T) {

	seed := "/foobar"
	random := New()
	output := random.StringLowerN(seed, 1)
	expected := "u"
	if (output != expected) {
			t.Errorf("Got: %s, expected: %s", output, expected)
	}

	output = random.StringLowerN(seed, 10)
	expected = "ubcxhzfrmp"
	if (output != expected) {
			t.Errorf("Got: %s, expected: %s", output, expected)
	}

	seed = "/"
	output = random.StringLowerN(seed, 10)
	expected = "grynotvhlk"
	if (output != expected) {
			t.Errorf("Got: %s, expected: %s", output, expected)
	}

} // End of TestStringLowerN()


func TestRandomIntn(t *testing.T) {

	seed := "/foobar"
	random := New()

	expected := []uint{38171, 815034, 971458}
	for i:=0; i<len(expected); i++ {
		result := random.Intn(seed, 1000000)
		row := expected[i]
		if (result != row) {
			t.Errorf("%d != %d", result, row)
		}
		seed += "1"
	}

	expected = []uint{614807, 777383, 679897}
	for i:=0; i<len(expected); i++ {
		result := random.Intn(seed, 1000000)
		row := expected[i]
		if (result != row) {
			t.Errorf("Got: %d, expected: %d", result, row)
		}
		seed += "1"
	}

	//
	// Edge case
	//
	seed = "12345"
	expected = []uint{0, 0, 0}
	for i:=0; i<len(expected); i++ {
		result := random.Intn(seed, 1)
		row := expected[i]
		if (result != row) {
			t.Errorf("Got: %d, expected: %d", result, row)
		}
	}

	//
	// Using a max of zero should panic.
	// Catch that panic and make sure we didn't not panic.
	//
	beenhere := false

	defer func() {
		if (beenhere) {
			panic("We should have had a panic earlier, but didn't.")
		}
		recover()
	}()


	expected = []uint{0, 0, 0}
	for i:=0; i<len(expected); i++ {
		result := random.Intn(seed, 0)
		row := expected[i]
		if (result != row) {
			t.Errorf("Got: %d, expected: %d", result, row)
		}
	}

	beenhere = true

} // End of TestRandomIntn()



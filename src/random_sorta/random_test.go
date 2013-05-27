
package random_sorta

//import "fmt"
import "testing"


func TestStringN(t *testing.T) {

	random := New(0, 1000000)
	output := random.StringN(1)
	expected := "P"
	if (output != expected) {
			t.Errorf("Got: %s, expected: %s", output, expected)
	}

	output = random.StringN(40)
	expected = "jeBTUGGjSxduTEVHkWCqqmYNnDGComdckLBFsslh"
	if (output != expected) {
			t.Errorf("Got: %s, expected: %s", output, expected)
	}

	random = New(12345, 1000000)
	output = random.StringN(10)
	expected = "CqjpmnVJwT"
	if (output != expected) {
			t.Errorf("Got: %s, expected: %s", output, expected)
	}

} // End of TestStringLowerN()


func TestStringLowerN(t *testing.T) {

	random := New(0, 1000000)
	output := random.StringLowerN(1)
	expected := "p"
	if (output != expected) {
			t.Errorf("Got: %s, expected: %s", output, expected)
	}

	output = random.StringLowerN(10)
	expected = "dbtuggdusr"
	if (output != expected) {
			t.Errorf("Got: %s, expected: %s", output, expected)
	}

	random = New(12345, 1000000)
	output = random.StringLowerN(10)
	expected = "ckdjgvhvjq"
	if (output != expected) {
			t.Errorf("Got: %s, expected: %s", output, expected)
	}

} // End of TestStringLowerN()


func TestRandomIntn(t *testing.T) {

	random := New(0, 1000000)
	expected := []uint{52687, 817315, 787998}
	for i:=0; i<len(expected); i++ {
		result := random.Intn()
		row := expected[i]
		if (result != row) {
			t.Errorf("%d != %d", result, row)
		}
	}

	random = New(12345, 1000000)
	expected = []uint{752770, 447658, 316259}
	for i:=0; i<len(expected); i++ {
		result := random.Intn()
		row := expected[i]
		if (result != row) {
			t.Errorf("Got: %d, expected: %d", result, row)
		}
	}

	//
	// Edge case
	//
	random = New(12345, 1)
	expected = []uint{0, 0, 0}
	for i:=0; i<len(expected); i++ {
		result := random.Intn()
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


	random = New(12345, 0)
	expected = []uint{0, 0, 0}
	for i:=0; i<len(expected); i++ {
		result := random.Intn()
		row := expected[i]
		if (result != row) {
			t.Errorf("Got: %d, expected: %d", result, row)
		}
	}

	beenhere = true

} // End of TestRandomIntn()


func TestRandomIntnChannel(t *testing.T) {

	random := New(0, 1000000)

	in := make(chan uint)
	out := make(chan []uint)

	go random.IntnChannel(in, out)

	in <- 3
	results := <-out

	expected := []uint{52687, 817315, 787998}
	for i:=0; i<len(expected); i++ {
		result := results[i]
		row := expected[i]
		if (result != row) {
			t.Errorf("%d != %d", result, row)
		}
	}


	in <- 3
	results = <-out

	expected = []uint{752770, 447658, 316259}
	for i:=0; i<len(expected); i++ {
		result := results[i]
		row := results[i]
		if (result != row) {
			t.Errorf("Got: %d, expected: %d", result, row)
		}
	}

} // End of TestRandomIntnChannel()



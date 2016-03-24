package ohyeah

import (
	"fmt"
	"log"
	"testing"
)

func TestConst(t *testing.T) {
	f := ConstantGen(44)

	for i := 0; i < 100; i++ {
		if v := f(); v != 44 {
			t.Errorf("Expected generator to produce 44, but got %v", v)
		}
	}
}

func TestCycle(t *testing.T) {
	a := ConstantGen(44)
	b := ConstantGen("hello")
	f := CycleGen(a, b)

	for i := 0; i < 100; i++ {
		if v1 := f(); v1 != 44 {
			t.Errorf("Expected generator to produce 44, but got %v", v1)
		}

		if v2 := f(); v2 != "hello" {
			t.Errorf(`Expected generator to produce "hello", but got %v`, v2)
		}
	}
}

func TestRandom(t *testing.T) {
	f := RandomFunc(373)

	numEqual := 0
	lastValue := int64(0)

	// Count the number of random numbers that are
	// equal to the previous number. Check to see if that
	// is less than 10% of the bunch. This is an extremely weak
	// randomness check.

	for i := 0; i < 100; i++ {
		value := f()
		if value == lastValue {
			numEqual++
		}
	}

	if numEqual > 10 {
		t.Errorf("More than 10% of the random numbers are equal to the previous number!")
	}
}

func TestPatternedStrings(t *testing.T) {
	var f Generator
	f = PatternedStringGen("foo")

	for i := 0; i < 100; i++ {
		value := f()
		expected := fmt.Sprintf("foo%d", i+1)
		if value != expected {
			t.Errorf("Expected value %v not == to actual %v", expected, value)
		}
	}
}

func XXTestStrings(t *testing.T) {
	r := RandomFunc(99)
	var f Generator
	f = StringGen(r)

	for i := 0; i < 100; i++ {
		log.Println(f())
	}
}

func XXTestValues(t *testing.T) {
	r := RandomFunc(99)

	intF := IntGen(r)
	sF := StringGen(r)

	f := RandomGen(r, intF, sF, sF)

	for i := 0; i < 100; i++ {
		log.Println(f())
	}
}

func XXTestArray(t *testing.T) {
	r := RandomFunc(99)

	intF := IntGen(r)
	f := ArrayGen(r, intF, 1)

	for i := 0; i < 100; i++ {
		log.Println(f())
	}
}

func XXTestMap(t *testing.T) {
	r := RandomFunc(99)

	kf := PatternedStringGen("key")
	vf := IntGen(r)

	f := MapGen(r, kf, vf, 5)

	for i := 0; i < 100; i++ {
		log.Println(f())
	}
}

func XXTestElement(t *testing.T) {
	r := RandomFunc(99)

	a := []interface{}{"russ", "olsen", "1234", "hello"}

	f := ElementGen(r, a)

	for i := 0; i < 100; i++ {
		log.Println(f())
	}
}

func XXTestBool(t *testing.T) {
	r := RandomFunc(99)
	f := BoolGen(r)

	for i := 0; i < 100; i++ {
		log.Println(f())
	}
}

func XXTestCrazy(t *testing.T) {
	r := RandomFunc(99)
	z := MapGen(r, PatternedStringGen("key"), ArrayGen(r, IntGen(r), 10), 25)

	for i := 0; i < 100; i++ {
		log.Println(z())
	}
}

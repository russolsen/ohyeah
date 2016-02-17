package ohyeah

import(
	"fmt"
)


type Int64F func()int64

// RandomFunc returns a simple int64 random number generator based
// on the given seed. The idea here is not to produce a great series
// of random numbers but rather a highly reproducable one, over differnt
// programming languages.
func RandomFunc(seed int64) Int64F {
	current := seed
	modValue := int64((1 << 31) - 1)
	g := int64(16807)

	return func() int64 {
		result := current
		current = (g * result) % modValue
		return result
	}
}


type Generator func() interface{}

func ConstantGen(x interface{}) Generator {
	return func() interface{} {
		return x
	}
}


func IntGen(r Int64F) Generator {
	return func() interface{} {
		return r()	
	}
}

func IntN(r Int64F, n int) int {
	return int(r() % int64(n))
}

func BoolGen(r Int64F) Generator {
	return func() interface{} {
		i := IntN(r, 2)
		return i == 0	
	}
}

func Float64Gen(r Int64F) Generator {
	return func() interface{} {
		i := r()
		j := r()
		return float64(i) / float64(j)
	}
}

func StringGen(r Int64F) Generator {
	return func() interface{} {
		l := IntN(r, 10)
		ba := make([]byte, l)

		for i := 0; i < l; i++ {
			ba[i] = byte(IntN(r, 100) + 12)
		}

		return string(ba)
	}
}

func PatternedStringGen(r Int64F, prefix string)Generator {
	i := 0
	return func() interface{} {
		i++
		return fmt.Sprintf("%s%d", prefix, i)
	}
}

// RandomGen picks a random generator from generators and evaluates it.
func RandomGen(r Int64F, generators ...Generator)Generator {
	return func() interface{} {
		i := IntN(r, len(generators))
		return generators[i]()	
	}
}

// CycleGen cycles thru all of the generators.
func CycleGen(generators ...Generator)Generator {
	if len(generators) == 0 {
		panic("No generators supplied to CycleGen")
	}

	i := 0
	return func() interface{} {
		result := generators[i]()
		i = (i + 1) % len(generators)
		return result
	}
}

// RepeatGen generates a repeated pattern of values. The first n times
// the repeat gen is called it will call g and return the values. On
// call n+1 the generator will return the first value, then the second
// and so on.

func RepeatGen(g Generator, n int)Generator {
	i := 0
	savedValues := make([]interface{}, n)

	return func() interface{} {
		if i < n {
			savedValues[i] = g()
		}
		result := savedValues[i%n]
		i++
		return result
	}
}

func ArrayGen(r Int64F, g Generator, maxLen int)Generator {
	return func() interface{} {
		l := IntN(r, maxLen + 1)
		result := make([]interface{}, l)
		for i, _ := range result {
			result[i] = g()
		}
		return result
	}
}


func MapGen(r Int64F, k Generator, v Generator, maxLen int)Generator {
	return func() interface{} {
		l := IntN(r, maxLen + 1)
		result := map[interface{}]interface{}{}

		for i :=0; i < l; i++ {
			result[k()] = v()
		}
		return result
	}
}

// ElementGen returns a generator that will produce random elements from the array.
func ElementGen(r Int64F, array []interface{}) Generator{
	return func() interface{} {
		i := IntN(r, len(array))
		return array[i]
	}
}



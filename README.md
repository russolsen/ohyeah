# ohyeah

> Oh YEAH? Well could your program handle an arrray of arrays of
> maps of strings to arrays filled with booleans? Huh?

Ohyeah is a small Go library that implements a set of functions that generate random data
and data structures, primarily intended for generating test data.

## Usage

The key abtraction in ohyeah is the `Generator`, which is a parameterless function
that produces a value:

```go
type Generator func() interface{}
```

Generators typically produce a series of randomly generated values, a new
one each time you call the function. One thing to note is that most
of the functions provided by ohyeah are not themselves generators: Instead
they are higher order functions that return generators. 
For example, the `PatternedStringGen` function returns a function
which produces a series of strings all of the form "<PAT><N>" where
<PAT> is the sring you pass to `PatternedStringGen` and <N> is an integer
that starts at 1.

Thus this:

```go
	idGen := ohyeah.PatternedStringGen("ID")
	nameGen := ohyeah.PatternedStringGen("Fred")
```

Will give you a function in `idGen` that will generated a sequence of ID strings:

```go
	a := idGen()  // a is "ID1"
	b := idGen()  // b is "ID2"

	n1 := nameGen()  // That's "Fred1"
	n2 := nameGen()  // That's "Fred2"
```

If you are not used to this style of programming, the whole function producing
a function thing may seem overly abstract, but it has a huge practical advantage:
Once you have set up your generator function -- the way we did with `idGen` and
`nameGen` above -- you no longer need to worry about what or how it is doing
it's thing. You just call the function.
 
While there isn't really much randomness coming out of `PatternedStringGen`,
many of the other ohyeah functions do and those functions all take a
random iterger generating function defined as:

```go
type Int64F func()int64
```

This function can be any function that returns a series of random int64's.
The ohyeah package supplies an higher order function to produce just
such a random function in the form of `RandomFunc`.

## Example

To make this all a bit more concrete, here is a generator that produces an 
map of strings => arrays filled with, well, _stuff_.


```go
	// Random function we will use through out.

	r := ohyeah.RandomFunc(99)

	// Generator which picks random element from the array.


	strs := []interface{}{"foo", "bar", "baz", "apple", "organge", "red", "x"}
	strs_gen := ohyeah.ElementGen(r, strs)

	// Generator which returns random ints. 

	ints_gen := ohyeah.IntGen(r)

	// Generator which always just returns true

	true_gen :=  ohyeah.ConstantGen(true)

	// Generator which returns "Key1", "Key2" ...

	key_gen := ohyeah.PatternedStringGen("Key")

	// Generator which will return values from the other generators

	value_gen := ohyeah.CycleGen(true_gen, key_gen, ints_gen, strs_gen)

	// Generator which will return randomly sized arrays of values, max len = 10

	array_gen := ohyeah.ArrayGen(r, value_gen, 10)

	// Generator which will return randomly sized maps of "KeyNNN" => array

	map_gen := ohyeah.MapGen(r, key_gen, array_gen, 10)
```

## Copyright and License
Copyright © 2016 Russ Olsen

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at
http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

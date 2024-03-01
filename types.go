package main

type intOrString interface { //abandon interface because methods cannot take type interface lists/generics
	~int32 | ~string
}

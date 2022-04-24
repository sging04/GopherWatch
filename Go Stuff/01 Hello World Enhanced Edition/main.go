package main

/*
	Patrick and Sean Ging, Period 1

	01 : Hello, world! Enhanced Edition.

	Description:

		Says hello to every argument in the Go function

	Glossary:
	
		fmt: the formatting package for printing, some strings, etc
		fmt.Printf: a printf function similar to Java, C, C++, Python3's String.format()

		range: Same as Python3's enumerate. Returns an iterable with the index and the item of an iterable.

		os: package relating to broader os stuff
		os.Args: the arguments of the function. The 0th item here is the name of the function "main".

	Notes:

		What is "_"?
			
			In GoLang, we're not allowed to declare variables or packages we don't use. So if we have a function that
			returns a value we won't use, we use the "_" character to tell the compiler we're not using it, allowing us to 
			avoid throwing an error.
*/



import (
	"fmt"
	"os"
	)

func main(){

	for _, item := range os.Args[1:] {
		fmt.Printf("Hello, %v\n", item)
	}

}

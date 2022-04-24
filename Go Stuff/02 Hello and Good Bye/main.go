package main

import (
	"fmt"
	"os"
	)

/*
	Patrick and Sean Ging, Period 1

	02 : Hello and Good Bye

	Description:

		Says hello to every odd numbered person and goodbye to every even numbered person.

	Glossary:
		
		len() : returns the length of an iterable

	Notes: 
		Why no newline inbetween if & else? 
			Explicit Choice made in Go. Imo it's to maintain readability (mostly consistency)

		Why does the 1st item in os.Args say hello? 
			Because we're not iterating through os.Args, we're iterating through
			os.Args[1:], meaning 1st item in os.Args is 0th in os.Args[1:]

		Why len(os.Args[1:]) > 1 ? 
			In os.Args, the 0th item is the name of the package, in this case main. Everything
			after that are command line arguments.
*/

func hello(name string, number int) {
	fmt.Printf("Hello, %v! You're person number %v\n", name, number)
}

func goodbye(name string, number int) {
	fmt.Printf("Goodbye, %v! You're person number %v\n", name, number)
}


func main(){
	if len(os.Args[1:]) > 1 {
		for index, name := range os.Args[1:] {
			if index % 2 != 0 {
				goodbye(name, index)
			} else { /*Note: You cannot have a newline inbetween if & else.*/
				hello(name, index)		
			}
		}
	} else {
		fmt.Println("Nobody's to greet : ( ")
	}

}

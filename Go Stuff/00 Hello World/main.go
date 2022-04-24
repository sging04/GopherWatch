package main

/*
	Patrick and Sean Ging, Period 1

	00 : Hello, World!

	Description:

		Says Hello, World!

	Glossary:
	
		fmt: the formatting package for printing, some strings, etc
		fmt.Printf: a printf function similar to Java, C, C++, Python3's String.format()

	Notes:
		
		Is ; mandatory?
			No, it isn't. Both of them are legal. Sometimes though, using the ; will make things
			more readable & also for loops, if statements, etc.
*/

import (
	"fmt"
	)

func main() {
	fmt.Printf("Hello, World")
	fmt.Printf("Hello, World");

}

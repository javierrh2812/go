package main // required for a standalone executable
// each go application contains one main

import (
	"fmt"
	"strconv"
	"time"
) // fmt implements formatted I/O
// import loads the public declarations from the compiled package

/*
import multiple packages

1.
import "fmt"
import "os"

2.
import "fmt"; import "os"

3.
import (
	"fmt"
	"os"
)
*/

/*
using alias:
import fm "fmt"
*/

/*
VARIABLE DECLARATION
var var1 type1 /// type1 is the type of var1
*/

var a int// declare a single variable

var ( // declare a group of variables
	b bool
	c float32
	d string
)

// when this program is executed, the firts function
// that runs is main.main()
func main (){// main is an entry point
	fmt.Print("Hello, world!\n")
	//elementaryTypes()
	//timeType()
	handlingError()
} // exiting the program

func elementaryTypes() {

	/*
	only declared values contains the default zero value depending upon its type
	*/
	fmt.Println(a)
	a = 42 				// assign single value
	b,c = true, 32.0 	// assign multiple
	d = "string" 		// strings must contain double quotes

	//REFACTORED (creating values without declaring them)
	a2 := 42 // initialize and assign to a single value
	b2, c2 := true, 32.4 // initialize and assign to multiple variables
	d2 := "string" // initializa and assign to a string
	fmt.Println("hi mom!")
	fmt.Println(a,b,c,d)
	fmt.Println(a2,b2,c2,d2)
	fmt.Println("casting c2", int(c2))
	/* User specified types */
	const a int32 = 12         // 32-bit integer
	const b float32 = 20.5      // 32-bit float
	var c complex128 = 1 + 4i  // 128-bit complex number
	var d uint16 = 14          // 16-bit unsigned integer

	/* Default types */
	n := 42              // int
	pi := 3.14           // float64
	x, y := true, false  // bool
	z := "Go is awesome" // string

	fmt.Printf("user-specified types:\n %T %T %T %T\n", a, b, c, d)
	fmt.Printf("default types:\n %T %T %T %T %T\n", n, pi, x, y, z)


	var ch byte = 'a'
	var ch1 int = '\u0041'
    var ch2 int = '\u03B2'
    var ch3 int = '\U00101234'
	fmt.Println(ch)
    fmt.Printf("%T - %T - %T\n", ch1, ch2, ch3)  // integer
    fmt.Printf("%d - %d - %d\n", ch1, ch2, ch3)  // integer
    fmt.Printf("%c - %c - %c\n", ch1, ch2, ch3)  // character
    fmt.Printf("%X - %X - %X\n", ch1, ch2, ch3)  // UTF-8 bytes
    fmt.Printf("%U - %U - %U", ch1, ch2, ch3)    // UTF-8 code poin 
}

func structuredTypes(){
	// struct
	// array
	// slice
	// map
	// channel
}

var week time.Duration

func timeType(){
	t := time.Now()
    fmt.Println(t)
    fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
    t = time.Now().UTC()
    fmt.Println("UTC =>", t)
    // calculating times:
    week = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
    week_from_now := t.Add(week)
    fmt.Println("in 1 week => ", week_from_now)
    // formatting times:
    fmt.Println(t.Format("02 Jan 2006 15:04"))
    s := t.Format("2006 01 02")
    fmt.Println(t, "=>", &s)
}
func handlingError() {
    var orig string = "ABC"
    var an int
    var err error
    // storing integer and error information
    an, err = strconv.Atoi(orig)

    if err != nil { // if it was an error, discontinue
        fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
        return
    }
    fmt.Printf("The integer is %d\n", an)
    // rest of the code
}

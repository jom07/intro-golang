// Functions are visible in...
package main

// We use...
import (
	"fmt"
	"math"
	"reflect"
)

// Main entry point
func main() {
	// Printing strings, escape characters, runes
	fmt.Println("Hello, Go!")
	fmt.Println("Runes vs string literals \n", '\n', '\t', '\\', "\n\"")
	fmt.Println(reflect.TypeOf(3))
	fmt.Println(reflect.TypeOf(3.2))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf('e'))
	fmt.Println(reflect.TypeOf("hello"))

	// Variable declaration and assignment
	var num1, num2 int = 2, 3
	var mult1, mult2 float64
	var div1 = 5.3
	var res, flooredRes, roundedRes float64
	var printRes string

	// Variable assignment
	mult1, mult2 = 3.5, 4.5

	// Some math
	res = (float64)(num1+num2) * mult1 * mult2 / div1
	flooredRes = math.Floor(res)
	roundedRes = math.Round(res)
	printRes = fmt.Sprint(res) + " floored is " + fmt.Sprint(flooredRes) + " and rounded is " + fmt.Sprint(roundedRes)
	fmt.Println(printRes)

	// Empty or zero values
	var aBool bool
	var aChar int32
	var aString string
	var aFloat float64

	fmt.Println(aBool)
	fmt.Println(aChar)
	fmt.Println(aString)
	fmt.Println(aFloat)

	// Short variable declaration and assignment
	var1, var2 := 2.2, 2.5
	name1 := "John Doe"
	fmt.Println(var1, var2, name1)

	// Careful... What int do you want?
	var a int8 = 16
	var b int16 = 8192
	var c int32 = 2_000_000
	var d int64 = 9_223_000_000_000_000
	// Seems like 64 bit to me
	var e int = 9_223_000_000_000_000_000

	fmt.Println(a, b, c, d, e)

	f := 'k'
	var g int32 = 'k'
	fmt.Println(reflect.TypeOf(f), f, reflect.TypeOf(g), g)

}

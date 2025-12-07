package main
import "fmt"

func main() {
	// integer
	// signed int
	var i int 
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64

	i=-3
	i8=-89
	i16=-898
	i32=-898998
	i64=-89899839898983

	// unsigned int
	var ui uint
	var ui8 uint8
	var ui16 uint16
	var ui32 uint32
	var ui64 uint64

	ui=3
	ui8=89
	ui16=8989
	ui32=89899
	ui64=89899839898983

	fmt.Println("signed integers",i, i8, i16, i32, i64)
	fmt.Println("unsigned integers",ui, ui8, ui16, ui32, ui64)

	// floating point
	// float32
	var f32 float32=10332332233443432
	// float64
	var f64 float64=10332332233443432
	fmt.Println("float32:",f32)
	fmt.Println("float64:",f64)

	// boolean
	var b1 bool=true
	var b2 bool=false
	fmt.Println("boolean1:",b1)
	fmt.Println("boolean2:",b2)

	// complex
	var c1 complex64=complex(3,4)
	var c2 complex128=complex(5,6)
	fmt.Println("complex64:",c1)
	fmt.Println("complex128:",c2)

	// strings
	var name string="Mekuannent"
	fmt.Println("my name is:",name)
}
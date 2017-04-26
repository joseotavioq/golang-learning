package main

func primitiveDataTypes() {
	//int
	var myInt int
	myInt = 33
	println(myInt)

	//float32
	var myFloat32 float32 = 33.
	println(myFloat32)

	//string
	myString := "Hi"
	println(myString)

	//complex
	myComplex := complex(2, 3)
	println(myComplex)
	println(real(myComplex))
	println(imag(myComplex))
}

/*
//Basic Types
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // pseudônimo para uint8

rune // pseudônimo para int32

float32 float64

complex64 complex128
*/

// Prints
// 33
// +3.300000e+001
// Hi
// (+2.000000e+000+3.000000e+000i)
// +2.000000e+000
// +3.000000e+000

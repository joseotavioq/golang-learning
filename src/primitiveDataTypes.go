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

// Prints
// 33
// +3.300000e+001
// Hi
// (+2.000000e+000+3.000000e+000i)
// +2.000000e+000
// +3.000000e+000
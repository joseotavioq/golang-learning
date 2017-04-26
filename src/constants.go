package main

const (
	simpleString = "Hello"
)

const (
	iotaExample = iota
	iotaExample1
	iotaExample2
)

const (
	tenTimesFirst = 1 << (10 * iota)
	tenTimesSecond
	tenTimesThird
)

func constants() {
	println(simpleString)

	println(iotaExample)
	println(iotaExample1)
	println(iotaExample2)

	println(tenTimesFirst)
	println(tenTimesSecond)
	println(tenTimesThird)
}

// Prints
// Hello
// 0
// 1
// 2
// 1
// 1024
// 1048576
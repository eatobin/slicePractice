package main

import "fmt"

var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
var aSlice, bSlice []byte
var cSlice []byte

func myMap() {
	// Initialize a map
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	// map has two return values. For the second return value, if the key doesn't
	// exist，'ok' returns false. It returns true otherwise.
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}
}

func main() {
	aSlice = array[:3]
	println(string(aSlice))

	aSlice = array[5:]
	println(string(aSlice))

	aSlice = array[:]
	println(string(aSlice))

	aSlice = array[3:7]
	println(string(aSlice))
	println(len(aSlice))
	println(cap(aSlice))

	bSlice = aSlice[1:3]
	println(string(bSlice))

	bSlice = aSlice[:3]
	println(string(bSlice))

	bSlice = aSlice[0:5]
	println(string(bSlice))

	println(string(aSlice[0:5]))

	println(string(aSlice[0:7]))

	cSlice = append(aSlice, 'k')
	println(string(cSlice))

	println(len(cSlice))
	println(cap(cSlice))
	println(string(cSlice[:7]))

	aSlice[0] = 'z'
	println(string(cSlice[:7]))
	println(string(aSlice[:7]))

	cSlice[1] = '1'
	println(string(cSlice))
	println(string(aSlice))

	myMap()

	//delete(rating, "C") // delete element with key "c"
}

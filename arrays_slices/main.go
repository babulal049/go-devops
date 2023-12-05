/*

Slice has 3 major components - pointer, length & capacity

len() length - number of elements slice contains
cap() capacity - start point from slice to endpoint of an underlying array
slice = array[0:3] ==> typically indicates slaice = array[min-index:max-index-1]
create slice using make function
In arrays capacity is equal to length of an array
when slice is part of an array, if you change value in slice, it will effect array as well
append function --> slice = append(slice, element1, element2)

good example of handling capacity when slices are part of arrays:
array  = 10 20 30 40
slice  = array[1:3] = 20 30, len(slice) = 2, cap(slice) = 3
slice1 = append(slice, 900) = 20 30 900 , len(slice1) = 3, cap(slice1) =3
slice2 = append(slice, -90, 500) -> since cap(slice) is 3, you can add only 1 item, but here you are trying to add 2, so it will double capacity of slice , cap(slice) is 6 now

appending slice with another slice, slice = append(slice, anotherslice...)
looping through slice --> for index, value := range arr
looping through slice --> for _, value := range arr
*/

package main

import "fmt"

func main() {
	// decalring arrays
	var grades [5]int
	fmt.Println(grades)
	var new_grades [3]int = [3]int{10, 20, 30}
	fmt.Println(new_grades)
	new_grades2 := [3]int{10, 20, 30}
	fmt.Println(new_grades2)
	new_grades3 := [...]int{10, 20, 30, 40, 50}
	fmt.Println(new_grades3)
	fmt.Println("----------------------------------")

	//length of an array
	fmt.Println(len(new_grades3))
	fmt.Println("----------------------------------")

	//looping through an array, normal & using range
	fmt.Println("Normal looping")
	for i := 0; i < len(new_grades3); i++ {
		fmt.Println(i, "-->", new_grades3[i])
	}
	fmt.Println("Looping using range keyword")
	for index, element := range new_grades3 {
		fmt.Println(index, "-->", element)
	}
	fmt.Println("----------------------------------")

	fmt.Println("Slice Tasks")
	// Declaring and initializing a slice
	new_slice := []int{10, 20, 30}
	fmt.Println(new_slice)
	fmt.Println("----------------------------------")

	//Creating a slice from an array
	mango := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println(mango)
	mango_slice := mango[1:9]
	fmt.Println(mango_slice)
	fmt.Println("----------------------------------")

	//Creating a slice using make function
	mango_slice2 := make([]int, 10, 12) //make([]<data_type>, length, capacity)
	fmt.Println(mango_slice2)

	//Appending slices
	mango_slice = append(mango_slice, 00, 01)
	fmt.Println(mango_slice)
	new_mango_slice := append(mango_slice, mango_slice2...)
	fmt.Println(new_mango_slice)
	fmt.Println("----------------------------------")

	//looping slices
	for index, value := range new_mango_slice {
		fmt.Println(index, "=>", value)
	}
	fmt.Println("----------------------------------")
	for _, value := range new_mango_slice {
		fmt.Println(value)
	}

}

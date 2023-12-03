package main

import "fmt"

func main() {
	/* Task1 & 2 :lets create 2 maps in 2 different ways */

	subject_marks := map[string]int{"English": 40, "Maths": 30} // Normal declaration - map[string]int ==> here "Key' is of string type, "Value" is of INT type
	exam_attempts := make(map[string]int) // Using make  command
	exam_attempts["English"] = 2
	exam_attempts["Maths"] = 3
	fmt.Println(subject_marks)
	fmt.Println(exam_attempts)
	fmt.Println("=============================")
	
  //3. Print length of maps
	fmt.Println(len(subject_marks))
	fmt.Println(len(exam_attempts))
	fmt.Println("=============================")
	
  //Accessing map
	fmt.Println(subject_marks["English"])
	fmt.Println(subject_marks["Maths"])
	fmt.Println(exam_attempts["English"])
	fmt.Println(exam_attempts["Maths"])
	fmt.Println("=============================")
	
  //Check if value is present or not basing on key
	value1, foundornot := subject_marks["English"]
	fmt.Println(value1, foundornot)
	value2, foundornot2 := subject_marks["Telugu"]
	fmt.Println(value2, foundornot2)
	fmt.Println("=============================")
	
  //Adding or updating key-value pair
	subject_marks["Spanish"] = 60
	subject_marks["English"] = 60
	fmt.Println(subject_marks)
	fmt.Println("=============================")
	
  //Deleting key-value pair
	delete(subject_marks, "Spanish")
	fmt.Println(subject_marks)
	fmt.Println("=============================")
	
  //Using for loops: Syntax ==> { for key, value := range map_name}
	for key, value := range subject_marks {
		fmt.Println(key, "=>", value)
	}
}

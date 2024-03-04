package main

import "fmt"

// Time to practice what you learned!
func main() {
	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.
	hobbies := [3]string{"Rock Climbing", "Coding", "Piano"}
	fmt.Println(hobbies)

	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list
	fmt.Println(hobbies[0])
	hobbies2 := [2]string{hobbies[1], hobbies[2]}
	fmt.Println(hobbies2)
	
	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	slice := hobbies[0:1]
	slice = hobbies[:1]
	fmt.Println(slice)
	
	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	slice = slice[1:2]
	fmt.Println(slice)
	
	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	goals := []string{"Become the greatest Go programmer ever", "Learn to hate Golang"}
	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	goals[1] = "Learn to love Golang"
	goals = append(goals, "Create something cool!")
	fmt.Println(goals)

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	type Product struct {
		title string
		id string
		price float64
	}

	products := []Product{{"Milk", "0001", 3.99}, {"Eggs", "0002", 4.99}}
	products = append(products, Product{"Tomato", "0003", 0.99})
	fmt.Println(products)
}


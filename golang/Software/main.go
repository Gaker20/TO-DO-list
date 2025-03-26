package main

import "fmt"

func main() {
	fmt.Println("ECU Form")
	fmt.Println("Enter your first name")
	var username string
	fmt.Scan(&username)

	fmt.Print("Enter student ID:")
	var userID int
	fmt.Scan(&userID)

	fmt.Print("Enter your major:")
	var userCourse string
	fmt.Scan(&userCourse)

	fmt.Printf("Your name is %v, your student ID is %v, and you are majoring in %v", username, userID, userCourse)
}

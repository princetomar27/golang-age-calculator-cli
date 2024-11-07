package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	// Step 1 : Ask user to enter your DOB
	fmt.Println("Please enter your date of birth in (format : YYYY-MM-DD) : ")

	// Step 2 : Read input from user

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error occurred while reading input : \n%s", err)
		return
	}

	// Trim extra space
	input = strings.TrimSpace(input)

	// 3. Parse the input date
	dob, err := time.Parse("2006-01-02", input)
	if err != nil {
		fmt.Println("Invalid date format. Please use the format YYYY-MM-DD.")
		return
	}

	// step 4 : get the current date
	currentDate := time.Now()

	// step 5 : calculate the age
	age := currentDate.Year() - dob.Year()

	if currentDate.Month() < dob.Month() || (currentDate.Month() == dob.Month() && currentDate.Day() < dob.Day()) {
		age--
	}

	// step 6 : Display the age
	fmt.Printf("You are %d years old.\n", age)

}

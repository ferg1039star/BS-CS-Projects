package main

import "fmt"

func main() {

	for {
		fmt.Print("Enter 'q' to exit the program or " +
			"Enter a number to FizzBuzz: ")

		var input string
		fmt.Scanln(&input)

		if input == "q" || input == "Q" {
			fmt.Println("Exiting program")
			break
		}
		var n int
		_, err := fmt.Sscanf(input, "%d", &n)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number or press 'q' to exit")
			continue
		}

		if n <= 0 {
			fmt.Println("Please enter a positive integer")
			continue
		}

		fizzBuzz(n)

	}
}

func fizzBuzz(n int) {

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
	fmt.Println("---- End of FizzBuzz for", n, "----")
}

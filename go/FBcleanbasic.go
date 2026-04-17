package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {

		var output = ""

		if i%3 == 0 && i%5 == 0 {
			output += "FizzBuzz"
		}
		if i%3 == 0 {
			output += "Fizz"
		}
		if i%5 == 0 {
			output += "Buzz"
		}
		if output == "" {
			output = fmt.Sprintf("%d", i)

		}

		fmt.Println(output)
	}

}

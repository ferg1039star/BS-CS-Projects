package main

import "fmt"

type Rule struct {
	Number int
	Word   string
}

type FizzBuzzConfig struct {
	Limit int
	Rules []Rule
}

func main() {

	config := FizzBuzzConfig{
		Limit: 100,
		Rules: []Rule{
			{Number: 3, Word: "Fizz"},
			{Number: 5, Word: "Buzz"},
		},
	}

	for i := 1; i <= config.Limit; i++ {
		output := ""

		for _, rule := range config.Rules {
			if i%rule.Number == 0 {
				output += rule.Word

			}
		}

		if output == "" {
			output = fmt.Sprintf("%d", i)
		}

		fmt.Println(output)
	}
}

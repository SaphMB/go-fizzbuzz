package fizzbuzz

import "fmt"

func RespondSlowly(input int) string {
	switch input {
	case 1:
		return fmt.Sprint(1)
	case 2:
		return fmt.Sprint(2)
	case 3:
		return "Fizz"
	case 4:
		return fmt.Sprint(4)
	case 5:
		return "Buzz"
	case 6:
		return "Fizz"
	case 7:
		return fmt.Sprint(7)
	case 8:
		return fmt.Sprint(8)
	case 9:
		return "Fizz"
	case 10:
		return "Buzz"
	case 11:
		return fmt.Sprint(11)
	case 12:
		return "Fizz"
	case 13:
		return fmt.Sprint(13)
	case 14:
		return fmt.Sprint(14)
	case 15:
		return "FizzBuzz"
	case 30:
		return "FizzBuzz"
	}

	return ""
}

func Respond(input int) string {
	if input%15 == 0 {
		return "FizzBuzz"
	}

	if input%3 == 0 {
		return "Fizz"
	}

	if input%5 == 0 {
		return "Buzz"
	}

	return fmt.Sprint(input)
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var choice int
	for {
		fmt.Println("1. Sum Values")
		fmt.Println("2. Convert Uppercase")
		fmt.Println("3. Sum Odd Numbers")
		fmt.Println("4. Order the Alphabet")
		fmt.Println("5. Spilt Words")
		fmt.Println("0. Exit")
		fmt.Print("Choice the option :")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			sum()
		case 2:
			uppercase()
		case 3:
			sumOddNumbers()
		case 4:
			ascendingAlphabet()
		case 5:
			splitWords()
		case 0:
			return
		default:
			fmt.Println("Invalid Option...")
		}
	}
}

// sum of two integer
func sum() {
	var firstNumber, secondNumber int
	fmt.Print("Enter the first number : ")
	_, firstErr := fmt.Scanln(&firstNumber)
	if firstErr != nil {
		fmt.Println("Please Enter Integer Value....")
		return
	}
	fmt.Print("Enter the second number : ")
	_, secondErr := fmt.Scanln(&secondNumber)
	if secondErr != nil {
		fmt.Println("Please Enter Integer Value....")
		return
	}
	sum := firstNumber + secondNumber
	fmt.Println("Sum of Two Integer is : ", sum)
}

// convert lowercase to uppercase
func uppercase() {
	var trigger bool = false
	fmt.Println("Enter the string : ")
	reader := bufio.NewReader(os.Stdin)
	inputs, _ := reader.ReadString('\n')
	for _, inputData := range inputs {
		if inputData >= 'a' || inputData == '\n' || inputData == ' ' {
			trigger = true
		} else {
			trigger = false
			break
		}
	}
	if trigger {
		Stringvalues := []byte(inputs)
		for i, input := range inputs {
			if input >= 'a' && input <= 'z' {
				Stringvalues[i] = byte(input) - 32
			}
		}
		fmt.Println(string(Stringvalues))
		return
	}
	fmt.Println("Please enter string value")
}

// sum of odd numbers
func sumOddNumbers() {
	var values [5]int = [5]int{}
	var total int
	fmt.Println("Enter", len(values), "numbers : ")
	for i := 0; i < len(values); i++ {
		fmt.Print("Element ", i+1, " : ")
		_, err := fmt.Scanln(&values[i])
		if err != nil {
			fmt.Println("Please Enter Integer Value")
			return
		}
	}
	for j := 0; j < len(values); j++ {
		if values[j]%2 == 1 {
			total = total + values[j]
		}
	}
	fmt.Println("Sum of all odd values : ", total)
}

// Sort a string in ascending order
func ascendingAlphabet() {
	var trigger bool = false
	fmt.Println("Enter the string : ")
	reader := bufio.NewReader(os.Stdin)
	inputs, _ := reader.ReadString('\n')
	for _, inputData := range inputs {
		if inputData >= 'a' || inputData == '\n' || inputData == ' ' {
			trigger = true
		} else {
			fmt.Println("Please Enter String value Only...")
			return
		}
	}
	if trigger {
		ascending(inputs)
	}
}

func ascending(inputs string) {
	input := []byte(inputs)
	for i := 1; i < len(inputs)-1; i++ {
		if input[i] < input[i-1] {
			input[i], input[i-1] = input[i-1], input[i]
			i = 0
		}
	}
	fmt.Print("Ascending order  : ", string(input))
}

// Split the Words
func splitWords() {

	fmt.Print("Input a string: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	words := splitBySpace(input)
	fmt.Println("After Split :")
	for _, word := range words {
		fmt.Println(word)
	}
}
func splitBySpace(s string) []string {
	var words []string
	word := ""
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			word += string(s[i])
		} else if word != "" {
			words = append(words, word)
			word = ""
		}
	}
	if word != "" {
		words = append(words, word)
	}
	return words
}

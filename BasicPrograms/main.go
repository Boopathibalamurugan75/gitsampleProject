package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var choice int
	for {
		fmt.Println("Enter your choice : ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			findWord()
		case 2:
			findFirstRepeatedChar()
		case 3:
			swapping()
		case 4:
			findMaxMin()
		case 5:
			maximumofThree()
		case 0:
			return
		default:
			fmt.Println("-------Invalid Option-------")
		}
	}
}

// Find the largest and Smallest word in given string
func findWord() {
	trigger := false
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the String : ")
	inputValue, _ := reader.ReadString('\n')
	for _, input := range inputValue {
		if input > 'a' || input == '\n' || input == ' ' {
			trigger = true
		} else {
			trigger = false
			break
		}
	}
	if trigger {
		wordlength := strings.Fields(inputValue)
		largest := wordlength[0]
		smallest := wordlength[0]
		for _, word := range wordlength {
			if len(word) > len(largest) {
				largest = word
			}
			if len(word) < len(smallest) {
				smallest = word
			}
		}
		fmt.Println("The Largest Word is : ", largest)
		fmt.Println("The Smallest word is : ", smallest)
	} else {
		fmt.Print("Please enter string value \n")
	}

}

// Find the repeated character in a given string.
func findFirstRepeatedChar() {
	checkChar := true
	trigger := false
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the String : ")
	input, _ := reader.ReadString('\n')
	for _, inputData := range input {
		if inputData >= 'a' || inputData == '\n' || inputData == ' ' {
			trigger = true
		} else {
			trigger = false
			break
		}
	}
	if trigger {
		for i := 0; i < len(input); i++ {
			for j := i + 1; j < len(input); j++ {
				if input[i] == input[j] {
					fmt.Printf("The first repetitive character is: %c\n", input[i])
					checkChar = false
					return
				}
			}
		}
		if checkChar {
			fmt.Printf("No repetitive character found in %s\n", input)
		}
	} else {
		fmt.Printf("Please enter string value \n")

	}
}

// Swap the values of two variables without a temporary variable.
func swapping() {
	var firstValue, secondValue int
	fmt.Print("Enter first value  : ")
	fmt.Scan(&firstValue)
	fmt.Print("Enter Second value  : ")
	fmt.Scan(&secondValue)
	fmt.Println("Before Swapping : ", firstValue, secondValue)
	firstValue, secondValue = secondValue, firstValue
	fmt.Println("After Swapping : ", firstValue, secondValue)
}

// Find the maximum and minimum element in an array
func findMaxMin() {
	inputData := [5]int{}
	trigger := true
	fmt.Println("Enter", len(inputData), "values : ")
	for i := 0; i < len(inputData); i++ {
		fmt.Printf("Element - %d : ", i)
		_, err := fmt.Scanln(&inputData[i])
		if err != nil {
			fmt.Println("Please enter Integer value")
			trigger = false
			return
		}
	}
	if trigger {
		for i := 1; i < len(inputData); i++ {
			if inputData[i] < inputData[i-1] {
				inputData[i], inputData[i-1] = inputData[i-1], inputData[i]
				i = 0
			}
		}
		fmt.Println("The Original Array is : ", inputData)
		fmt.Println("Minimum Number is : ", inputData[0])
		fmt.Println("Maximum Number is : ", inputData[len(inputData)-1])
	}
}

// Find the maximum of three.
func maximumofThree() {
	var firstNumber, secondNumber, thirdNumber int
	fmt.Print("Enter the first number : ")
	fmt.Scan(&firstNumber)
	fmt.Print("Enter the second number : ")
	fmt.Scan(&secondNumber)
	fmt.Print("Enter the third number : ")
	fmt.Scan(&thirdNumber)
	if firstNumber > secondNumber && firstNumber > thirdNumber {
		fmt.Println("The Biggest number is : ", firstNumber)
	} else if secondNumber > firstNumber && secondNumber > thirdNumber {
		fmt.Println("The Biggest number is : ", secondNumber)
	} else {
		fmt.Println("The Biggest number is : ", thirdNumber)
	}
}
